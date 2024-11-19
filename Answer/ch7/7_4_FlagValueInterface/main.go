package main

import (
	"flag"
	"fmt"
)

// 练习 7.6： 对tempFlag加入支持开尔文温度。

type Celsius float64    // 摄氏度
type Fahrenheit float64 // 华式温度
type Kelvin float64     // 开尔文温度

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

/*
package flag
// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
*/
// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	// flag.CommandLine.Var 用于向命令行解析器注册一个自定义类型的命令行参数
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

type kelvinFlag struct{ Kelvin }

// 练习 7.7： 解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。
// 这里为啥直接输出fmt.Println(*temp)是20°C呢，是因为Celsius类型实现了String方法，
// 在执行fmt.Println(*temp)时，*temp的类型`Celsius`，是会默认调用*temp.String()方法，于是输出会带°C

func main() {
	var temp = CelsiusFlag("temp", 20.0, "the temperature(Celsius)")
	flag.Parse()
	fmt.Println(*temp)
}
