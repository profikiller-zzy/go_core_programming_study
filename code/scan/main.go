package main

import "fmt"

func main() {
	//age := 20
	//switch {
	//case age == 20:
	//	fmt.Println("age == 20")
	//case age == 30:
	//	fmt.Println("age == 20")
	//default:
	//	fmt.Println("unknown")
	//}
	//
	//// switch 穿透
	//var num int = 10
	//switch num {
	//case 10:
	//	fmt.Println("ok1")
	//	fallthrough // 使用fallthrough来实现switch穿透，即使下一个case条件不满足，也会执行下一个分支的代码
	//	//默认只会穿透一层
	//case 20:
	//	fmt.Println("ok2")
	//case 30:
	//	fmt.Println("ok3")
	//}
	//
	//num = 3
	//switch num {
	//case 1:
	//	fmt.Println("num is 1")
	//case 2:
	//	fmt.Println("num is 2")
	//	fallthrough
	//case 3:
	//	fmt.Println("This will be executed due to fallthrough")
	//default:
	//	fmt.Println("default case")
	//}

	//var x interface{}
	//var y = 10.0
	//x = y
	//x = func(int2 int) float64 {
	//	return 0.0
	//}
	//
	//switch i := x.(type) {
	//case nil:
	//	fmt.Printf("x 的类型: %T\n", i)
	//case int:
	//	fmt.Printf("x 是 int 型\n")
	//case float64:
	//	fmt.Printf("x 是 float64 型\n")
	//case func(int) float64:
	//	fmt.Printf("x 是 func(int) float64 型\n")
	//case bool, string:
	//	fmt.Printf("x 是 bool 或 string 型\n")
	//default:
	//	fmt.Printf("未知类型\n")
	//}

	var s string = "我爱中国"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
	for i, ch := range s {
		fmt.Printf("索引：%d\t字符：%c\tUnicode码点:%U\n", i, ch, ch)
	}

}
