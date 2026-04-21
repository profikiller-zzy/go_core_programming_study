package main

import (
	"context"
	"errors"
	"net"
	"sync"
)

// Factory 底层创建连接的函数
type Factory func() (net.Conn, error)

type Pool struct {
	factory       Factory    // 底层创建连接的函数
	maxConnCount  int        // 最大连接数
	openConnCount int        // 已经打开的连接数
	idleConn      []net.Conn // 空闲的连接数
	mu            sync.Mutex
	closed        bool          // 标记是否关闭
	waitChan      chan struct{} // 阻塞拿不到连接的协程
}

// Conn 对底层连接进行封装
type Conn struct {
	conn net.Conn
	pool *Pool // pool非nil，代表该连接是从连接池中拿的
	once sync.Once
}

func NewPool(factory Factory, maxConnCount int) (*Pool, error) {
	if factory == nil {
		return nil, errors.New("factory is nil")
	}
	if maxConnCount <= 0 {
		maxConnCount = 100
	}
	return &Pool{
		factory:      factory,
		maxConnCount: maxConnCount,
	}, nil
}

// Acquire 获取一个连接
// 1. 如果pool已经关闭了，直接返回空连接和错误
// 2. 如果有空闲连接，直接返回一个空闲连接
// 3. 如果没有空闲连接并且现在的连接数没有到达上限，创建一个新连接并返回
// 4. 协程拿不到连接，那么就阻塞在Acquire中
func (p *Pool) Acquire(ctx context.Context) (*Conn, error) {
	for {
		p.mu.Lock()

		if p.closed {
			p.mu.Unlock()
			return nil, errors.New("pool is closed")
		}

		if n := len(p.idleConn); n > 0 {
			c := p.idleConn[n-1]
			p.idleConn[n-1] = nil
			p.idleConn = p.idleConn[:n-1]
			p.mu.Unlock()
			return &Conn{
				conn: c,
				pool: p,
			}, nil
		}

		if p.openConnCount < p.maxConnCount {
			p.openConnCount++
			p.mu.Unlock()
			conn, err := p.factory()
			if err != nil {
				p.mu.Lock()
				p.openConnCount--
				p.signalWaited() // 这里的语义是，本来该协程占用一个连接名额，但是创建conn失败了，所以唤醒阻塞的协程来抢名额
				p.mu.Unlock()
				return nil, err
			}
			return &Conn{
				conn: conn,
				pool: p,
			}, nil
		}

		waitChan := p.waitChan // 这里一定要在锁里面拿该waitChan
		p.mu.Unlock()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-waitChan:
			// 阻塞在这里，等待被唤醒并进行下一轮重试
		}
	}
}

// put 归还一个连接
// 1. 如果归还时连接池已经被关闭了，直接销毁这个连接
// 2. 如果没有关闭，则放回连接池的空闲连接中
func (p *Pool) put(conn net.Conn) error {
	if conn == nil {
		return nil
	}
	var needClose bool // 是否需要销毁这个连接
	p.mu.Lock()
	switch {
	case p.closed:
		needClose = true
		p.openConnCount--
	default:
		p.idleConn = append(p.idleConn, conn)
	}

	p.signalWaited()
	p.mu.Unlock()
	if needClose {
		return conn.Close()
	}
	return nil
}

// Close 关闭连接池
// 1. 首先释放所有空闲连接
// 2. 对于那些由协程持有的连接，等待协程完成任务后调用Conn.Close()关闭，函数发现连接池已经关闭则直接销毁这个连接
func (p *Pool) Close() error {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return nil
	}
	p.closed = true
	idleConns := p.idleConn
	p.openConnCount -= len(idleConns)
	p.signalWaited()
	p.idleConn = nil
	p.mu.Unlock()

	var err error
	for _, conn := range idleConns {
		err = conn.Close()
	}
	return err
}

// signalWaited 关闭当前p中的waitChan以唤醒所有阻塞的协程，并且将一个新的协程赋值在它之
// 注意：调用该函数的调用方必须持有锁
func (p *Pool) signalWaited() {
	close(p.waitChan)
	p.waitChan = make(chan struct{})
}

// Close 对于调用方，调用了Close()方法的连接不应该继续使用
func (c *Conn) Close() error {
	var err error
	// 这里使用once.Do的原因是，如果同一个Conn重复归还到池子中，会出大问题
	c.once.Do(func() {
		if c.pool == nil { // 如果这里这个连接不是从池子拿的，可能是用户自己创建的
			if c.conn != nil {
				err = c.conn.Close()
			}
			return
		}
		err = c.pool.put(c.conn)
	})
	return err
}
