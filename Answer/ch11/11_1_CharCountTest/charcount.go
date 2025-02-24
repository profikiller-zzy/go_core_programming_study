package _1_1_CharCountTest

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func charCount(data []byte) (unicodeCounts map[rune]int, utfCounts [utf8.UTFMax + 1]int, invalid int) {
	unicodeCounts = make(map[rune]int) // Unicode 字符统计
	//var utflen [utf8.UTFMax + 1]int    // 统计UTF-8字符各个字节数的字符有多少个
	invalid = 0 // 非法UTF-8字符数

	reader := bytes.NewReader(data)
	in := bufio.NewReader(reader)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		unicodeCounts[r]++
		utfCounts[n]++
	}
	//fmt.Printf("rune\tcount\n")
	//for c, n := range unicodeCounts {
	//	fmt.Printf("%q\t%d\n", c, n)
	//}
	//fmt.Print("\nlen\tcount\n")
	//for i, n := range utflen {
	//	if i > 0 {
	//		fmt.Printf("%d\t%d\n", i, n)
	//	}
	//}
	//if invalid > 0 {
	//	fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	//}
	return
}
