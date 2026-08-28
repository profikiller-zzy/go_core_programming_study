package main

import (
	"fmt"
	"unicode/utf8"
)

/*
题目描述：
流式输出场景下，如何去检测敏感词。你的输入并不是一整个文本，而是被切成长度不等的一些小字符串[]string。
这就会出现一个问题：上一个字符串的结尾和下一个字符串的开头，可能拼起来才是一个完整的关键词。那么，你这种边界情况要如何检测到？关键词有很多，如果每个关键词都去遍历全文，时间复杂度就会变成 O(NM)，其中 M 是关键词的个数，N 是整个文本的长度。这样时间复杂度太低了。
如何让时间复杂度变成 O(N) 呢？就是你只需要把输入的所有文本过一遍，就可以检测到所有的敏感词。
*/

type Match struct {
	match   string
	runeLen int
}

type Trie struct {
	Children map[rune]*Trie
	Rune     rune
	// 失败指针，当前字符串中所有真后缀在前缀树中的最长前缀
	// 例如，假设敏感词是abc、bcd，那么 Fail(ab) = b，Fail(abc) = bc
	Fail *Trie
	// Fail 链上最近的关键词结束节点
	Output *Trie
	// 记录以当前节点为结尾的所有关键词
	Pattern []Match
}

func NewTrieNode(r rune) *Trie {
	return &Trie{
		Children: make(map[rune]*Trie),
		Rune:     r,
	}
}

// ACAutomaton 完整AC自动机
type ACAutomaton struct {
	Root       *Trie
	state      *Trie // 当前流式匹配状态
	runeOffset int   // 当前已经读取了多少个rune字符
	built      bool
}

func NewACAutomaton() *ACAutomaton {
	root := NewTrieNode('0')
	return &ACAutomaton{
		Root:  root,
		state: root,
	}
}

func (ac *ACAutomaton) Insert(word string) {
	if ac.built { // AC 自动机构建完成后不能继续插入关键词
		return
	}
	if word == "" {
		return
	}
	cur := ac.Root
	for _, r := range word {
		if cur.Children == nil {
			cur.Children = make(map[rune]*Trie)
		}
		next, ok := cur.Children[r]
		if !ok {
			next = NewTrieNode(r)
			cur.Children[r] = next
		}
		cur = next
	}
	cur.Pattern = append(cur.Pattern, Match{
		word,
		utf8.RuneCountInString(word),
	})
	return
}

// Build 使用BFS构建Fail指针，规则：要计算 child.Fail：
// 1. 从 parent.Fail 开始。
// 2. 检查它有没有字符 r 的子节点。
// 3. 有则将那个子节点作为 child.Fail。
// 4. 没有则继续沿 Fail 指针回退。
// 5. 一直找不到就指向 root。
func (ac *ACAutomaton) Build() {
	root := ac.Root
	root.Fail = root
	root.Output = nil

	queue := make([]*Trie, 0)
	// 初始化第一层节点，他们的Fail一定指向root，因为只有一个字符的string的真后缀为空
	for _, child := range root.Children {
		child.Fail = root
		child.Output = nil
		queue = append(queue, child)
	}

	for head := 0; head < len(queue); head++ {
		parent := queue[head]
		for r, child := range parent.Children {
			fail := parent.Fail

			// 沿着parent的fail链路一直找能够接受r的状态
			for fail != root {
				if _, ok := fail.Children[r]; ok {
					break
				}
				fail = fail.Fail
			}
			// 这里退出时可能是fail已经为root，循环退出没有检查，再检查一遍
			if next, ok := fail.Children[r]; ok {
				child.Fail = next
			} else {
				child.Fail = root
			}
			// 找到fail链上最近的敏感词结束节点
			if len(child.Fail.Pattern) > 0 {
				child.Output = child.Fail
			} else {
				child.Output = child.Fail.Output
			}

			queue = append(queue, child)
		}
	}
	ac.state = root
	ac.runeOffset = 0
	ac.built = true
}

func (ac *ACAutomaton) Advance(r rune) {
	root := ac.Root
	for ac.state != root {
		if next, ok := ac.state.Children[r]; ok {
			ac.state = next
			return
		}
		ac.state = ac.state.Fail
	}
	if next, ok := ac.state.Children[r]; ok {
		ac.state = next
	} else {
		ac.state = root
	}
}

type Pattern struct {
	word       string
	start, end int
}

func (ac *ACAutomaton) Feed(chunk string) []Pattern {
	if !ac.built {
		return nil
	}
	var result []Pattern
	for _, r := range chunk {
		ac.Advance(r)
		end := ac.runeOffset + 1
		for _, pattern := range ac.state.Pattern {
			result = append(result, Pattern{
				pattern.match, end - pattern.runeLen, end,
			})
		}

		// 还要检查 ac.state fail链上的后缀关键词
		for output := ac.state.Output; output != nil; output = output.Output {
			for _, pattern := range output.Pattern {
				result = append(result, Pattern{
					pattern.match, end - pattern.runeLen, end,
				})
			}
		}
		ac.runeOffset = end
	}
	return result
}

func SensitiveWordDetection(chunks []string, sensitiveWords []string) []Pattern {
	ac := NewACAutomaton()
	for _, word := range sensitiveWords {
		ac.Insert(word)
	}
	ac.Build()
	var result []Pattern
	for _, chunk := range chunks {
		curRes := ac.Feed(chunk)
		result = append(result, curRes...)
	}
	return result
}

func main() {
	chunks := []string{
		"xxab",
		"cd这是敏",
		"感",
		"词结束",
	}
	sensitiveWords := []string{
		"abc", "bcd", "敏感", "敏感词",
	}

	result := SensitiveWordDetection(chunks, sensitiveWords)
	fmt.Println(result)
}
