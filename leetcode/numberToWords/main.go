package main

import "fmt"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero" // 处理0的特殊情况
	}

	thousand := 1000
	million := 1000000
	billion := 1000000000

	digitWords := []string{"", "One", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine"}
	tenDigitWords := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	dozenDigWords := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen",
		"Sixteen", "Seventeen", "Eighteen", "Nineteen"}

	// 匿名函数：将三位数以下的数转换为英文单词()
	// 注意只有正确处理空格，确保只有需要的时候再添加
	wordFunc := func(num int) string {
		var resStr string
		if num/100 != 0 {
			resStr += digitWords[num/100] + " Hundred"
		}
		if (num%100)/10 >= 2 {
			if resStr != "" { // 保证不同部分之间有空格
				resStr += " "
			}
			resStr += tenDigitWords[(num%100)/10]
			if num%10 != 0 {
				resStr += " " + digitWords[num%10]
			}
		} else if (num%100)/10 == 1 {
			if resStr != "" { // 保证不同部分之间有空格
				resStr += " "
			}
			resStr += dozenDigWords[num%10]
		} else if num%10 != 0 {
			if resStr != "" { // 保证不同部分之间有空格
				resStr += " "
			}
			resStr += digitWords[num%10]
		}
		return resStr
	}

	resStr := ""

	// 处理 Billion
	if num/billion != 0 {
		resStr += wordFunc(num/billion) + " Billion"
	}

	// 处理 Million
	if (num%billion)/million != 0 {
		if resStr != "" {
			resStr += " "
		}
		resStr += wordFunc((num%billion)/million) + " Million"
	}

	// 处理 Thousand
	if (num%million)/thousand != 0 {
		if resStr != "" {
			resStr += " "
		}
		resStr += wordFunc((num%million)/thousand) + " Thousand"
	}

	// 处理小于一千的部分
	if num%thousand != 0 {
		if resStr != "" {
			resStr += " "
		}
		resStr += wordFunc(num % thousand)
	}

	return resStr
}

func main() {
	fmt.Println(numberToWords(12345))

}
