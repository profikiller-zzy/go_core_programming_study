package main

type Trie struct {
	Byte     byte
	Children map[byte]*Trie
	Fail     *Trie
	Output   *Trie
	Match    []string
}

func newTrieNode(b byte) *Trie {
	return &Trie{
		Children: make(map[byte]*Trie),
		Byte:     b,
	}
}

type StreamChecker struct {
	Root  *Trie
	State *Trie
	Built bool
}

func (sc *StreamChecker) InsertWord(word string) {
	if sc.Built {
		return
	}
	if word == "" {
		return
	}
	cur := sc.Root
	for index := 0; index < len(word); index++ {
		b := word[index]
		if cur.Children == nil {
			cur.Children = make(map[byte]*Trie)
		}
		next, ok := cur.Children[b]
		if !ok {
			next = newTrieNode(b)
			cur.Children[b] = next
		}
		cur = next
	}
	cur.Match = append(cur.Match, word)
}

func (sc *StreamChecker) Build() {
	root := sc.Root
	root.Fail = root
	root.Fail = nil
	var queue []*Trie
	for _, child := range sc.Root.Children {
		child.Fail = root
		child.Output = nil
		queue = append(queue, child)
	}
	for head := 0; head < len(queue); head++ {
		parent := queue[head]
		for b, child := range parent.Children {
			fail := parent.Fail
			for fail != root {
				if _, ok := fail.Children[b]; ok {
					break
				}
				fail = fail.Fail
			}
			next, ok := fail.Children[b]
			if ok {
				child.Fail = next
			} else {
				child.Fail = root
			}
			if len(child.Fail.Match) > 0 {
				child.Output = child.Fail
			} else {
				child.Output = child.Fail.Output
			}
			queue = append(queue, child)
		}

	}
	sc.State = root
	sc.Built = true
}

func (sc *StreamChecker) Advance(b byte) {
	root := sc.Root
	for sc.State != root {
		if _, ok := sc.State.Children[b]; ok {
			break
		}
		sc.State = sc.State.Fail
	}
	next, ok := sc.State.Children[b]
	if ok {
		sc.State = next
	} else {
		sc.State = root
	}
}

func Constructor(words []string) StreamChecker {
	root := newTrieNode('0')
	sc := &StreamChecker{
		Root: root,
	}
	for _, word := range words {
		sc.InsertWord(word)
	}
	sc.Build()
	return *sc
}

func (this *StreamChecker) Query(letter byte) bool {
	this.Advance(letter)
	if len(this.State.Match) > 0 {
		return true
	}
	if this.State.Output != nil {
		return true
	}
	return false
}
