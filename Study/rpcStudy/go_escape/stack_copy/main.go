package main

import (
	"os"
	"strconv"
	"strings"
	"unsafe"
)

// ptr0 存储指针最开始的值，他的类型是uintptr
var ptr0 uintptr

// depth 存储递归深度
var depth int

func main() {
	var n int
	ptr0 = uintptr(unsafe.Pointer(&n))
	f(&n)
}

// f 会一直递归，随着递归进行栈不断增长，超出初始栈空间后，go会分配一个更大的栈空间，将原本的栈空间的内容复制到新的栈空间
// 然后栈中的变量地址也会改变，runtime会将所有受影响的变量的地址更新为新的地址
// 当栈指针改变时，打印递归深度和原始指针地址以及改变后的指针地址
func f(ptr *int) {
	depth++
	if uintptr(unsafe.Pointer(ptr)) != ptr0 {
		print(ptr)
		os.Exit(0)
	}
	f(ptr)
}

func print(ptr *int) {
	var buf strings.Builder
	buf.WriteString("ptr0 = 0x")
	buf.WriteString(strconv.FormatUint(uint64(ptr0), 16))
	buf.WriteString("; ")
	buf.WriteString("ptrn = 0x")
	buf.WriteString(strconv.FormatUint(uint64(uintptr(unsafe.Pointer(ptr))), 16))
	buf.WriteString("; depth = ")
	buf.WriteString(strconv.Itoa(depth))
	buf.WriteRune('\n')
	os.Stdout.WriteString(buf.String())
}
