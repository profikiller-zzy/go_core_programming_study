package main

import (
	"fmt"
	"unsafe"
)

// 模拟 runtime/internal 的 hmap 结构（仅保留常用字段）
type hmap struct {
	count      int // map 中实际存储的键值对数量
	flags      uint8
	B          uint8 // 2^B 个桶
	noverflow  uint16
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra      unsafe.Pointer
}

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// 强制将 map 转为 *hmap
	mapHeader := *(*uintptr)(unsafe.Pointer(&m)) // 获取 map 接口底层指针
	h := (*hmap)(unsafe.Pointer(mapHeader))      // 转为 hmap 指针

	// 输出信息
	fmt.Println("🚀 访问 map 的底层结构：")
	fmt.Printf("map.count:     %d\n", h.count)
	fmt.Printf("map.B:         %d\n", h.B)
	fmt.Printf("map.buckets:   %p\n", h.buckets)
	fmt.Printf("map.hashSeed:  %d\n", h.hash0)
	fmt.Printf("map.noverflow: %d\n", h.noverflow)
}
