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
    "webCal/http"
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

对上图说明 (1) 在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间 和其它的栈的空间区分开来

(2) 在每个函数对应的栈中，数据空间是独立的，不会混淆 (3) 当一个函数调用完毕(执行完毕)后，程序会销毁这个函数对应的栈空间。 2. 计算两个数,并返回

### 1.1 函数递归

对上面代码分析的示意图：

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

在 Go 语言内部，切片的结构大致是如下这样（简化表示）：

```go
type Slice struct {
    ptr *array // 指向底层数组的指针
    len int    // 当前切片的长度
    cap int    // 当前切片的容量
}
```

```go
func main() {
    intArr := [5]int{1, 22, 33, 44, 55}
	intSlice := intArr[2:4]
	for i := 0; i < 5; i++ {
		fmt.Printf("intArr[%d]地址:%v\n", i, &intArr[i])
	}
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("intSlice[%d]地址:%v\n", i, &intSlice[i])
	}
	fmt.Println(cap(intSlice))
	fmt.Println(len(intSlice))

	intSlice[0] = 2
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\t", intArr[i])
	}
	fmt.Println()
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("%d\t", intSlice[i])
	}
}
/* 输出
intArr[0]地址:0xc00000e330
intArr[1]地址:0xc00000e338
intArr[2]地址:0xc00000e340
intArr[3]地址:0xc00000e348
intArr[4]地址:0xc00000e350
intSlice[0]地址:0xc00000e340
intSlice[1]地址:0xc00000e348
3
2
1       22      2       44      55
2       44
*/
```

Slice是引用类型，所以对Slice中的元素进行修改会引起其引用的底层数组的元素的修改。

### 3.2 切片的引用

1. 方式 1

第一种方式：定义一个切片，然后让切片去引用一个已经创建好的数组：

样的。

2. 方式 2

第二种方式：通过 **make** 来创建切片. 

基本语法：**`var 切片名` `[]type = make([]type, len, [cap])`**

参数说明: type: 就是数据类型 len : 大小 cap ：指定切片容量，**可选， 如果你分配了`cap`,则要求`cap>=len`.** 

```go
var floatSlice []float64 = make([]float64, 5, 10)
```

直接声明切片和使用`make`声明切片是由一些区别的

**`var floatSlice []float64`**

- 这是定义了一个名为 `floatSlice` 的空切片，类型为 `[]float64`。
- 此时 `floatSlice` 为 `nil`，没有分配任何内存，也没有初始容量或长度。
- 直接对它进行访问或操作会报错，需要先使用 `make()` 函数来分配内存。

**`var floatSlice []float64 = make([]float64, 5, 10)`**

- 使用 `make([]float64, 5, 10)` 创建了一个 `[]float64` 类型的切片并赋值给 `floatSlice`。
- 该切片的**初始长度**为 `5`，**容量**为 `10`。
- 这意味着可以直接对 `floatSlice` 的前五个元素进行访问或赋值，而不需要额外的 `make()` 操作。
- 如果追加元素超出初始长度 `5`，但不超过容量 `10`，Go 将自动扩展其长度。
- 一旦切片的元素数量超过容量 `10`，Go 会重新分配更大的底层数组。

3. 方式 3

第 3 种方式：定义一个切片，直接就指定具体数组，使用原理类似 make 的方式

```go
var strSlice []string = []string{"tom", "jack", "mary"}
```



### 3.3 切片的使用细节

#### 3.3.1 切片的cap大小和切片引用的底层数组长度大小的关系

- **容量与底层数组的长度关系**：切片的容量总是与它的底层数组的长度相等，或者小于底层数组的总长度。也就是说，切片容量和底层数组的容量是绑定在一起的。
- **切片容量不足时扩展**：当通过 `append` 函数向切片添加元素时，如果切片的容量不足以容纳新的元素，Go 会为切片分配一个新的、更大的底层数组。这个新数组的大小通常是原容量的两倍（具体增长规则与实现相关）。

#### 3.3.2 切片扩容的底层细节

在 Go 中，当切片的长度超过其当前容量时，`append` 函数会创建一个新的底层数组来存储切片的数据，并将切片的引用指向这个新数组，而不再引用原来的数组。这种行为背后的原理和实现细节如下：

1. **切片的结构**：切片在 Go 语言内部是一个结构体，包含三个字段：

   - `指针`：指向底层数组的起始位置。
   - `长度`：切片当前持有的元素数量。
   - `容量`：从切片的起始位置到底层数组末尾的最大可用元素数量。

   当 `append` 操作导致切片超出其容量时，Go 会自动创建一个新的底层数组，以容纳新增的元素。

2. **扩展容量的策略**：当容量不足时，Go 会按照一定的策略来扩展容量。具体的扩展策略因实现不同有所差异，但通常来说，当切片容量较小时，Go 会选择加倍扩展；当切片容量较大时，Go 会以更小的增量扩展，以避免内存浪费。

3. **生成新的底层数组**：

   - 当 `append` 被调用且超出当前切片容量时，Go 会自动分配一个新的、更大的底层数组。
   - `append` 函数会将原有切片的数据复制到新数组中，并将新增的元素追加到新数组中。
   - `append` 返回一个指向新底层数组的切片，因此此时新的 `floatSlice1` 就不再引用原来的 `floatArray`，而是指向了新分配的数组。

4. **原底层数组不变**：原来的底层数组（`floatArray`）依然保持不变，且仍然包含最初的内容，只是 `floatSlice1` 不再指向它了。这种做法是为了保持 Go 语言切片在内存管理上的安全性和效率。

```go
func main() {
    var floatArray [5]float64
	for i, _ := range floatArray {
		floatArray[i] = float64(i)
	}
	var floatSlice1 []float64 = floatArray[0:5]
	//var floatSlice2 []float64 = make([]float64, 5, 10)
	fmt.Printf("len floatSlice1:%d\n", len(floatSlice1))
	fmt.Printf("cap floatSlice1:%d\n", cap(floatSlice1))
	fmt.Printf("len floatArray:%d\n", len(floatArray))
	floatSlice1 = append(floatSlice1, 5, 6)
	fmt.Printf("after append: len floatSlice1:%d\n", len(floatSlice1))
	fmt.Printf("after append: cap floatSlice1:%d\n", cap(floatSlice1))
	fmt.Printf("after append: len floatArray:%d\n", len(floatArray))
}
/* output：
len floatSlice1:5
cap floatSlice1:5
len floatArray:5
after append: len floatSlice1:7
after append: cap floatSlice1:10
after append: len floatArray:5
*/
```

总结：

切片 append 操作的本质就是对数组扩容;

go 底层会创建一下新的数组 `newArr` (安装扩容后大小);

将 slice 原来包含的元素拷贝到新的数组 `newArr`;

slice 重新引用到` newArr`;

注意 `newArr`是在底层来维护的，程序员不可见;

关于`appand`函数的使用细节可以参考这篇文章：[一篇文章带你看懂Go append方法]https://juejin.cn/post/6951672096699187207

#### 3.2.3 切片的拷贝

- 浅拷贝

**浅拷贝**是指复制切片的结构体信息，但不复制底层数组，因此新旧切片仍然引用同一个底层数组。

1. ​	在 Go 中，可以通过直接赋值实现浅拷贝：

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice1 // 浅拷贝
/*
拷贝后效果：slice1 和 slice2 都引用同一个底层数组。
修改影响：修改 slice2 的元素会影响到 slice1，因为它们引用的是同一片内存空间。
*/
```

- 深拷贝

**深拷贝**是指复制整个切片，包括底层数组，使得新旧切片拥有独立的底层数组。这样，两者之间的修改不会互相影响。

1. 在 Go 中，可以使用内置的 `copy` 函数实现深拷贝：

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := make([]int, len(slice1)) // 创建一个新的切片，用于存放拷贝的数据
copy(slice2, slice1) // 深拷贝
/*
拷贝后效果：slice1 和 slice2 拥有独立的底层数组。
修改影响：修改 slice2 不会影响 slice1，反之亦然。
*/
```

##### `copy` 函数的使用细节

在使用 `copy` 函数进行深拷贝时，有一些需要注意的细节：

- **复制的长度**：`copy` 函数会复制两个切片中较短的那个的长度。例如，如果源切片比目标切片长，`copy` 只会复制目标切片的长度。

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := make([]int, 3) // 长度小于 slice1
copy(slice2, slice1) // 只会复制前 3 个元素

fmt.Println("slice2:", slice2) // 输出：[1, 2, 3]
```

![image-20241111151755812](image\image-20241111151755812.png)



# 六、map

## 1. 有序遍历 map

在 Go 中，`map` 是无序的，因此无法直接保证遍历的顺序。但是，可以通过将 `map` 中的键排序后，再按排序后的顺序遍历 `map`，从而实现有序遍历。以下是具体实现方法。

1. **提取键**：将 `map` 中的所有键提取出来，并存储到一个切片中。
2. **排序键**：对键的切片进行排序。可以使用标准库中的 `sort` 包来排序。
3. **按排序后的键遍历 `map`**：根据排序后的键顺序，按顺序访问 `map` 中的值。

假设我们有一个 `map[string]int`，并希望按键的字母顺序进行遍历：

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 示例 map
    scores := map[string]int{
        "Alice": 90,
        "Bob":   85,
        "Charlie": 92,
        "David":  88,
    }

    // 1. 提取键
    keys := make([]string, 0, len(scores))
    for k := range scores {
        keys = append(keys, k)
    }

    // 2. 对键排序
    sort.Strings(keys)

    // 3. 按排序后的键遍历 map
    for _, k := range keys {
        fmt.Printf("%s: %d\n", k, scores[k])
    }
}
```

```
Alice: 90
Bob: 85
Charlie: 92
David: 88
```

## 2. map 使用细节

1. map 是引用类型，遵守引用类型传递的机制，在一个函数接收 map，修改后，会直接修改原本的 map：

![image-20241111171108167](C:\Users\hasee\GolandProjects\go_core_programming\md\image-20241111171108167.png)

2. map 的容量达到后，再想 map 增加元素，会自动扩容，并不会发生 panic，也就是说 map 能动态的增长键值对(key-value)

# 七、 GO中的面向对象

## 1. `golang` 中面向对象简介

1.  `Golang` 也支持面向对象编程(`OOP`)，但是和传统的面向对象编程有区别，并不是纯粹的面向对象语言。所以我们说 `Golang` **支持面向对象编程特性**是比较准确的。

2. `Golang` 没有类(class)，Go 语言的结构体(`struct`)和其它编程语言的类(class)有同等的地位，你可以理解 `Golang` 是基于 `struct` 来实现 `OOP` 特性的。

3.  `Golang` 面向对象编程非常简洁，去掉了传统 `OOP` 语言的继承、方法重载、构造函数和析构函数、隐藏的 this 指针等等。
4. `Golang` 仍然有面向对象编程的**继承，封装和多态**的特性，只是实现的方式和其它 `OOP` 语言不一样，比如继承 ：`Golang` 没有 `extends` 关键字，继承是通过匿名字段来实现。
5. `Golang` 面向对象(`OOP`)很优雅，`OOP` 本身就是语言类型系统(type system)的一部分，通过接口(`interface`)关联，耦合性低，也非常灵活。后面同学们会充分体会到这个特点。也就是说在 `Golang` 中面向接口编程是非常重要的特性。

## 2. 结构体

### 7.2.1 结构体使用的一些注意事项

1. 结构体进行 type 重新定义(相当于取别名)，`Golang` 认为是新的数据类型，**但是相互间可以强转**:

```go
type Student struct {
    Name string
    Age  int
}

type Stu Student

func main() {
    var stu1 Student
    var stu2 Stu
    stu2 = stu1 // 错误吗？错误，可以这样修改 stu2 = Stu(stu1) // ok
    fmt.Println(stu1, stu2)
}
```

尽管 `Stu` 的结构与 `Student` 相同，但 Go 语言不会将 `Stu` 和 `Student` 视为相同类型，因此不能直接将 `stu1` 赋值给 `stu2`。



2. `struct` 的每个字段上，可以写上一个 **tag**, 该 **tag** 可以通过反射机制获取，常见的使用场景就是**序列化和反序列化**:

- 首先，什么是 `struct` 字段的 tag？

在 Go 结构体的字段声明后，可以写上一段反引号包裹的字符串，这就是 **tag**。`tag` 通常以 `key:"value"` 的形式表示，可以包含多个键值对。

```go
type Person struct {
    Name string `json:"name" xml:"name"`
    Age  int    `json:"age" xml:"age"`
}
```

在这个结构体 `Person` 中，字段 `Name` 和 `Age` 都有 `tag`。例如，`Name` 字段的 tag 是 `json:"name" xml:"name"`，表示在 `JSON` 序列化时将字段 `Name` 映射为 `"name"`，在 XML 序列化时同样映射为 `"name"`。

- `tag` 的作用：序列化和反序列化:

**序列化**是将数据结构（如结构体）转换为特定格式的字符串（如 JSON、XML），而**反序列化**是将格式化的字符串转换回数据结构。`tag` 为 Go 提供了灵活的字段映射方式，可以在序列化和反序列化时指定特定的字段名。

例如，Go 中的 `encoding/json` 包会根据 `tag` 中的 `json:"name"` 来决定 JSON 字段的名称，而不是直接使用 Go 结构体字段的名称。这对控制输出格式、字段重命名以及数据转换非常有用。

```GO
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{Name: "Alice", Age: 30}

    // 序列化为 JSON
    jsonData, _ := json.Marshal(p)
    fmt.Println(string(jsonData)) // 输出：{"name":"Alice","age":30}

    // 反序列化 JSON
    var p2 Person
    json.Unmarshal(jsonData, &p2)
    fmt.Println(p2) // 输出：{Alice 30}
}
```

在这个例子中，`json:"name"` 和 `json:"age"` 告诉 `json.Marshal` 和 `json.Unmarshal` 在 JSON 中使用 `"name"` 和 `"age"` 作为字段名，而不是 `Name` 和 `Age`。这对于 JSON 的格式化和字段控制非常有帮助。

- 使用反射获取 `tag`:

```GO
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string `json:"name" xml:"name"`
    Age  int    `json:"age" xml:"age"`
}

func main() {
    p := Person{Name: "Alice", Age: 30}

    t := reflect.TypeOf(p)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field: %s, JSON tag: %s, XML tag: %s\n", 
                   field.Name, 
                   field.Tag.Get("json"), 
                   field.Tag.Get("xml"))
    }
}
```

```
Field: Name, JSON tag: name, XML tag: name
Field: Age, JSON tag: age, XML tag: age
```

## 3. 方法

在某些情况下，我们要需要声明(定义)方法。比如 Person 结构体:除了有一些字段外( 年龄，姓名..),`Person` 结构体还有一些行为比如:可以说话、跑步..,通过学习，还可以做算术题。这时就要用方法才能完成。

`Golang` 中的方法是**作用在指定的数据类型上的**(即：和指定的数据类型绑定)，因此**自定义类型都可以有方法**，而不仅仅是 **`struct`**。

### 7.3.1 方法的声明和调用

在 Go 中，方法的声明与函数类似，但有一个额外的接收者参数。接收者参数在 `func` 关键字和方法名称之间，用于指定方法所属的类型。方法声明的一般格式如下：

```go
func (receiver Type) MethodName(parameters) returnType {
    // 方法体
}
```

```
receiver：接收者，用于指定方法的所属类型。接收者类型可以是值类型或指针类型。
MethodName：方法的名称。
parameters：方法的参数列表。
returnType：方法的返回类型，可以有多个返回值。
```

#### 值类型接收者的方法

使用值类型作为接收者时，方法只能作用于该类型的副本，而**不会修改原始值**。

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// 值接收者的方法
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}
```

#### 指针类型接收者的方法

使用指针类型作为接收者时，方法可以直接修改原始值，因为指针传递的是原始值的地址。

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// 指针接收者的方法
// UpdateAge 方法的接收者是 *Person（指针类型），因此可以修改 Person 的原始值。
func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}
```

#### 方法与函数的区别

1. 调用方式不一样：

   - **函数**：函数是独立的，可以接受任何类型的参数。

   - **方法**：方法是特定类型的函数，通过接收者关联到特定类型上。


2. 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然；但是对于方法（如 `struct` 的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以:

也就是说方法的调用对比函数的灵活性更强：

- **当接收者是值类型时**，可以使用指针类型的变量调用该方法，Go 会自动解引用指针，将其转换为值类型。

- **当接收者是指针类型时**，可以使用值类型的变量调用该方法，Go 会自动取该值的地址，将其转换为指针类型。

### 7.3.2 方法的调用和传参机制

方法的调用和传参机制和函数基本一样，不一样的地方是方法调用时，**会将调用方法的变量作为接收者（receiver）传递给方法**，即该变量会自动成为方法的一个隐式参数。下面我们举例说明：

```go
package main

import "fmt"

// 定义结构体类型 Person
type Person struct {
    Name string
    Age  int
}

// 定义方法 Greet，接收者是 Person
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    // 创建 Person 类型的实例
    person := Person{Name: "Alice", Age: 25}

    // 调用方法
    person.Greet()
}
```

```tex
接收者参数：在方法 Greet 中，p 是 Person 类型的接收者，这相当于一个特殊的参数。
隐式传递调用者：在 main 函数中，当我们调用 person.Greet() 时，Go 会自动将变量 person 作为实参传递给方法 Greet 的接收者 p。
访问接收者属性：在方法内部，可以通过 p.Name 和 p.Age 来访问 person 的字段。
```

#### 7.3.3 方法使用的注意事项

如果一个类型实现了 `String()` 方法，那么在使用 `fmt.Println`、`fmt.Printf` 或其他 `fmt` 包的打印函数时，会自动调用该类型的 `String()` 方法来获取其字符串表示形式。这是因为 `fmt` 包会检查该类型是否实现了 `Stringer` 接口：

```go
// Stringer 接口定义
type Stringer interface {
    String() string
}
```

```go
package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

// 实现 String 方法
func (p Person) String() string {
    return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Alice", Age: 30}

    // fmt.Println 会调用 String 方法
    fmt.Println(p) // 输出：Person{Name: Alice, Age: 30}

    // println 不会调用 String 方法
    println(p) // 输出的格式不确定，通常是结构体字段的默认值
}
```

### 3. 面向对象部分中的工厂模式

在 Go 语言中，尽管没有传统的面向对象机制（如继承），并且`Golang`的结构体中没有构造函数，但是我们仍然可以实现一些面向对象的设计模式，比如工厂模式。**工厂模式**是一种创建对象的设计模式，它可以根据不同的需求创建并返回对象，而不直接暴露对象的创建细节。

```go
package model

type student struct {
	Name  string
	grade float64 // 假设学生的分数不对外导出
}

func (s student) getScore() float64 {
	return s.grade
}

func NewStudent(name string, grade float64) *student {
	return &student{
		Name:  name,
		grade: grade,
	}
}
```

# 八、 面向对象编程思想



## 1. 面向对象编程思想-抽象

抽象的介绍：我们在前面去定义一个结构体时候，实际上就是把一类事物的共有的**属性**(字段)和**行为**(方法)提取出来，形成一个**物理模型**(结构体)。这种研究问题的方法称为抽象。

举例说明：

![image-20241112160510145](image\image-20241112160510145.png)

```go
package main

import (
	"fmt"
)

// Account 定义一个结构体
type Account struct {
	AccountNo string  // 账户号
	Pwd       string  // 密码
	Balance   float64 // 余额
}

// Deposit 1. 存款
func (account *Account) Deposit(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功~~")
}

// WithDraw 2.取款
func (account *Account) WithDraw(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功~~")
}

// Query 3.查询余额
func (account *Account) Query(pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v 余额=%v \n", account.AccountNo, account.Balance)
}

func main() {
	// 测试一把
	account := Account{
		AccountNo: "gs1111111",
		Pwd:       "666666",
		Balance:   100.0}
	// 这里可以做得更加灵活，就是让用户通过控制台来输入命令... //菜单.... account.Query("666666")
	account.Deposit(200.0, "666666")
	account.Query("666666")
	account.WithDraw(150.0, "666666")
	account.Query("666666")
}
```

## 2. 面向对象编程三大特性-封装

### 8.2.1 基本介绍

`Golang` 仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其它 `OOP` 语言不一样，下面我们一一为同学们进行详细的讲解 `Golang` 的三大特性是如何实现的。

### 8.2.2 封装

在 Go 语言中，封装（Encapsulation）并不像传统面向对象编程（如 Java 或 C++）中那样直接支持 `class` 或 `access modifiers`（如 `private`, `public`）来控制数据和方法的可见性和访问级别。

封装可以**隐藏实现细节**；并且可以提可以对**数据进行验证**，保证安全合理(Age)

### 8.2.2 封装的实例

1. 结构体字段私有化

`accountNo`、`pwd` 和 `balance` 字段都使用了小写字母开头。通过将 `account` 结构体中的所有字段设为私有，外部包和函数无法直接读取或更改这些字段，从而保护了账户信息的完整性和安全性。

2. 提供公开的构造函数来创建对象

提供了一个**工厂函数** `NewAccount`，用来创建 `account` 对象，而不是让外部代码直接创建。这样做的好处是可以在对象创建时，对传入的参数进行合法性检查（如账号长度、密码长度、初始余额等），避免了不合规数据的出现。

```go
package account

import (
	"fmt"
)

// account 定义一个结构体
type account struct {
	accountNo string  // 账户号
	pwd       string  // 密码
	balance   float64 // 余额
}

// NewAccount 工厂模式的函数-构造函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对...")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对...")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数目不对...")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance}
}

// Deposit 1. 存款
func (account *account) Deposit(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.balance += money
	fmt.Println("存款成功~~")
}

// WithDraw 2.取款
func (account *account) WithDraw(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看取款金额是否正确
	if money <= 0 || money > account.balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.balance -= money
	fmt.Println("取款成功~~")
}

// Query 3.查询余额
func (account *account) Query(pwd string) {
	//看下输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v 余额=%v \n", account.accountNo, account.balance)
}
```

```go
package main

import (
	"fmt"
	"go_core_programming/code/chapter_8/account/account"
)

func main() {
	//创建一个 account 变量
	account := account.NewAccount("jzh11111", "123456", 40)
	if account != nil {
		fmt.Println("创建成功=", *account)
	} else {
		fmt.Println("创建失败")
	}
}
```

## 3. 面向对象编程三大特性-继承

```go
package main

import "fmt"

// 编写一个学生考试系统

// Pupil 小学生
type Pupil struct {
	Name  string
	Age   int
	Score int
}

// ShowInfo 显示他的成绩
func (p *Pupil) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Score)
}

func (p *Pupil) SetScore(score int) {
	//业务判断
	p.Score = score
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中.....")
}

// 大学生, 研究生。。

// Graduate 大学生
type Graduate struct {
	Name  string
	Age   int
	Score int
}

// ShowInfo 显示他的成绩
func (p *Graduate) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Score)
}
func (p *Graduate) SetScore(score int) {
	//业务判断
	p.Score = score
}
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}

func main() {
	//测试
	var pupil = &Pupil{
		Name: "tom",
		Age:  10}
	pupil.testing()
	pupil.SetScore(90)
	pupil.ShowInfo()
	//测试
	var graduate = &Graduate{
		Name: "mary",
		Age:  20}
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
}
```

对上面代码的小结

1) Pupil 和 Graduate 两个结构体的字段和方法几乎，但是我们却写了相同的代码， 代码复用性不

强

2) 出现代码冗余，而且代码**不利于维护**，同时**也不利于功能的扩展**。

3) 解决方法-通过**继承**方式来解决。

### 8.3.1 Go 语言的继承：结构体嵌入

Go 中使用结构体的嵌入可以实现类似继承的效果。一个结构体可以嵌入另一个结构体，嵌入的结构体的字段和方法会被“继承”下来，成为外部结构体的一部分。这种方式提供了一种灵活的代码复用和对象行为扩展机制。

```go
package main

import "fmt"

// 编写一个学生考试系统

// Student 学生共有属性
type Student struct {
	Name  string
	Age   int
	Score int
}

// ShowInfo 显示学生的成绩
func (s *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(score int) {
	//业务判断
	s.Score = score
}

// Pupil 小学生
type Pupil struct {
	Student //嵌入了 Student 匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中.....")
}

// Graduate 大学生, 研究生。。
type Graduate struct {
	Student //嵌入了 Student 匿名结构体
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}

func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "mary~"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
}
```

Go 中没有直接的继承机制，但通过嵌入实现了类似继承的效果。`Pupil` 和 `Graduate` 通过嵌入 `Student`，获得了 `Student` 的所有属性和方法，并且还可以定义自己的方法 `testing`，达到继承和扩展的目的。

如果 `Student`接着定义了一个方法，则`Pupil` 和 `Graduate`都可以调用，同时`Pupil` 和 `Graduate` 结构体可以在 `Student` 基础上扩展自己的方法，而无需重新定义已有的功能。

### 8.3.2 继承的深入讨论

1.  结构体可以**使用嵌套匿名结构体所有的字段和方法**，即：首字母大写或者小写的字段、方法，都可以使用。
2. 匿名结构体字段访问可以简化。比如说 8.3.1中`mian`函数的语句可以等价于以下代码：

```
pupil.Name = "Tom"   // 等价于 pupil.Student.Name
```

**当我们直接通过** **`pupil`** **访问字段或方法时，其执行流程如下**：

-  编译器会先看 `pupil` 对应的类型有没有 `Name`, 如果有，则直接调用 `Pupil` 类型的 `Name` 字段

-  如果没有就去看 `Pupil`  中嵌入的匿名结构体 `Student` 有没有声明 `Name` 字段，如果有就调用,如果没有则继续递归查找.如果都找不到就报错。

3. 当**结构体**和**匿名结构体**有相同的字段或者方法时，**编译器采用就近访问原则访问**，如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分：

```go
// Student 学生共有属性
type Student struct {
	Name  string
	Age   int
	Score int
}

// Pupil 小学生
type Pupil struct {
	Student //嵌入了 Student 匿名结构体
	Name    string
}

func main() {
    //当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
    pupil.Name = "jack"
    fmt.Println(pupil.Name)
	fmt.Println(pupil.Student.Name)
}
```

```
jack
tom~
```

4. 结构体嵌入两个(或多个)匿名结构体，如**两个匿名结构体有相同的字段和方法**(**同时结构体本身没有同名的字段和方法**)，在访问时，就必须明确指定匿名结构体名字，否则编译报错。

5. 如果一个 `struct` 嵌套了一个有名结构体，这种模式就是**组合**，如果是组合关系，那么**在访问组合**

   **的结构体的字段或方法时，必须带上结构体的名字**。

```go
package main

import "fmt"

// 基础结构体 Student
type Student struct {
	Name  string
	Score int
}

// Student 的方法 ShowInfo
func (s *Student) ShowInfo() {
	fmt.Printf("Student Name: %v, Score: %v\n", s.Name, s.Score)
}

// 外层结构体 Pupil，包含组合关系
type Pupil struct {
	StudentInfo Student // 使用有名字段 StudentInfo 表示组合关系
	Level       int
}

func main() {
	// 创建 Pupil 实例
	pupil := Pupil{
		StudentInfo: Student{Name: "Tom", Score: 85},
		Level:       3,
	}

	// 访问组合结构体的字段和方法时，必须使用字段名称
	fmt.Println("Pupil Name:", pupil.StudentInfo.Name) // 访问 Student 的 Name 字段
	fmt.Println("Pupil Score:", pupil.StudentInfo.Score) // 访问 Student 的 Score 字段

	// 调用组合结构体的方法时，也必须使用字段名称
	pupil.StudentInfo.ShowInfo() // 调用 Student 的 ShowInfo 方法

	// 访问 Pupil 自己的字段
	fmt.Println("Pupil Level:", pupil.Level)
}
```

```
Pupil Name: Tom
Pupil Score: 85
Student Name: Tom, Score: 85
Pupil Level: 3
```

## 4. 接口

在`golang`中，多态主要是通过接口(`interface`)实现的。接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会表现出它们自己的方法。也就是说当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。

示例：不管是什么类型，也可以是自定义的类型，只要满足满足了这个接口定义的所有方法，那么就可以说这个类型是这个接口，如下我们可以说`*ByteCounter`满足了接口`io.Writer`：

```go
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12, = len("hello, Dolly")
}
```

`fmt`中这几个函数的定义如下：

```go
package fmt

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
func Printf(format string, args ...interface{}) (int, error) {
    return Fprintf(os.Stdout, format, args...)
}
func Sprintf(format string, args ...interface{}) string {
    var buf bytes.Buffer
    Fprintf(&buf, format, args...)
    return buf.String()
}
```

其中接口`io.Writer`的定义如下

```go
package io

// Writer is the interface that wraps the basic Write method.
type Writer interface {
    // Write writes len(p) bytes from p to the underlying data stream.
    // It returns the number of bytes written from p (0 <= n <= len(p))
    // and any error encountered that caused the write to stop early.
    // Write must return a non-nil error if it returns n < len(p).
    // Write must not modify the slice data, even temporarily.
    //
    // Implementations must not retain p.
    Write(p []byte) (n int, err error)
}
```

因为`*ByteCounter`满足`io.Writer`接口，所以可以把它传入`Fprintf`函数中；`Fprintf`函数执行字符串格式化的过程不会去关注`*ByteCounter`具体是如何实现这个接口的，这样体现了**多态同时又具有高内聚和低耦合**的特性。

### 8.4.1 接口的注意事项

1. 在 `Golang` 中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口，才能将该自定义类型的实例(变量)赋给这个接口类型。

2. 一个接口(比如 A 接口)可以继承多个别的接口(比如 B,C 接口)，这时如果要实现 A 接口，也必须将 B,C 接口的方法也全部实现。这里说的其实是**接口组合**：



### 8.4.2 实现接口的条件

我们之前说过，一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。即：

```go
var w io.Writer
w = os.Stdout           // OK: *os.File has Write method
w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
w = time.Now()         // compile error\

var rwc io.ReadWriteCloser
rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
```

```
Cannot use 'time.Now()' (type Time) as the type io.Writer Type does not implement 'io.Writer' as some methods are missing: Write(p []byte) (n int, err error)
```

这里必须要注意的一点是：在T类型的参数上调用一个`*T`的方法是合法的，编译器隐式的获取了它的地址。但这仅仅是一个语法糖：T类型的值不拥有所有`*T`指针的方法，这样对于类型`T`或者`*T`，各自定义的方法是独立的，这样它们实现的接口更少。例如：

```go
type IntSet struct { /* ... */ }
func (*IntSet) String() string
var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
```

就像信封封装和隐藏起信件来一样，接口类型封装和隐藏具体类型和它的值。即使具体类型有其它的方法，也只有接口类型暴露出来的方法会被调用到：

```go'
os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
os.Stdout.Close()                // OK: *os.File has Close method

var w io.Writer
w = os.Stdout
w.Write([]byte("hello")) // OK: io.Writer has Write method
w.Close()                // compile error: io.Writer lacks Close method
```

关于`interface{}`类型，它没有任何方法，可以说任何具体类型包括自定义类型都实现了`interface{}`，但是同样地因为interface{}没有任何方法，我们当然不能直接对它持有的值做任何操作(指调用方法)。

### 8.4.3 接口值

概念上讲一个接口的值，接口值，由两个部分组成，一个具体的类型和那个类型的值。它们被称为接口的动态类型和动态值。在编译阶段，**接口类型是已知的**，并且编译器会检查实现接口的类型是否满足接口的方法集要求。然而，**接口的具体内容（即接口的动态类型和动态值）在编译时是未知的**，因为它依赖于实际运行时赋值给接口的变量类型。

例如：

```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

首先，编译的时候就确定了`w`的类型是`io.Writer`，第一个语句定义了变量w:

```
var w io.Writer
```

初始化时，`Go`会将无论是变量还是接口都初始化为零值，对于一个接口的零值就是它的类型和值的部分都是`nil`。如图所示：

![img](image\ch7-01.png)

**一个接口值基于它的动态类型被描述为空或非空**，所以这是一个空的接口值。你可以通过使用w==nil或者w!=nil来判断接口值是否为空。调用一个空接口值上的任意方法都会产生panic:

```go
w.Write([]byte("hello")) // panic: nil pointer dereference
```

第二个语句将一个`*os.File`类型的值赋给变量w:

```
w = os.Stdout
```

这个赋值过程调用了一个**具体类型(`*os.File`)到接口类型(`io.Writer`)的隐式转换**，这个接口值的动态类型被设为`*os.File`指针的类型描述符，它的动态值持有`os.Stdout`的拷贝；这是一个代表处理标准输出的`os.File`类型变量的指针:

![img](image\ch7-02.png)

调用一个包含`*os.File`类型指针的接口值的Write方法，使得`(*os.File).Write`方法被调用。这个调用输出“hello”:

```go
w.Write([]byte("hello")) // "hello"
```

第三个语句给接口值赋了一个`*bytes.Buffer`类型的值:

```
w = new(bytes.Buffer)
```

现在动态类型是`*bytes.Buffer`并且动态值是一个指向新分配的缓冲区的指针:

![img](image\ch7-03.png)

Write方法的调用也使用了和之前一样的机制：

```go
w.Write([]byte("hello")) // writes "hello" to the bytes.Buffers
```

最后，第四个语句将nil赋给了接口值：

```
w = nil
```

这个重置将它所有的部分都设为nil值，把变量w恢复到和它之前定义时相同的状态:

![img](image\ch7-01.png)

接口值可以使用`==`和`!＝`来进行比较。两个接口值相等仅当它们都是`nil`值，或者它们的动态类型相同并且动态值也根据这个动态类型的`==`操作相等。因为接口值是可比较的，所以它们可以用在`map`的键或者作为`switch`语句的操作数。

如果两个接口值的动态类型相同，但是这个动态类型是不可比较的（比如切片），将它们进行比较就会失败并且panic:

```go
var x interface{} = []int{1, 2, 3}
fmt.Println(x == x) // panic: comparing uncomparable type []int
```

当我们处理错误或者调试的过程中，得知接口值的动态类型是非常有帮助的。所以我们使用fmt包的%T动作:

```go
var w io.Writer
fmt.Printf("%T\n", w) // "<nil>"
w = os.Stdout
fmt.Printf("%T\n", w) // "*os.File"
w = new(bytes.Buffer)
fmt.Printf("%T\n", w) // "*bytes.Buffer"
```

##### 警告：一个包含nil指针的接口不是nil接口

一个不包含任何值的`nil`接口值和一个刚好包含`nil`指针的接口值是不同的。这个细微区别产生了一个容易绊倒每个Go程序员的陷阱。

```go
const debug = true

func main() {
    var buf *bytes.Buffer
    if debug {
        buf = new(bytes.Buffer) // enable collection of output
    }
    f(buf) // NOTE: subtly incorrect!
    if debug {
        // ...use buf...
    }
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
    // ...do something...
    if out != nil {
        out.Write([]byte("done!\n"))
    }
}
```

实际上在`out.Write`方法调用时程序发生了panic，这是因为在`main`函数中将`buf`初始化为动态类型为`*bytes.Buffer`，而动态值为`nil`的指针，在调用函数`f`的时候做了接口类型的隐式转换，对于接口类型，判断一个接口类型变量为`nil`的条件是该变量的动态类型和动态值都是`nil`，所以语句`out != nil`输出为`True`，而对一个动态值为`nil`的变量调用`Write`函数会发生`panic`。

### 8.4.4 类型断言

语法上它看起来像`x.(T)`被称为断言类型，这里x表示一个接口的类型和T表示一个类型。

类型断言分为两种，第一种，如果断言的类型T是一个具体类型，然后类型断言检查x的动态类型是否和T相同:

```go
var w io.Writer
w = os.Stdout
f := w.(*os.File)      // success: f == os.Stdout
c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```

第二种，如果相反地断言的类型T是一个接口类型，**然后类型断言检查是否x的动态类型满足T**:

```go
package main

import "fmt"

// 定义两个接口
type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

// 定义一个结构体
type File struct{}

func (f File) Read() string {
	return "Reading from file"
}

func (f File) Write(s string) {
	fmt.Println("Writing to file:", s)
}

func main() {
	var x Reader = File{} // x 是一个 Reader 接口的值

	// 类型断言，将 x 转为 Writer 接口
	y, ok := x.(Writer)
	if ok {
		fmt.Println("断言成功")
		// y 是一个 Writer 接口的值，但动态部分仍然是 File
		y.Write("Hello")
	}
}
```

`x` 是一个 `Reader` 接口，动态类型是 `File`，动态值是一个 `File` 的实例。类型断言 `x.(Writer)`：

- Go 检查 `x` 的动态类型 `File` 是否实现了 `Writer` 接口，结果是实现了，因此断言成功。
- 断言后，`y` 是一个新的接口值，静态类型是 `Writer`，但动态类型和值仍然是 `File`。

也就是说，当`x.(T)`断言的`T`是接口类型时，类型断言只是“改变”了接口值的 **静态类型**，不会影响到动态类型和动态值部分。

对一个接口值的动态类型我们是不确定的，当我们要对它进行确定类型的断言时，可以指定第二个返回的接收值，这个结果是一个标识成功与否的布尔值：

```go
var w io.Writer = os.Stdout
f, ok := w.(*os.File)      // success:  ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
```

第二个结果通常赋值给一个命名为`ok`的变量。如果这个操作失败了，那么`ok`就是`false`值，第一个结果等于被断言类型的零值，在这个例子中就是一个nil的`*bytes.Buffer`类型。

# 九、 文件操作

操作文件主要借助`os.File`，这个结构体的指针定义了很多对文件操作的方法，如下所示：

![image-20241116165211118](image\image-20241116165211118.png)

## 1. 文件操作的方法

### 9.1.1 打开文件

#### `os.Open`

**功能**：以**只读模式**打开文件。**方法签名**：

```go
func Open(name string) (*File, error)
```

示例：

```go
file, err := os.Open("example.txt")
if err != nil {
    fmt.Println("文件打开失败:", err)
    return
}
defer file.Close() // 确保文件关闭
fmt.Println("文件打开成功:", file.Name())
```

#### `os.OpenFile`

![image-20241118193352871](image\image-20241118193352871.png)

**`os.OpenFile`*可以以指定的模式和权限打开文件。方法的签名如下：

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

参数说明：

`name`: 文件路径。

`flag`: 文件操作模式，常用模式包括：

- `os.O_RDONLY`：只读模式。
- `os.O_WRONLY`：只写模式。
- `os.O_RDWR`：读写模式。
- `os.O_CREATE`：文件不存在时创建文件。
- `os.O_APPEND`：追加写入模式。
- `os.O_TRUNC`：打开文件时清空内容。

`perm`: 文件权限（如 `0666` 表示文件可读写）。

示例：

```go
file, err := os.OpenFile("example.txt", os.O_CREATE|os.O_WRONLY, 0666)
if err != nil {
    fmt.Println("文件打开失败:", err)
    return
}
defer file.Close() // 确保文件关闭
fmt.Println("文件成功打开或创建:", file.Name())
```

#### `bufio.Reader`和它所持有的方法`ReadString`

`bufio.Reader`默认是带缓冲区的，其缓冲区的大小为4096字节，也就是说每次调用方法最多也就读取4096字节，也可以指定这个数，在创建`bufio.Reader`的时候。

```go
package bufio

const (
	defaultBufSize = 4096
)

func NewReaderSize(rd io.Reader, size int) *Reader
// NewReaderSize创建一个具有最少有size尺寸的缓冲、从r读取的*Reader。如果参数r已经是一个具有足够大缓冲的* Reader类型值，会返回r。
```

例如：

```go
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// 示例缓冲区小于数据大小
	data := "Hello, this is a test without the delim."
	reader := bufio.NewReaderSize(bytes.NewReader([]byte(data)), 10) // 缓冲区 10 字节

	// 尝试读取直到分隔符（分隔符不存在）
	result, err := reader.ReadString(',')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Result: %q\n", result)
}
```

结果如下：

```
Result: "Hello,"
```

### 9.1.2 读取文件

#### `ioutil.ReadFile`

函数签名如下：

```go
func ReadFile(filename string) ([]byte, error)
```

例如：

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "code/chapter_9/fileOperation/test.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error: %q\n", err)
	}
	fmt.Printf("%v\n", string(data))
}
```

```
我爱祖国
我爱祖国
我爱祖国  fkafaf
```

注意，`ioutil.ReadFile`有明显的**性能问题**，如下：

- `ReadFile` 会一次性将整个文件读入内存，因此不适合处理大文件。
- 对于大文件，建议使用流式读取（例如 `os.File` 的 `Read` 方法）。

### 9.1.3 创建文件并写入

#### **使用 `os.WriteFile`**

Go 1.16 引入了 `os.WriteFile`，是写文件的简单方法。

```go
import (
	"os"
)

func main() {
	data := []byte("Hello, World!")
	err := os.WriteFile("example.txt", data, 0644) // 文件权限为 0644
	if err != nil {
		panic(err)
	}
}
```

特点：

- 简单高效，一次性写入整个文件。
- 自动创建文件，如果文件存在则覆盖内容。
- **适用场景**：小型文件写入，不需要处理文件内容追加。

------

#### **使用 `os.File` 的 `Write` 和 `WriteString`**

手动控制文件写入，适用于追加写入或需要更高灵活性的场景。

```go
import (
	"os"
)

func main() {
	file, err := os.Create("example.txt") // 创建文件，如果文件存在会清空
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入字节
	_, err = file.Write([]byte("Hello, "))
	if err != nil {
		panic(err)
	}

	// 写入字符串
	_, err = file.WriteString("World!\n")
	if err != nil {
		panic(err)
	}
}
```

特点:

- 灵活，适合多段内容写入或文件内容追加。
- 文件操作需要手动关闭 `file.Close()`。
- **适用场景**：中大型文件，或需要逐步写入内容时。

------

#### **使用 `bufio.Writer`**

`bufio.Writer` 提供了缓冲写入，性能更高，适合频繁小量写入的场景。

```go
import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// 缓冲写入
	_, err = writer.WriteString("Buffered Hello, World!\n")
	if err != nil {
		panic(err)
	}

	// 由于bufio.NewWriter返回的bufio.Writer是带缓冲区的，需要调用Flush()
    // 方法确保数据写入到文件中
	writer.Flush()
}
```

特点

- 提高写入效率，减少磁盘 I/O 操作。
- 需要调用 `Flush()` 将缓冲区内容写入文件。
- **适用场景**：频繁写入小块数据，性能敏感的场景。

------

#### **使用 `ioutil.WriteFile`**（已废弃）

在 Go 1.16 之前，`ioutil.WriteFile` 是常用的写文件方法。

```go
import (
	"io/ioutil"
)

func main() {
	data := []byte("Hello, World!")
	err := ioutil.WriteFile("example.txt", data, 0644) // 文件权限为 0644
	if err != nil {
		panic(err)
	}
}
```

**特点**

- 与 `os.WriteFile` 类似，简单易用。
- **已废弃**，推荐使用 `os.WriteFile`。

------

#### **追加写入（Append）**

对于需要在文件末尾追加内容的场景，可以使用 `os.OpenFile`。



```go
import (
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 追加内容
	_, err = file.WriteString("Appended content!\n")
	if err != nil {
		panic(err)
	}
}
```

**特点**

- 适合日志写入等需要追加内容的场景。
- 支持设置文件打开模式（如只写、只读、追加等）。

------

#### **使用 `fmt.Fprintf` 写入格式化内容**

`fmt.Fprintf` 提供了格式化写入的能力，适合生成结构化文本。

```go
import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入格式化内容
	_, err = fmt.Fprintf(file, "Name: %s, Age: %d\n", "Alice", 30)
	if err != nil {
		panic(err)
	}
}
```

**特点**

- 支持格式化输出，适合生成配置文件或结构化内容。

### 9.1.4 拷贝文件

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CopyFile 接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过 srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)
	//打开 dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//通过 dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}
func main() {
	//将 d:/flower.jpg 文件拷贝到 e:/abc.jpg
	//调用 CopyFile 完成文件拷贝
	srcFile := "d:/flower.jpg"
	dstFile := "e:/abc.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}
}
```

### 9.1.5 统计英文、数字、空格和其他字符数量

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CharCount 定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

func main() {
	//思路: 打开一个文件, 创建一个 Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	fileName := "code/chapter_9/charCount/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	//定义个 CharCount 实例
	var count CharCount
	//创建一个 Reader
	reader := bufio.NewReader(file)
	//开始循环的读取 fileName 的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读到文件末尾就退出
			break
		}
		//为了兼容中文字符, 可以将 strRune 转成 []rune
		strRune := []rune(str)
		//遍历 strRune ，进行统计
		for _, v := range strRune {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	//输出统计的结果看看是否正确
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
```

### 2. flag 解析命令行参数

```go
package main

import (
	"flag"
	"fmt"
)

type flagValue struct {
	user string // -u
	pwd  string // -pwd
	host string // -h
	port int    // -p
}

func main() {
	var f flagValue
	// 注册flag
	flag.StringVar(&f.user, "u", "", "用户名，默认为空")
	flag.StringVar(&f.pwd, "pwd", "", "密码, 默认为空")
	flag.StringVar(&f.host, "h", "localhost", "主机名, 默认为localhost")
	flag.IntVar(&f.port, "port", 3306, "端口号, 默认为3306")

	// 所有变量注册flag完毕之后，调用该方法解析
	flag.Parse()
	// 输出结果
	fmt.Printf("user=%v pwd=%v\nhost=%v port=%v\n",
		f.user, f.pwd, f.host, f.port)
}
```

### 3. `JSON`

`JSON` 应用场景：

![image-20241118203701391](image\image-20241118203701391.png)

#### 9.3.1 `JSON` 的序列化

`json` 序列化是指，将有 **key-valu**e 结构的数据类型(比如**结构体、map、切片**)序列化成 `json` 字符串的操作。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	//演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳"}
	//将 monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster 序列化后=%v\n", string(data))
}

// 将 map 进行序列化
func testMap() {
	//定义一个 map
	var a map[string]interface{}
	//使用 map,需要 make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"
	//将 a 这个 map 进行序列化
	//将 monster 序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

// 演示对切片进行序列化, 我们这个切片 []map[string]interface{}
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用 map 前，需要先 make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)
	var m2 map[string]interface{}
	//使用 map 前，需要先 make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)
	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

// 对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67
	//对 num1 进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func main() {
	//演示将结构体, map , 切片进行序列号
	testStruct()
	testMap()
	testSlice() //演示对切片的序列化
	testFloat64() //演示对基本数据类型的序列化
}
```

#### 9.3.1 `JSON` 的反序列化

`json` 反序列化是指，将 `json` 字符串反序列化成对应的数据类型(比如结构体、map、切片)的操作。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string //.... Sal float64
	Skill    string
}

// 演示将 json 字符串，反序列化成 struct
func unmarshalStruct() {
	//说明 str 在项目开发中，是通过网络传输获取到.. 或者是读取文件获取到
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	//定义一个 Monster 实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v \n", monster, monster.Name)
}

// 演示将 json 字符串，反序列化成 map
func unmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	//定义一个 map
	var a map[string]interface{}
	//反序列化
	//注意：反序列化 map,不需要 make,因为 make 操作被封装到 Unmarshal 函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}

// 演示将 json 字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"
	//定义一个 slice
	var slice []map[string]interface{}
	//反序列化，不需要 make,因为 make 操作被封装到 Unmarshal 函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
```



# 十、 单元测试

在我们工作中，我们会遇到这样的情况，就是去确认一个函数，或者一个模块的结果是否正确，如：

![image-20241118210329296](image\image-20241118210329296.png)

Go 语言中自带有一个轻量级的测试框架 testing 和自带的 go test 命令来实现单元测试和性能测试，testing 框架和其他语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例。通过单元测试，可以解决如下问题:

- 确保**每个函数是可运行，并且运行结果是正确**的；

- 确保写出来的代码**性能是好**的

- 单元测试能及时的发现程序设计或实现的**逻辑错误**，使问题及早暴露，便于问题的定位解决，而**性能测试**的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定。
