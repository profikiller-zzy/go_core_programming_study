package main

func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		result <<= 1 // 左移一位
		curBit := num & 1
		result |= curBit // 将当前位添加到结果中
		num >>= 1        // 右移一位，处理下一个位
	}
	return result
}
