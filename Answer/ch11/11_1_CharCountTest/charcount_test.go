package _1_1_CharCountTest

import (
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	// 首先定义输入数据
	tests := []struct {
		name            string // 各种不同的输入场景
		input           []byte // 输入的数据
		expectedUnicode map[rune]int
		expectedUtf     [utf8.UTFMax + 1]int
		expectedInvalid int
	}{
		{
			name:  "ASCII characters",
			input: []byte("hello world"),
			expectedUnicode: map[rune]int{
				'h': 1, 'e': 1, 'l': 3, 'o': 2, ' ': 1, 'w': 1, 'r': 1, 'd': 1,
			},
			expectedUtf:     [utf8.UTFMax + 1]int{0, 11, 0, 0, 0},
			expectedInvalid: 0,
		},
		{
			name:  "Mixed ASCII and Unicode",
			input: []byte("hello 世界"),
			expectedUnicode: map[rune]int{
				'h': 1, 'e': 1, 'l': 2, 'o': 1, ' ': 1, '世': 1, '界': 1,
			},
			expectedUtf:     [utf8.UTFMax + 1]int{0, 7, 0, 2, 0},
			expectedInvalid: 0,
		},
		{
			name:            "Invalid UTF-8 characters",
			input:           []byte{0xff, 0xfe, 0xfd},
			expectedUnicode: map[rune]int{},
			expectedUtf:     [utf8.UTFMax + 1]int{0, 0, 0, 0, 0},
			expectedInvalid: 3,
		},
	}

	for _, tt := range tests {
		// 执行名字为 name 的子测试 f ，并报告 f 在执行过程中是否出现了任何失败。Run 将一直阻塞直到 f 的所有并行测试执行完毕。
		t.Run(tt.name, func(t *testing.T) {
			unicodeCounts, utfCounts, invalid := charCount(tt.input)

			// 比较 Unicode 字符统计
			for r, count := range tt.expectedUnicode {
				if unicodeCounts[r] != count {
					t.Errorf("For rune %q: expected %d, got %d", r, count, unicodeCounts[r])
				}
			}

			// 比较 UTF 字节长度统计
			for i, count := range tt.expectedUtf {
				if utfCounts[i] != count {
					t.Errorf("For UTF length %d: expected %d, got %d", i, count, utfCounts[i])
				}
			}

			// 比较非法字符数量
			if invalid != tt.expectedInvalid {
				t.Errorf("Expected invalid count %d, got %d", tt.expectedInvalid, invalid)
			}
		})
	}
}
