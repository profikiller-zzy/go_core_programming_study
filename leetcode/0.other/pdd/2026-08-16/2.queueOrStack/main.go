package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述：

多多君在游乐场发现了一条神奇的通道。入口处有一串彩球按顺序
滚入通道，经过通道内部的暂存区后，再从出口按某种顺序滚出。

通道有两种工作模式，但多多君忘记当时使用的是哪一种：

1. 队列模式：
   暂存区像队列一样，先进入的彩球先出来，即先进先出。

2. 栈模式：
   暂存区像栈一样，后进入的彩球先出来，即后进先出。

两种模式的工作流程相同：

彩球按照给定顺序从入口依次滚入。多多君可以在任意时刻选择将
暂存区最前面或最上面的彩球取出，让它从出口滚出；也可以选择
暂时不取出彩球，继续让更多彩球进入暂存区。

最终，所有彩球都必须从出口滚出。

现在给出彩球进入通道的顺序 S 和从通道出来的顺序 T，请判断
通道当时使用的是队列模式、栈模式，还是两种模式都有可能。


输入描述：

输入共两行。

第一行输入一个字符串 S，表示彩球进入通道的顺序。

第二行输入一个字符串 T，表示彩球从通道出来的顺序。

1 <= |S| = |T| <= 100000

S 和 T 均由小写英文字母组成。

S 和 T 的长度相同，并且包含的字符及其数量相同，
即 T 是 S 的一个排列。


输出描述：

输出一个字符串，表示判断结果：

如果只能是队列模式，输出：

queue

如果只能是栈模式，输出：

stack

如果队列模式和栈模式都可能，输出：

both

如果两种模式都不可能，输出：

neither


样例 1：

输入：
abc
abc

输出：
both

样例解释：

队列模式为先进先出，可以得到 abc。

栈模式下，可以依次执行：

a 进入，a 出来；
b 进入，b 出来；
c 进入，c 出来。

同样可以得到 abc，因此输出 both。


样例 2：

输入：
abc
cba

输出：
stack

样例解释：

队列模式只能按照先进先出的顺序得到 abc，无法得到 cba。

栈模式可以让 a、b、c 全部进入暂存区，然后依次出栈，
得到 cba，因此输出 stack。
*/

func solution(s, t string) string {
	isStack, isQueue := checkStack(s, t), s == t
	if isStack && isQueue {
		return "both"
	} else if isStack {
		return "stack"
	} else if isQueue {
		return "queue"
	}
	return "neither"
}

func checkStack(s, t string) bool {
	n := len(s)
	stack := make([]byte, 0, n)
	var checkedIndex int
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		for checkedIndex < n && len(stack) > 0 && stack[len(stack)-1] == t[checkedIndex] {
			stack = stack[:len(stack)-1]
			checkedIndex++
		}
	}
	return checkedIndex == n
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var s, t string
	fmt.Fscan(in, &s, &t)
	fmt.Println(solution(s, t))
}
