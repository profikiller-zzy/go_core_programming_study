# 一、GO语言简述

### 1. 垃圾回收机制

内存自动回收，不用开发员自己管理，程序员只用关心代码逻辑，不用担心内存泄漏。

### 2. 原生支持并发（重要）

- 从语言方面支持并发，实现简单；
- goroutine，轻量级线程，**可实现大量并发**，高效率利用核；
- 基于CPS并发模型(Communicating Sequential Processes)实现；

### 3. 管道通信机制

通过管道Channel，可以实现不同的goroutine之间的相互通信

### 4. 函数可以返回多个值

```go
func sumAndSub(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
```

### 5. 创新：切片、延时执行函数defer语句等

### 6. GO 语言开发注意事项

- GO应用程序执行的入口是main()函数

### 

# 二、GO变量

## 概念

什么是变量：**变量**是用来存储值的命名实体。变量在编程中扮演着重要角色，允许你存储和操作数据。Go 中的变量可以存储不同的数据类型，比如整数、浮点数、字符串等。

## 1. 使用变量

### 1.1  使用 `var` 关键字显式声明

这种方式适用于需要显式声明变量类型的情况，并且可以选择是否同时赋值。

```go
// 声明变量并赋初值
var name string = "Alice"
var age int = 30

// 只声明变量，不赋值，使用默认的零值
var height float64
```

在这种方式下，如果不赋初值，Go 会给变量赋一个默认的**零值**。如上例中的 `height` 默认值为 `0.0`。

### 1.2 使用 `var` 关键字并省略类型（类型推断）

当为变量赋值时，Go 可以根据赋值的内容自动推断出变量的类型，因此你可以省略类型声明。

```go
var name = "Bob"   // Go 自动推断为 string 类型
var age = 25       // Go 自动推断为 int 类型
```

### 1.3 简短声明（仅限函数内部）

在函数内部，可以使用简短声明符 `:=` 进行变量声明并赋值。Go 会根据右侧的值自动推断变量类型。

```go
name := "Charlie"
age := 22
height := 1.80
```

> 注意：简短声明只能在函数内部使用，不能用于全局变量声明。

### 1.4 多变量声明

Go 语言支持一次性声明多个变量，可以使用多种形式：

- **同时声明多个变量并赋值：**

```go
var x, y int = 10, 20
```

- **使用简短声明同时声明多个变量：**

```go
x, y := 10, 20
```

- **多行声明变量：**

```go
var (
    name   string = "David"
    age    int    = 35
    height float64 = 1.75
)
```

### 1.5 如何在程序中查看一个变量的数据类型和所占的字节数：

```go
var x int64
fmt.Printf("x 的类型为 %T\\nx 所占的字节数为 %d", x, unsafe.Sizeof(x))
```

在使用变量类型的时候，在保证程序正常运行的前提下，尽量使用占用字节较小的变量类型。

在不指定类型的情况下，将变量赋值为一个整形变量，GO会自动将这个变量的类型设置为int

```go
var y = 100
fmt.Printf("y 的类型为 %T\\n", y)

// y 的类型为 int
```

## 2. 浮点数 float

在 Go 语言中，小数类型主要通过 **浮点数类型** 来表示。Go 提供了两种主要的浮点数类型，用于处理不同精度的浮点运算：

### 2.1 **`float32`**

- `float32` 是 32 位的浮点数，表示精度较低的小数，符合 IEEE-754 标准。
- 它能精确表示小数点后大约 7 位的有效数字。

### 2.2 **`float64`**

- `float64` 是 64 位的浮点数，表示高精度的小数，也是 Go 语言中常用的浮点数类型，符合 IEEE-754 标准。
- 它能精确表示小数点后大约 15 位的有效数字。

> 在 Go 中，默认的小数类型是 float64，如果没有显式声明类型，带有小数点的数字会自动被识别为 float64。

```go
package main

import (
    "fmt"
)

func main() {
    var f1 float32 = 3.14159  // 显式声明为 float32
    var f2 float64 = 3.141592653589793  // 显式声明为 float64
    f3 := 2.71828  // 自动推断为 float64

    fmt.Printf("f1 (float32): %f\\n", f1)
    fmt.Printf("f2 (float64): %.15f\\n", f2) // 打印 15 位小数
    fmt.Printf("f3 (default float64): %.5f\\n", f3)
}
```

结果：

```go
f1 (float32): 3.141590
f2 (float64): 3.141592653589793
f3 (default float64): 2.71828
```

`Golang`的浮点数申明默认为`float64`

```go
var z = 1.01
fmt.Printf("z 的类型为 %T\\n", z)

// z 的类型为 float64
```

### 2.3 浮点数练习

用`Printf`函数的`%g`参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适。所有的这三个打印形式都可以指定打印的宽度和控制打印精度。

```Go
for x := 0; x < 8; x++ {
    fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
}
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
    "fmt"
    "log"
    "math"
    "net/http"
)

const (
    width, height = 600, 320            // canvas size in pixels
    cells         = 100                 // number of grid cells
    xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit
    angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
    http.HandleFunc("/", handler)
    err := http.ListenAndServe("localhost:8080", nil)
    log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
    //不加Content-Type，浏览器显示的是svg文件内容，而不是解析显示为图片。
    w.Header().Set("Content-Type", "image/svg+xml")

    fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j)
            bx, by := corner(i, j)
            cx, cy := corner(i, j+1)
            dx, dy := corner(i+1, j+1)
            fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z.
    z := f(x, y)

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

// f 用于计算z z = Sin(r) / r
func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0,0)
    return math.Sin(r) / r
}

//!-
```

运行效果

![image.png](C:\Users\hasee\GolandProjects\go_core_programming\md\image.png)

# 3. 字符串

Go语言中没有专门的字符类型，如果要存储一个字符的话，通常使用`rune`类型。

在 Go 语言中，字符串是由字节（`byte`）组成的。Go 的字符串实际上是一个字节序列，且是不可变的。这意味着每个字符串都是一组按顺序排列的字节，底层存储为 UTF-8 编码的字节序列。

```go
s := "你好"
fmt.Println(len(s)) // 输出 6，因为 "你" 和 "好" 各自占用 3 个字节
```

### 3.1 字符类型使用细节

1. 字符常量是用单引号(`'`)括起来的单个字符。例如：

```go
var c1 byte = 'a'
var c2 int = '中'
var c3 byte = '9'
```

1. Go 中允许使用转义字符`'\\'`来将其后的字符转义为特殊字符型常量。例如：`var c3 char = '\\n' // '\\n'表示换行符`
2. Go 语言的字符使用 **UTF-8 编码**，这是一种变长编码，英文字母 1 个字节，汉字 3 个字节。
3. 在 Go 中，字符的本质是一个整数，直接输出时，是该字符对应的 UTF-8 编码的码值。
   - 可以直接给某个变量赋一个数字，然后按格式化输出时 `%c` 会输出该数字对应的 Unicode 字符。
4. 字符类型是可以进行运算的，相当于一个整数，因为它有对应的 Unicode 码。

在 Go 语言中，使用反引号（```）可以定义**原始字符串字面量**，这意味着字符串中的所有内容都将按照原样保存，不会处理转义字符（如 `\\n` 不会被解释为换行符）。示例：

```go
package main

import "fmt"

func main() {
    // 使用反引号定义的原始字符串
    rawString := `这是一个
多行字符串，
其中的\\n不会被转义为换行符。`

    fmt.Println(rawString)
}
```

### 3.2 字符串拼接方式

#### 3.2.1 使用 `+` 操作符拼接

这是最简单直接的拼接方式，适合少量字符串的拼接。

```go
import "fmt"

func main() {
    str1 := "Hello"
    str2 := "World"
    result := str1 + " " + str2
    fmt.Println(result) // 输出：Hello World
}
```

#### 3.2.2 使用 `fmt.Sprintf` 格式化拼接

`fmt.Sprintf` 可以通过格式化操作拼接字符串，适合需要格式化输出的场景。

```go
import "fmt"

func main() {
    name := "Go"
    version := "1.18"
    result := fmt.Sprintf("Language: %s, Version: %s", name, version)
    fmt.Println(result) // 输出：Language: Go, Version: 1.18
}
```

#### 3.2.3 使用 `strings.Join` 函数

如果要拼接字符串切片，可以使用 `strings.Join`，这种方法效率高于直接使用 `+`，特别是当需要拼接大量字符串时。

```go
import (
    "fmt"
    "strings"
)

func main() {
    words := []string{"Hello", "Go", "Language"}
    result := strings.Join(words, " ")
    fmt.Println(result) // 输出：Hello Go Language
}
```

#### 3.2.4 使用 `bytes.Buffer`

当需要拼接大量字符串时，使用 `bytes.Buffer` 比 `+` 更高效，因为它避免了频繁的内存分配。

```go
import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer
    buffer.WriteString("Hello")
    buffer.WriteString(" ")
    buffer.WriteString("Go")

    result := buffer.String()
    fmt.Println(result) // 输出：Hello Go
}
```

#### 3.2.5 使用 `strings.Builder` （Go 1.10 及之后）

`strings.Builder` 是 Go 1.10 引入的，专门用于高效的字符串拼接，它是比 `bytes.Buffer` 更合适的选择，因为它直接面向字符串操作。

```go
func main() {
    var builder strings.Builder
    builder.WriteString("Hello")
    builder.WriteString(" ")
    builder.WriteString("Go")

    result := builder.String()
    fmt.Println(result) // 输出：Hello Go
}
```

注意一点，由于Go 语言使用行终止符来分隔语句，换句话说，Go 编译器在解析代码时会自动将每一行末尾（除非行尾有明确的连接符号，如 `+`、`,` 等）看作语句结束。这样做的目的是为了避免显式的语句终止符（如 C/C++ 中的分号 `;`）。

正确情况：

```go
str := "Hello, " +
       "World!"
```

错误情况：

```go
str := "Hello, " 
+ "World!" // 这里会报错
```

这样设计有很多好处：

- **提高可读性**：开发者可以通过视觉检查很容易地确定一行代码是否已经结束。
- **减少分号使用**：自动插入分号机制使得代码简化，不需要每次都写分号。

### 3.3 基本数据类型的转换

在 Go 语言中，类型转换是显式的，并且必须通过一种特定的语法来完成。与某些动态类型语言不同，Go 是静态类型语言，因此类型之间的转换需要开发者明确指定，不会自动进行隐式转换。这可以帮助避免潜在的类型不匹配和错误。

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i)  // 将 int 转换为 float64
    var u uint = uint(f)        // 将 float64 转换为 uint

    fmt.Println(i)  // 输出：42
    fmt.Println(f)  // 输出：42.000000
    fmt.Println(u)  // 输出：42
}
```

注意，**类型转换**并不会改变原变量的类型，而是创建了一个新值，该值的类型是你指定转换的目标类型，而原变量本身的数据类型保持不变。

```go
package main

import "fmt"

func main() {
    var i int = 42      // 定义一个 int 类型的变量 i
    var f float64 = float64(i)  // 将 i 的值转换为 float64 类型，赋值给新变量 f

    fmt.Println(i)  // 输出：42，i 的类型仍然是 int
    fmt.Println(f)  // 输出：42.000000，f 是一个新的 float64 类型的变量
}
```

这个例子中，`float64(i)` 的意思是，将 `i` 存储的值 `42` 转换为 `float64` 类型，并将转换后的值赋给 `f` 。`i` 仍然是 `int` 类型。

### 3.4 字符串与基本类型之间的转换

- **字符串转整数**：可以使用 `strconv` 包中的 `Atoi` 函数。
- **整数转字符串**：可以使用 `strconv` 包中的 `Itoa` 函数。

```go
import (
    "fmt"
    "strconv"
)

func main() {
    i, err := strconv.Atoi("123") // 将字符串转换为整数
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(i) // 输出：123
    }

    s := strconv.Itoa(123)  // 将整数转换为字符串
    fmt.Println(s)          // 输出："123"
}
```

也可以使用`fmt.Sprintf("%参数", 表达式)`的形式进行转换



### 3.5 字符串练习

```Go
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}
```

输入comma函数的参数是一个字符串。如果输入字符串的长度小于或等于3的话，则不需要插入逗号分隔符。否则，comma函数将在最后三个字符前的位置将字符串切割为两个子串并插入逗号分隔符，然后通过递归调用自身来得出前面的子串。

**练习** 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

```Go
func nonRecursiveComma(s string) string {
    var buf bytes.Buffer
    for i := 0; i <= len(s)-1; i++ {
        if i > 0 && (len(s)-i)%3 == 0 {
            buf.WriteByte(',')
        }
        buf.WriteByte(s[i])
    }
    return buf.String()
}
```

**练习** 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

```Go
func comma(s string) string {
    var buf bytes.Buffer
    // 如果以 +/- 号开头，设置 start 为 1，将 +/- 号写入 buf
    var start int
    if s[0] == '+' || s[0] == '-' {
        start = 1
        buf.WriteByte(s[0])
    } else {
        start = 0
    }

    // 截取整数部分
    end := strings.Index(s, ".")
    if end == -1 {
        end = len(s)
    }
    intStr := s[start:end]

    // 对整数部分进行','分割
    for i := 0; i <= len(intStr)-1; i++ {
        if i > 0 && (len(intStr)-i)%3 == 0 {
            buf.WriteByte(',')
        }
        buf.WriteByte(intStr[i])
    }

    // 拼接小数点后面的部分
    rest := s[end:]
    for i := 0; i < len(rest); i++ {
        buf.WriteByte(rest[i])
    }

    return buf.String()
}
```

**练习** 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序：

```Go
func isMixedTwoStrings(s1, s2 string) bool {
    len1, len2 := len(s1), len(s2)
    if len1 != len2 {
        return false
    }

    // key是字符，value是该字符的个数
    runeMap1 := make(map[rune]int)
    runeMap2 := make(map[rune]int)

    for _, v := range s1 {
        runeMap1[v]++
    }
    for _, v := range s2 {
        runeMap2[v]++
    }

    for k, _ := range runeMap1 {
        if runeMap1[k] != runeMap2[k] {
            return false
        }
    }
    return true
}
```

# 三、指针

在 Go 语言中，指针是一个非常重要的概念，它指向一个变量的内存地址。通过指针，你可以直接访问和修改存储在内存中的值，而不是值的副本。在 Go 中，指针类型的使用方式与其他语言（如 C/C++）类似，但也有一些 Go 特有的特性。

### 1. Go 中的指针基础

1. 定义指针

   ：

   - 指针类型使用 `` 表示，表示一个指向某个类型的指针。例如：`int` 表示指向 `int` 类型的指针。
   - 使用 `&` 操作符可以获取一个变量的内存地址。
   - 使用 `` 操作符可以解引用一个指针，访问指针指向的值。

### 例子：

```go
import "fmt"

func main() {
    // 定义一个整数变量
    var a int = 42

    // 定义一个指针，指向变量 a 的地址
    var p *int = &a

    // 打印变量 a 的值和地址
    fmt.Println("Value of a:", a)         // 42
    fmt.Println("Address of a:", &a)      // a 的内存地址
    fmt.Println("Pointer p:", p)          // 与 &a 相同的地址
    fmt.Println("Value via pointer p:", *p) // 通过指针解引用，获取 a 的值 42

    // 修改通过指针修改 a 的值
    *p = 100
    fmt.Println("New value of a:", a)     // 100
}
```

### 2. Go 中指针的特性

1. **Go 中不支持指针运算**： 与 C/C++ 不同，Go 不允许对指针进行算术运算。你不能像在 C 中那样通过增加或减少指针的值来访问相邻的内存位置。这是因为 Go 语言希望减少内存管理中的错误，避免低级的指针操作导致的错误行为。

2. **零值指针**： 指针的零值是 `nil`。如果一个指针没有被初始化，它默认指向 `nil`。尝试解引用一个 `nil` 指针会导致运行时错误（`runtime panic`）。

   ```go
   var p *int
   fmt.Println(p)  // 输出: <nil
   ```

3. **指针与函数**： Go 中的函数传参是值传递，意味着当你将一个变量传递给函数时，函数得到的是这个变量的副本。如果希望函数修改传入的变量，你需要传递变量的指针。

#### 例子：

```go
package main

import "fmt"

// 通过指针修改传入的值
func updateValue(val *int) {
    *val = 100
}

func main() {
    x := 42
    fmt.Println("Before:", x) // 输出 42

    updateValue(&x)  // 传入 x 的指针
    fmt.Println("After:", x)  // 输出 100，值被修改了
}
```

1. **new() 函数**： Go 提供了一个内置的 `new()` 函数，它用于分配内存并返回指向该类型的指针。`new()` 函数不初始化内存，它只是分配零值。

   ```go
   var p *int = new(int)
   fmt.Println(*p)  // 输出: 0，p 指向的是一个零值初始化的 int
   *p = 42
   fmt.Println(*p)  // 输出: 42，修改了该指针指向的值
   ```
   
2. **make() 与指针**： 虽然 `new()` 分配内存用于基本类型或结构体，但 `make()` 是专门用于创建并初始化引用类型（如切片、映射和通道）的函数。这些类型本质上也是指针，但它们有更复杂的底层结构，因此 `make()` 会负责内存分配和初始化。

### 3. 指针与结构体

在 Go 中，使用指针可以更高效地处理结构体类型。通过传递结构体指针，你可以避免拷贝整个结构体的数据，只需传递其地址，并通过指针修改结构体的字段。

#### 例子：

```go
package main

import "fmt"

// 定义一个结构体类型
type Person struct {
    name string
    age  int
}

// 修改结构体的函数，使用指针传递
func updatePerson(p *Person) {
    p.name = "Alice"
    p.age = 30
}

func main() {
    // 创建一个结构体实例
    p1 := Person{name: "John", age: 25}

    fmt.Println("Before:", p1) // 输出: {John 25}

    updatePerson(&p1) // 传递结构体指针

    fmt.Println("After:", p1)  // 输出: {Alice 30}
}
```

# 四、函数

## 1. 函数-调用过程

介绍：为了让大家更好的理解函数调用过程, 看两个案例，并画出示意图，这个很重要

1. 传入一个数+1

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/c2892132-f556-4b19-9f1d-9a2554cd25ec/a31a342a-4489-48ed-9d0c-734aa1527e47/image.png)

对上图说明 (1) 在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间 和其它的栈的空间区分开来

(2) 在每个函数对应的栈中，数据空间是独立的，不会混淆 (3) 当一个函数调用完毕(执行完毕)后，程序会销毁这个函数对应的栈空间。 2. 计算两个数,并返回：

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/c2892132-f556-4b19-9f1d-9a2554cd25ec/c0ec86da-46da-4cad-9969-922644381d6b/image.png)

### 1.1 函数递归

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/c2892132-f556-4b19-9f1d-9a2554cd25ec/ed5562aa-86bc-40dc-862f-f08dc521e4f0/image.png)

对上面代码分析的示意图：

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/c2892132-f556-4b19-9f1d-9a2554cd25ec/eade0c21-11b6-4ae7-9ba1-a2919538240d/image.png)

**函数递归需要遵守的重要原则**:

1. 执行一个函数时，就创建一个新的受保护的独立空间(新函数栈)
2. 函数的局部变量是独立的，不会相互影响
3. 递归必须向退出递归的条件逼近，否则就是无限递归，死龟了:)
4. 当一个函数执行完毕，或者遇到 return，就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或者返回时，该函数本身也会被系统销毁

### 1.2 函数使用的注意事项

1. **基础数据类型**和**数组**在传入函数时，默认情况下是**值传递**。

这意味着函数会接收到原始数据的一个副本，而不是对原始数据的引用。因此，在函数内部对这些数据的任何修改都不会影响到外部的原始数据。

````go
package main

import "fmt"

func modifyArray(arr [3]int) {
    arr[0] = 100
    fmt.Println("Inside function: ", arr)  // 输出 [100 2 3]
}

func main() {
    a := [3]int{1, 2, 3}
    modifyArray(a)
    fmt.Println("Outside function: ", a)  // 输出 [1 2 3]，原始数组未被修改
}
````

1. 如果希望函数内的变量能修改函数外的变量(指的是默认以值传递的方式的数据类型)，可以传 入**变量的地址&**，函数内以指针的方式操作变量。从效果上看类似引用 。
2. **Go 语言不支持函数重载。函数重载**即多个函数可以有相同的名称但不同的参数列表或返回类型。
3. 在 Go 中，**函数也是一种数据类型**，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该**变量可以对函数调**用：

- 将函数赋值给变量：

```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    // 将函数赋值给变量
    var sumFunc func(int, int) int
    sumFunc = add

    // 通过变量调用函数
    result := sumFunc(3, 4)
    fmt.Println("Result:", result)  // 输出 7
}
```

- 函数作为参数传递，或者作为函数的返回值

```go
package main

import "fmt"

// 函数作为参数
func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func add(a, b int) int {
    return a + b
}

func main() {
    result := applyOperation(3, 4, add)
    fmt.Println("Result:", result)  // 输出 7
}
package main

import "fmt"

// 函数返回一个函数
func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    multiplyBy2 := createMultiplier(2)
    multiplyBy3 := createMultiplier(3)

    fmt.Println(multiplyBy2(5))  // 输出 10
    fmt.Println(multiplyBy3(5))  // 输出 15
}
```

1. Go 支持可变参数 在函数定义中，通过在最后一个参数的类型之前添加 `...`，表示这个参数是可变参数。可变参数在函数内部表现为一个切片（`slice`），类型是传入**参数类型的切片**。

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2))           // 输出 3
    fmt.Println(sum(1, 2, 3, 4, 5))  // 输出 15
}
```



## 2. init函数

**每一个源文件都可以包含一个** **init** **函数**，该函数会在 main 函数执行前，被 Go 运行框架调用，也就是说 init() 会在 main 函数前被调用。

### init函数的特点

- **自动调用**：`init` 函数会在包被导入时自动执行，开发者不需要显式调用它。

- **顺序执行**：如果一个文件同时包含**全局变量定义**，**init函数**和 **main** **函数**，则执行的流程**全局变量初始哈**->**执行init函数**->执行**main函数**；当一个包被多个其他包导入时，`init` 函数会根据依赖关系先后执行，先执行被导入的包：

```go
package main

import "fmt"

// 全局变量定义
var globalVar = initializeGlobal()

// 用于初始化全局变量的函数
func initializeGlobal() int {
    fmt.Println("Initializing global variable")
    return 100
}

// init 函数
func init() {
    fmt.Println("Executing init function")
}

// main 函数
func main() {
    fmt.Println("Executing main function")
    fmt.Println("Global variable:", globalVar)
}
```

执行顺序输出：

```go
Initializing global variable
Executing init function
Executing main function
Global variable: 100
```

- **每个包可以有多个 `init` 函数**：同一个包可以在不同的文件中定义多个 `init` 函数，它们会按顺序执行。

- **不能被显式调用**：`init` 函数不能被其他函数或代码直接调用，且不能有任何参数和返回值。

- **与 `main` 函数的关系**：`init` 函数会在 `main` 函数之前执行。`main` 是程序的入口，而 `init` 是在初始化时自动执行的函数。



## 3. 匿名函数

介绍：Go 支持匿名函数，匿名函数就是没有名字的函数，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用。

#### 3.1 立即执行的匿名函数

匿名函数可以定义后立即调用，这种方式非常适合临时执行一次性操作。

```
package main

import "fmt"

func main() {
    // 定义并立即调用匿名函数
    result := func(a, b int) int {
        return a + b
    }(3, 4) // 传递参数3和4
    fmt.Println("Result:", result)  // 输出：Result: 7
}
```

在这个示例中，匿名函数 `(a, b int) int` 被定义后立即执行，通过 `()` 调用并传递了参数 `3` 和 `4`。

#### 3.2 将匿名函数赋值给变量

匿名函数可以赋值给变量，这使得它可以多次调用，类似于具名函数。

```
package main

import "fmt"

func main() {
    // 将匿名函数赋值给变量
    add := func(a, b int) int {
        return a + b
    }

    // 调用匿名函数
    fmt.Println(add(5, 6))  // 输出：11
    fmt.Println(add(10, 20)) // 输出：30
}
```

在这个例子中，匿名函数被赋值给变量 `add`，并可以通过该变量多次调用。

#### 3.3 匿名函数作为参数传递

匿名函数可以作为参数传递给其他函数。这个特性在回调函数、排序、过滤等操作中非常有用。

```
package main

import "fmt"

// 接受函数作为参数
func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    // 传递匿名函数作为参数
    result := applyOperation(5, 3, func(x, y int) int {
        return x - y
    })
    fmt.Println("Result:", result)  // 输出：Result: 2
}
```

在这个示例中，匿名函数 `func(x, y int) int` 被传递给 `applyOperation` 函数，并作为参数在函数内部执行。

#### 3.4 匿名函数作为返回值

匿名函数也可以作为另一个函数的返回值，从而生成不同的函数行为。这个特性常用于函数生成器或工厂模式。

```
package main

import "fmt"

// 返回一个匿名函数
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    mulBy2 := multiplier(2)  // 返回一个将数值乘以2的函数
    mulBy3 := multiplier(3)  // 返回一个将数值乘以3的函数

    fmt.Println(mulBy2(5))   // 输出：10
    fmt.Println(mulBy3(5))   // 输出：15
}
```

在这个例子中，`multiplier` 函数返回了一个匿名函数，该匿名函数捕获了外部作用域的 `factor` 参数，从而形成了一个**闭包**。

#### 3.5 闭包

介绍：闭包就是**一个函数**和与**其相关的引用环境**组合的**一个整体**(实体)

```go
package main

import "fmt"

// 生成一个计数器闭包
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter() // 返回一个匿名函数

	fmt.Println(c()) // 输出：1
	fmt.Println(c()) // 输出：2
	fmt.Println(c()) // 输出：3
}
```

对上面代码的说明和总结

1. counter()是一个函数，返回的数据类型是 fun () int
2. 闭包的说明：

​		返回的是一个匿名函数, 但是这个匿名函数引用到函数外的 count ,因此这个匿名函数就和 count 形成一个整体，构成**闭包**。

3. 可以这样理解:  闭包是类, 函数是操作，count 是字段。函数和它使用到 count 构成闭包。
4. 当我们反复的调用函数`c`时，这个闭包引用了外部作用域的变量（即这里的 `count`），并且外部作用域的变量只初始化一次，闭包会持续访问和修改这个变量的值，也就是说闭包可以保持某些状态或上下文(在这里上下文是count)。
5. 我们要搞清楚闭包的关键，就是要分析出**返回的函数**和**它使用(引用)到哪些变量**，因为函数和它引用到的变量共同构成闭包。

##### 闭包的例题：

```go
package main

import (
	"fmt"
	"strings"
)

/*
闭包的最佳实践

请编写一个程序，具体要求如下：

编写一个函数 makeSuffix(suffix string)，可以接收一个文件后缀名（比如 .jpg），并返回一个闭包。
调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如 .jpg），则返回 文件名.jpg，如果已经有 .jpg 后缀，则返回原文件名。
要求使用闭包的方式完成。
使用 strings.HasSuffix
*/
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 判断文件名是否以指定后缀结尾
		if !strings.HasSuffix(name, suffix) {
			// 如果没有指定后缀，添加后缀
			return name + suffix
		}
        // 有后缀，直接返回源文件名
		return name
	}
}

func main() {
	// 定义一个函数闭包
	addJpgSuffix := makeSuffix(".jpg")
	fmt.Println(addJpgSuffix("file"))      // 输出 file.jpg
	fmt.Println(addJpgSuffix("image.jpg")) // 输出 image.jpg
}
```

说明：定义的`addJpgSuffix`是函数闭包，也就是`addJpgSuffix`返回的匿名函数和外部的`suffix`构成了函数闭包，也就是返回的匿名函数`addJpgSuffix`引用到了外部变量`suffix`。



#### 3.6 defer

##### 为什么需要defer

在函数中，程序员经常需要创建资源(比如：数据库连接、文件句柄、锁等) ，为了在**函数执行完毕后，及时的释放资源**，Go 的设计者提供 defer (延时机制)。

##### defer示例

```go
package main

import (
    "fmt"
)

func sum(n1 int, n2 int) int {
    // 当执行到 defer 时，暂时不执行，会将 defer 后面的语句压入到独立的栈 (defer 栈)
    // 当函数执行完毕后，再从 defer 栈按照先入后出的方式出栈，执行
    defer fmt.Println("ok1 n1=", n1)  // defer
    defer fmt.Println("ok2 n2=", n2)  // defer

    res := n1 + n2
    fmt.Println("ok3 res=", res)

    return res
}

func main() {
    res := sum(10, 20)
    fmt.Println("res=", res)
}
```

程序执行结果，打印的顺序：

```go
ok3 res= 30
ok2 n2= 20
ok1 n1= 10
res= 30
```

##### defer的细节说明

1. 当 go 执行到一个 defer 时，不会立即执行 defer 后的语句，而是将 defer 后的语句压入到一个栈中[我为了讲课方便，暂时称该栈为 defer 栈], 然后继续执行函数下一个语句。
2. 当函数执行完毕后，在从 defer 栈中，依次从栈顶取出语句执行(注：遵守栈 先入后出的机制)。
3. 在 defer 将语句放入到栈时，**同时也会将相关的值拷贝入栈**。

```go
package main

import (
    "fmt"
)

func sum(n1 int, n2 int) int {
    // 当执行到 defer 时，暂时不执行，会将 defer 后面的语句压入到独立的栈 (defer 栈)
    // 当函数执行完毕后，再从 defer 栈，按照先入后出的方式出栈，执行
    defer fmt.Println("ok1 n1=", n1)  // defer 3. ok1 n1 = 10
    defer fmt.Println("ok2 n2=", n2)  // defer 2. ok2 n2 = 20

    // 增加一句话
    n1++ // n1 = 11
    n2++ // n2 = 21

    res := n1 + n2 // res = 32
    fmt.Println("ok3 res=", res)  // 1. ok3 res= 32

    return res
}

func main() {
    res := sum(10, 20)
    fmt.Println("res=", res)  // 4. res= 30
}
```

在 `sum` 函数中：

- `defer fmt.Println("ok1 n1=", n1)` 和 `defer fmt.Println("ok2 n2=", n2)` 被压入栈中。
- 虽然 `n1` 和 `n2` 在函数体中自增了，但 `defer` 中引用的 `n1` 和 `n2` 的值是在 `defer` 语句**执行时**捕获的（即 `n1=10` 和 `n2=20`），因此打印时还是捕获时的值。
- `n1++` 和 `n2++` 之后，新的值 `n1=11` 和 `n2=21` 用于计算 `res`。
- 输出 `ok3 res=32`，返回 `res=32`。

执行顺序：

```go
ok3 res=32（计算结果为 32）
ok2 n2=20（defer 中的 n2 被捕获时的值）
ok1 n1=10（defer 中的 n1 被捕获时的值）
res=30
```

##### defer最佳实践

​	defer 最主要的价值是在，当函数执行完毕后，可以及时的释放函数创建的资源。看下**模拟代码**:

1. 示例:关闭文件资源：

```go
func test() {
    // 关闭文件资源
    file = openfile(文件名)
    defer file.close()
    // 其他代码
}
```

2. 示例:释放数据库资源：

```go
func test() {
    // 释放数据库资源
    connect = openDatabase()
    defer connect.close()
    // 其他代码
}
```

- 在golang编程中的通常做法是，创建资源后，比如(打开了文件，获取了数据库的链接，或者是锁资源)，可以执行deferfile.Close()deferconnect.Close()
- 在defer后，可以继续使用创建资源.
- 当函数完毕后，系统会依次从defer栈中，取出语句，关闭资源.
- 这种机制，非常简洁，程序员不用再为在什么时机关闭资源而烦心。



## 4. 函数参数传递形式

​		我们在讲解函数注意事项和使用细节时，已经讲过值类型和引用类型了，这里我们再系统总结一下，因为这是重难点，值类型参数默认就是值传递，而引用类型参数默认就是引用传递。

### 两种传递形式

1) 值传递

2) 引用传递

其实，不管是值传递还是引用传递，传递给函数的都是变量的副本，不同的是，值传递的是值的拷贝，引用传递的是**地址的拷贝**，一般来说，地址拷贝效率高，因为数据量小，而值拷贝决定拷贝的数据大小，数据越大，效率越低。

## 5. 包 package

在Golang（Go语言）中，**包（package）** 是代码组织和模块化的基本单元，包的设计简化了代码复用和依赖管理，同时提供了封装和命名空间的概念。

### 5.1 包的组织结构

Go中的每个源文件都属于某个包。包名通常与文件所在的目录名称一致，所有文件必须在第一行用 `package` 关键字声明它所属的包。

**包的结构**

- **主程序包**：`package main` 用于声明一个可独立运行的程序，包含 `main` 函数作为程序入口。
- **库包**：其他任何不以 `package main` 开头的包称为库包，通常提供被其他包复用的功能。

```go
package main  // 主程序包

func main() {
    // 主函数
}
package mylib  // 自定义的库包

func MyFunction() {
    // 库包中的函数
}
```

### 5.2 **包的导入机制**

Go使用 `import` 关键字来导入包，包名通常是其路径（相对于 `GOPATH` 或 `go mod` 的模块路径）。导入包后，可以通过包名访问其中的公开标识符（如函数、变量、类型等）。

#### 导入标准库

标准库直接通过包名导入：

```go
import "fmt"
```

#### 导入自定义包或第三方包

Go的导入路径基于工作区的设置：

- 如果使用 `GOPATH`，导入路径基于 `src` 文件夹。
- 如果使用 `Go Modules`，则基于模块的路径和版本。

```go
import "mypackage"
```

### 5.3 **包的初始化**

Go中的包在**首次被导入时**会自动进行初始化，初始化的顺序遵循以下规则：

- 每个包只会被初始化一次。
- 包的初始化顺序是按照包的依赖关系进行的，最底层的依赖包会先被初始化。
- 包的初始化通过：
  - **全局变量的初始化**：包级别的变量在包初始化时首先会被赋值。
  - **`init`函数**：每个包可以有一个或多个 `init` 函数，用于执行初始化逻辑，`init` 函数在全局变量初始化之后自动调用。

例如：

```go
package mylib

var counter int = 0  // 全局变量初始化

func init() {
    counter = 10  // init函数中的初始化逻辑
}
```

- `init` 函数是隐式调用的，不能从代码中显式调用。
- 如果一个包导入了其他包，则会按照依赖关系优先初始化其他包。

**初始化顺序：**

包的初始化遵循以下步骤：

1. **包级别变量初始化**：所有全局变量会被先赋初值。
2. **`init` 函数调用**：在变量初始化完成后，调用 `init` 函数。

### 5.4 **包的可见性**

Go使用标识符的命名规则来控制包中的标识符的可见性（即封装性）。这与传统OOP语言中的访问控制符（如`public`, `private`）不同。

- **导出的标识符**：如果标识符（如函数、变量、类型、常量等）的首字母大写，则它是导出的，即其他包可以访问它。这类似于其他语言中的 `public` 访问修饰符。

  ```go
  package mylib
  
  var ExportedVar = 10  // 导出的变量，可以在其他包中访问
  
  func ExportedFunc() {  // 导出的函数，可以在其他包中调用
      fmt.Println("This is an exported function")
  }
  ```

- **未导出的标识符**：如果标识符的首字母小写，则它是未导出的，只能在定义它的包内访问。这类似于其他语言中的 `private` 访问修饰符。

  ```go
  package mylib
  
  var unexportedVar = 5  // 未导出的变量，仅在本包内可见
  
  func unexportedFunc() {  // 未导出的函数，仅在本包内可见
      fmt.Println("This is an unexported function")
  }
  ```

通过这种大写/小写规则，Go实现了简单而有效的封装。

### 5.5 **包的导入别名与匿名导入**

Go允许在导入包时使用别名或进行匿名导入。

#### 导入别名

如果导入的包名称与当前包中的某些名称冲突，或为了方便，可以给包起一个别名：

```go
import f "fmt"  // 给 fmt 包起别名 f

func main() {
    f.Println("Hello, World!")
}
```

#### 匿名导入

有时候，我们需要导入一个包但不直接使用它的标识符，仅为了执行它的 `init` 函数，这时可以使用匿名导入：

```go
import _ "mypackage"  // 仅导入并执行 mypackage 的 init 函数
```

### 5.6 **包的循环依赖**

Go **不允许包之间存在循环依赖**。如果包A导入了包B，同时包B也试图导入包A，这将导致编译失败。Go强制包之间的依赖是单向的，这样可以简化依赖关系，并确保初始化顺序不会混乱。

如果遇到循环依赖，通常可以通过将公共代码提取到一个单独的包中来解决。

## 6. GO中与字符串操作相关的常用函数

1. 统计字符串的长度，按字节 `len(str)`，对于字符串，返回的是该字符串的字节个数。

```go
func len(v Type) int
```

2. 字符串遍历，同时处理有中文的问题 `r := []rune(str)`
3. 字符串转整数:  `n, err := strconv.Atoi("12")`，函数`strconv.Atoi(str)`其实是`strconv.Atoi(str, 10, 0)`，不过`strconv.Atoi`返回的类型是`(int, error)`

```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
/* 
返回字符串表示的整数值，接受正负号。
base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
*/ 
```

4. 整数转字符串 `str = strconv.Itoa(12345)`，`strconv.Itoa` 将整数转换为对应的字符串形式。
5. 字符串 转 `[]byte`: `var bytes = []byte("hello go")`
6. `[]byte` 转 字符串: `str = string([]byte{97, 98, 99})`
7. 10 进制转 2, 8, 16 进制: `str = strconv.FormatInt(123, 2)` // 2-> 8 , 16

```go
package strconv
func FormatInt(i int64, base int) string
// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
```

8. 查找子串是否在指定的字符串中: `strings.Contains("seafood", "foo") //true`

```go
func Contains(s, substr string) bool
// 判断字符串s是否包含子串substr。
```

9. 统计一个字符串有几个指定的子串 ： `strings.Count("ceheese", "e") //4`

```go
func Count(s, sep string) int
// 返回字符串s中有几个不重复的sep子串。
```

```go
func main() {
	str := "ababa"
	substr := "aba"
    // 这里的不重复意味着子串在主串不共用字符
	fmt.Println(strings.Count(str, substr)) // 输出 1
}
```

10. 不区分大小写的字符串比较(==是区分字母大小写的): `fmt.Println(strings.EqualFold("abc", "Abc")) // true`

11. 返回子串在字符串第一次出现的 index 值，返回的index指的是第几个字节开始的位置，如果没有返回-1 : `strings.Index("NLT_abc", "abc") // 4`

```go
func Index(s, sep string) int
// 子串sep在字符串s中第一次出现的位置，不存在则返回-1。
```

12. 返回子串在字符串最后一次出现的 index，返回的index指的是第几个字节开始的位置，如没有返回-1 : `strings.LastIndex("go golang", "go")`
13. 将指定的子串替换成 另外一个子串: `strings.Replace("go go hello", "go", "go 语言", n)` n 可以指定你希望替换几个，如果 n=-1 表示全部替换:

```go
func Replace(s, old, new string, n int) string
// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
```

14. 按 照 指 定 的 某 个 字 符 ， 为 分 割 标 识 ， 将 一 个 字 符 串 拆 分 成 **字 符 串 数 组** ：`strings.Split("hello,wrold,ok", ",")`

```go
func main() {
	strArr := strings.Split("I love you", " ")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("strArr[%d]: %s\n", i, strArr[i])
	}
	fmt.Printf("strArr： %v", strArr)
}
/* 
strArr[0]: I
strArr[1]: love
strArr[2]: you
strArr： [I love you]
*/
```

14. 将字符串的字母进行大小写的转换:

```go	
strings.ToLower("Go") // go 
strings.ToUpper("go") // GO
```

15. 将字符串左右两边的空格去掉： `strings.TrimSpace(" tn a lone gopher ntrn     ")`
16. 将字符串左右两边指定的字符去掉 ：`strings.Trim("! hello! ", " !") // ["hello"] //将左右两边 !和 " "去掉`
17. 将字符串左边指定的字符去掉 ： `strings.TrimLeft("! hello! ", " !") // ["hello"] //将左边 ! 和 " "去掉`
18. 将字符串右边指定的字符去掉 ：`strings.TrimRight("! hello! ", " !") // ["hello"] //将右边 ! 和 " "去掉`
19. 判断字符串是否以指定的字符串开头: `strings.HasPrefix("ftp://192.168.10.1", "ftp") // true`

```go
package strings
func HasPrefix(s, prefix string) bool
// 判断s是否有前缀字符串prefix。

func HasSuffix(s, suffix string) bool
// 判断s是否有后缀字符串suffix。
```

## 7. GO中日期和时间相关的函数

### 7.1 格式化输出time.Time

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    // 完整日期和时间
    fmt.Println("完整日期和时间:", t.Format("2006-01-02 15:04:05"))
    // 仅日期
    fmt.Println("仅日期:", t.Format("2006-01-02"))
    // 仅时间
    fmt.Println("仅时间:", t.Format("15:04:05"))
    // 美国格式
    fmt.Println("美国格式日期:", t.Format("01/02/2006"))
    // 自定义格式 (24小时制时间)
    fmt.Println("自定义格式:", t.Format("2006年01月02日 15:04"))
}

```

### 7.2 基本时间常量

```go
type Duration int64
// Duration类型代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约290年。
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

### 7.3 time的 `Unix` 和 `UnixNano` 方法

```go
func (t Time) Unix() int64
// Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）。

func (t Time) UnixNano() int64
// UnixNano将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）。如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话，结果是未定义的。
```

```go
func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
}
/*
1731133596
1731133596049038800
*/
```

## 8. 内置函数

Golang 设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称为 Go 的内置函数。文档：https://studygolang.com/pkgdoc -> builtin

1) `len`：用来求长度，比如 string、array、slice、map、channel

2) `new`：用来分配内存，主要用来分配值类型，比如 `int`、`float32`、`struct`...返回的是指针

**举例说明** **new** **的使用**：

```GO
func new(Type) *Type
// 内建函数new分配内存。其第一个实参为类型，而非值。其返回值为指向该类型的新分配的零值的指针。
```

3) `make`：用来**分配内存**，主要用来**分配引用类型**，比如 channel、map、slice。这个我们后面讲解。

````GO
func make(Type, size IntegerType) Type
// 内建函数make分配并初始化一个类型为切片、映射、或通道的对象。其第一个实参为类型，而非值。make的返回类型与其参数相同，而非指向它的指针。其具体结果取决于具体的类型：
```
切片：size指定了其长度。该切片的容量等于其长度。切片支持第二个整数实参可用来指定不同的容量；
     它必须不小于其长度，因此 make([]int, 0, 10) 会分配一个长度为0，容量为10的切片。
映射：初始分配的创建取决于size，但产生的映射长度为0。size可以省略，这种情况下就会分配一个
     小的起始大小。
通道：通道的缓存根据指定的缓存容量初始化。若 size为零或被省略，该信道即为无缓存的。
```
````

## 9. 错误处理

1) 在默认情况下，当发生错误后(panic) ,程序就会退出（崩溃.）

2) 如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。还可以在捕获到错误后，给管理员一个提示(邮件,短信。。。）

3) 这里引出我们要将的错误处理机制。

### 9.1 错误处理基本说明

1) Go 语言追求简洁优雅，所以，Go 语言不支持传统的 try…catch…finally 这种处理。

2) Go 中引入的处理方式为：**defer**, **panic**, **recover**

3) 这几个异常的使用场景可以这么简单描述：Go 中可以抛出一个 panic 的异常，然后在 defer 中通过 recover 捕获这个异常，然后正常处理

### 9.2 使用 defer + recover 来处理错误

当你使用 `defer` 和 `recover` 时，可以在函数中捕获并处理 `panic`，避免程序因未捕获的错误而崩溃。下面是代码的工作原理：

1. 当函数中的代码执行时，如果没有错误，`defer` 注册的代码块在函数返回时执行。
2. 如果在函数中遇到 `panic`，程序会跳出正常的执行流程，执行 `defer` 注册的代码块。
3. 在 `defer` 中调用 `recover` 会捕获 `panic` 的错误信息，使程序恢复正常流程，**避免崩溃**。

```go
package main

import "fmt"

func main() {
	test()
	// code ...
	fmt.Println("code...")
}

func test() {
	// 使用defer + recover 来处理函数中遇到的错误
	defer func() {
		err := recover()
		if err != nil { // 说明在函数执行过程中遇到了错误
			fmt.Printf("err=%v\n", err)
			// 这里可以针对错误进行各种各样的处理，或者是将错误信息发送给管理员等等
		}
	}()
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	fmt.Println(num3)
}
```

### 9.3 自定义错误

Go 程序中，也支持自定义错误， 使用 errors.New 和 panic 内置函数。

1) `errors.New("错误说明")` , 会返回一个 error 类型的值，表示一个错误

2) panic 内置函数 ,接收一个` interface{}`类型的值作为参数。可以接收 error 类型的变量，**输出错误信息**，**并退出程序**

```go
func New(text string) error
// 使用字符串创建一个错误,请类比fmt包的Errorf方法，差不多可以认为是New(fmt.Sprintf(...))。
```



# 五、 数组和切片

## 1. 初始化数组的方式

### 1.1 使用字面量初始化数组

可以在定义数组的同时使用字面量赋值进行初始化：

```go
// 定义一个长度为 5 的整型数组，并指定值
var arr1 = [5]int{1, 2, 3, 4, 5}

// 使用简短声明方式
arr2 := [5]int{1, 2, 3, 4, 5}
```

### 1.2 使用 `[...]` 自动推断数组长度

可以使用 `[...]` 来让编译器自动推断数组长度：

```go
arr := [...]int{1, 2, 3, 4, 5} // 长度会自动推断为 5
```

### 1.3 定义并初始化特定索引的元素

可以只为特定索引初始化值，其余索引会自动初始化为元素类型的零值：

```go
arr := [5]int{0: 10, 3: 30} // 初始化索引 0 和 3，其余为 0
// 结果为 [10, 0, 0, 30, 0]
```

### 1.4 创建空数组并手动赋值

先定义一个空数组，然后逐个元素赋值：

```go
var arr [5]int // 数组元素默认为 0
arr[0] = 1
arr[1] = 2
```

### 1.5 多维数组初始化

可以定义多维数组，并在声明时为每个维度赋值：

```go
// 2x3 的二维数组
arr := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

// 或者
arr := [...][3]int{
    {1, 2, 3},
    {4, 5, 6},
} // 使用 [...] 自动推断第一维长度
```

## 2. 数组使用的注意事项

1. 数组是多个相同类型数据的组合,一个数组一旦声明/定义了,其**长度是固定的**, **不能动态变化**：

![image-20241109154109497](image\image-20241109154109497.png)

2. `var arr []int `

   这时 arr 就是一个 slice 切片.

3.  Go 的数组属值类型， 在默认情况下是值传递， 因此会进行值拷贝。数组间不会相互影响:

![image-20241109154627472](image\image-20241109154627472.png)

4. 长度是数组类型的一部分，在传递函数参数时 需要考虑数组的长度，看下面案例:

![image-20241109155042558](image\image-20241109155042558.png)

## 3. 切片

先看一个需求：我们需要一个数组用于保存学生的成绩，但是**学生的个数**是不确定的，请问怎么办？解决方案：-》使用**切片**。切片的基本介绍：

1) 切片的英文是 slice

2) 切片是数组的一个引用，因此**切片是引用类型**，在进行传递时，遵守引用传递的机制。

3) 切片的**使用和数组类似**，遍历切片、访问切片的元素和求切片长度 len(slice)都一样。

4) 切片的长度是可以变化的，因此切片是一个**可以动态变化数组**。

5) 切片定义的基本语法:

var 切片名 []类型

比如：`var a [] int`

### 3.1 切片在内存中的形式（重要）

为了让大家更加深入的理解切片，我们画图分析一下切片在内存中是如何布局的，这个是一个非常重要的知识点：(以前面的案例来分析)

画出前面的切片内存布局:

![image-20241109170835941](image\image-20241109170835941.png)
