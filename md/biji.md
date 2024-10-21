# GO语言简述

### 垃圾回收机制

内存自动回收，不用开发员自己管理，程序员只用关心代码逻辑，不用担心内存泄漏。

### 原生支持并发（重要）

- 从语言方面支持并发，实现简单；
- goroutine，轻量级线程，**可实现大量并发**，高效率利用核；
- 基于CPS并发模型(Communicating Sequential Processes)实现；

### 管道通信机制

通过管道Channel，可以实现不同的goroutine之间的相互通信

### 函数可以返回多个值

```go
func sumAndSub(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
```

### 创新：切片、延时执行函数defer语句等

### GO 语言开发注意事项

- GO应用程序执行的入口是main()函数

### 

# GO变量

## 概念

什么是变量：**变量**是用来存储值的命名实体。变量在编程中扮演着重要角色，允许你存储和操作数据。Go 中的变量可以存储不同的数据类型，比如整数、浮点数、字符串等。

## 使用变量

### 1. 使用 `var` 关键字显式声明

这种方式适用于需要显式声明变量类型的情况，并且可以选择是否同时赋值。

```go
// 声明变量并赋初值
var name string = "Alice"
var age int = 30

// 只声明变量，不赋值，使用默认的零值
var height float64
```

在这种方式下，如果不赋初值，Go 会给变量赋一个默认的**零值**。如上例中的 `height` 默认值为 `0.0`。

### 2. 使用 `var` 关键字并省略类型（类型推断）

当为变量赋值时，Go 可以根据赋值的内容自动推断出变量的类型，因此你可以省略类型声明。

```go
var name = "Bob"   // Go 自动推断为 string 类型
var age = 25       // Go 自动推断为 int 类型
```

### 3. 简短声明（仅限函数内部）

在函数内部，可以使用简短声明符 `:=` 进行变量声明并赋值。Go 会根据右侧的值自动推断变量类型。

```go
name := "Charlie"
age := 22
height := 1.80
```

> 注意：简短声明只能在函数内部使用，不能用于全局变量声明。

### 4. 多变量声明

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

### 如何在程序中查看一个变量的数据类型和所占的字节数：

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

## 浮点数 float

在 Go 语言中，小数类型主要通过 **浮点数类型** 来表示。Go 提供了两种主要的浮点数类型，用于处理不同精度的浮点运算：

### 1. **`float32`**

- `float32` 是 32 位的浮点数，表示精度较低的小数，符合 IEEE-754 标准。
- 它能精确表示小数点后大约 7 位的有效数字。

### 2. **`float64`**

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

Golang的浮点数申明默认为float64

```go
var z = 1.01
fmt.Printf("z 的类型为 %T\\n", z)

// z 的类型为 float64
```

### 

# 字符串

Go语言中没有专门的字符类型，如果要存储一个字符的话，通常使用`rune`类型。

在 Go 语言中，字符串是由字节（`byte`）组成的。Go 的字符串实际上是一个字节序列，且是不可变的。这意味着每个字符串都是一组按顺序排列的字节，底层存储为 UTF-8 编码的字节序列。

```go
s := "你好"
fmt.Println(len(s)) // 输出 6，因为 "你" 和 "好" 各自占用 3 个字节
```

### 字符类型使用细节

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

### 字符串拼接方式

### 1. 使用 `+` 操作符拼接

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

### 2. 使用 `fmt.Sprintf` 格式化拼接

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

### 3. 使用 `strings.Join` 函数

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

### 4. 使用 `bytes.Buffer`

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

### 5. 使用 `strings.Builder` （Go 1.10 及之后）

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

### 基本数据类型的转换

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

### 字符串与基本类型之间的转换

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

# 指针类型

在 Go 语言中，指针是一个非常重要的概念，它指向一个变量的内存地址。通过指针，你可以直接访问和修改存储在内存中的值，而不是值的副本。在 Go 中，指针类型的使用方式与其他语言（如 C/C++）类似，但也有一些 Go 特有的特性。

### Go 中的指针基础

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

### Go 中指针的特性

1. **Go 中不支持指针运算**： 与 C/C++ 不同，Go 不允许对指针进行算术运算。你不能像在 C 中那样通过增加或减少指针的值来访问相邻的内存位置。这是因为 Go 语言希望减少内存管理中的错误，避免低级的指针操作导致的错误行为。

2. **零值指针**： 指针的零值是 `nil`。如果一个指针没有被初始化，它默认指向 `nil`。尝试解引用一个 `nil` 指针会导致运行时错误（`runtime panic`）。

   ```go
   var p *int
   fmt.Println(p)  // 输出: <nil
   ```

3. **指针与函数**： Go 中的函数传参是值传递，意味着当你将一个变量传递给函数时，函数得到的是这个变量的副本。如果希望函数修改传入的变量，你需要传递变量的指针。

### 例子：

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
   go
   复制代码
   var p *int = new(int)
   fmt.Println(*p)  // 输出: 0，p 指向的是一个零值初始化的 int
   *p = 42
   fmt.Println(*p)  // 输出: 42，修改了该指针指向的值
   ```

2. **make() 与指针**： 虽然 `new()` 分配内存用于基本类型或结构体，但 `make()` 是专门用于创建并初始化引用类型（如切片、映射和通道）的函数。这些类型本质上也是指针，但它们有更复杂的底层结构，因此 `make()` 会负责内存分配和初始化。

### 指针与结构体

在 Go 中，使用指针可以更高效地处理结构体类型。通过传递结构体指针，你可以避免拷贝整个结构体的数据，只需传递其地址，并通过指针修改结构体的字段。

### 例子：

```go
go
复制代码
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

# 函数

## 函数-调用过程

介绍：为了让大家更好的理解函数调用过程, 看两个案例，并画出示意图，这个很重要

1. 传入一个数+1

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/c2892132-f556-4b19-9f1d-9a2554cd25ec/a31a342a-4489-48ed-9d0c-734aa1527e47/image.png)

对上图说明 (1) 在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间 和其它的栈的空间区分开来

(2) 在每个函数对应的栈中，数据空间是独立的，不会混淆 (3) 当一个函数调用完毕(执行完毕)后，程序会销毁这个函数对应的栈空间。 2. 计算两个数,并返回：

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/c2892132-f556-4b19-9f1d-9a2554cd25ec/c0ec86da-46da-4cad-9969-922644381d6b/image.png)

### 函数递归

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/c2892132-f556-4b19-9f1d-9a2554cd25ec/ed5562aa-86bc-40dc-862f-f08dc521e4f0/image.png)

对上面代码分析的示意图：

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/c2892132-f556-4b19-9f1d-9a2554cd25ec/eade0c21-11b6-4ae7-9ba1-a2919538240d/image.png)

**函数递归需要遵守的重要原则**:

1. 执行一个函数时，就创建一个新的受保护的独立空间(新函数栈)
2. 函数的局部变量是独立的，不会相互影响
3. 递归必须向退出递归的条件逼近，否则就是无限递归，死龟了:)
4. 当一个函数执行完毕，或者遇到 return，就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或者返回时，该函数本身也会被系统销毁

### 函数使用的注意事项

1. **基础数据类型**和**数组**在传入函数时，默认情况下是**值传递**。

这意味着函数会接收到原始数据的一个副本，而不是对原始数据的引用。因此，在函数内部对这些数据的任何修改都不会影响到外部的原始数据。

````
```go
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

```
````

1. 如果希望函数内的变量能修改函数外的变量(指的是默认以值传递的方式的数据类型)，可以传 入**变量的地址&**，函数内以指针的方式操作变量。从效果上看类似引用 。
2. **Go 语言不支持函数重载。函数重载**即多个函数可以有相同的名称但不同的参数列表或返回类型。
3. 在 Go 中，**函数也是一种数据类型**，可以赋值给一个变量，则该变量就是一个函数类型的变量

了。通过该**变量可以对函数调**用：

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

## 包 package

在Golang（Go语言）中，**包（package）** 是代码组织和模块化的基本单元，包的设计简化了代码复用和依赖管理，同时提供了封装和命名空间的概念。

### 1. 包的组织结构

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

### 2. **包的导入机制**

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

### 3. **包的初始化**

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

### 4. **包的可见性**

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

### 5. **包的导入别名与匿名导入**

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

### 6. **包的循环依赖**

Go **不允许包之间存在循环依赖**。如果包A导入了包B，同时包B也试图导入包A，这将导致编译失败。Go强制包之间的依赖是单向的，这样可以简化依赖关系，并确保初始化顺序不会混乱。

如果遇到循环依赖，通常可以通过将公共代码提取到一个单独的包中来解决。