package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 练习6.1: 为bit数组实现下面这些方法
//
//
//func (*IntSet) Len() int      // return the number of elements
//func (*IntSet) Remove(x int)  // remove x from the set
//func (*IntSet) Clear()        // remove all elements from the set
//func (*IntSet) Copy() *IntSet // return a copy of the set

func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			if word&(1<<uint(i)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	index, offset := x/64, x%64
	if index > len(s.words) {
		return
	}
	s.words[index] &= ^(1 << uint(offset))
}

func (s *IntSet) Clear() {
	s.words = nil // 直接将s.words置为nil
}

func (s *IntSet) Copy() *IntSet {
	// 1.创建一个新的IntSet
	var newIntSet IntSet
	// 2.使用可变参数将s.words赋值给newIntSet.words
	newIntSet.words = append(newIntSet.words, s.words...)
	// 3.返回新的IntSet
	return &newIntSet
}

// 练习 6.2： 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。

func (s *IntSet) AddAll(nums ...int) {
	if len(nums) > 0 {
		for _, num := range nums {
			s.Add(num)
		}
	}
}

// 练习 6.3： (*IntSet).UnionWith会用|操作符计算两个集合的并集，我们再为IntSet实现另外的几个函数IntersectWith（交集：元素在A集合B集合均出现），
// DifferenceWith（差集：元素出现在A集合，未出现在B集合），SymmetricDifference（并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A）。

// IntersectWith 将s变为s与t之间的交集
func (s *IntSet) IntersectWith(t *IntSet) {
	for index, _ := range s.words {
		// 如果越界，表示在t中没有这些数，那么将s中对应的数全部置为0
		if index >= len(t.words) {
			s.words[index] = 0
		} else {
			// 与运算的结果就是s和t的交集
			s.words[index] &= t.words[index]
		}
	}
}

// DifferenceWith 计算出现在s中但未出现在t中的集合
func (s *IntSet) DifferenceWith(t *IntSet) {
	// 只需要去除s中两个集合共有的部分即可，故只需要遍历s即可，无需处理越界部分
	for index, word := range s.words {
		if index >= len(t.words) { // 越界部分直接返回
			return
		}
		// interWord 是两个集合的交集部分
		var interWord = t.words[index] & word
		// 将interWord取反之后再与即可剔除两个集合之间的交集部分
		s.words[index] &= ^interWord
	}
}

// SymmetricDifference 就是计算s与t之间的并集与两者交集之间的差
func (s *IntSet) SymmetricDifference(t *IntSet) {
	// 创建一个空的IntSet
	var u IntSet
	// 将s和t的并集赋值给u, 相当于u里面有s和t的所有元素
	u.UnionWith(s)
	u.UnionWith(t)
	// 将s和t的交集赋值给s, 相当于s中只有s和t中共同都有的元素了, 这也就是交集
	s.IntersectWith(t)
	// 将u和s的差集赋值给u, 相当于u扣掉了s和t的交集
	u.DifferenceWith(s)
	// 将u赋值给s
	*s = u
}

// 练习 6.4： 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。

// Elems 返回一个包含集合中所有元素的切片。
func (s *IntSet) Elems() (elems []int) {
	// 创建一个空的切片
	// 遍历s.words
	for index, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*index+j)
			}
		}
	}
	return elems
}

// 练习 6.5： 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
// 修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。当然了，这里我们可以不用简单粗暴地除64，
// 可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)

const PLATFORM_ARCH = 32 << (^uint(0) >> 63)

// 在Go语言中，实现了随平台位数的变化而变化，在32位系统中uint长度与uint32一致
// 在64位系统中长度和uint64一致
// 修改只需要将64替换成PLATFORM_ARCH即可

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(9)
	x.Add(144)
	fmt.Println("x:", x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println("y:", y.String()) // "{9 42}"

	x.IntersectWith(&y)
	fmt.Println("x IntersectWith y:", x.String()) // "{9}"

	x.Clear()
	x.Add(1)
	x.Add(9)
	x.Add(144)

	x.DifferenceWith(&y)
	fmt.Println("x DifferenceWith y:", x.String()) // "{1 144}"

	x.Clear()
	x.Add(1)
	x.Add(9)
	x.Add(144)

	x.SymmetricDifference(&y)
	fmt.Println("x SymmetricDifference y:", x.String()) // "{1 42 144}"
	fmt.Println(x.Elems())
}
