package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

// 每个文件中存储的随机数数量
const numbersPerFile = 1000

// 文件数量
const fileCount = 5

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	// 用于通知排序协程的 channel
	done := make(chan string, fileCount)
	var wg sync.WaitGroup

	// 开一个协程负责写文件
	wg.Add(1)
	go func() {
		defer wg.Done()
		writeDataToFile(done)
	}()

	// 主协程负责监听文件生成并排序
	wg.Add(1)
	go func() {
		defer wg.Done()
		sortAndWriteFiles(done)
	}()

	wg.Wait()
	close(done)
	fmt.Println("所有任务完成！")
}

// 生成随机数据并写入文件
func writeDataToFile(done chan<- string) {
	for i := 0; i < fileCount; i++ {
		fileName := fmt.Sprintf("data_%d.txt", i)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("无法创建文件 %s: %v\n", fileName, err)
			continue
		}

		writer := bufio.NewWriter(file)
		for j := 0; j < numbersPerFile; j++ {
			num := rand.Intn(10000) // 生成随机数
			writer.WriteString(strconv.Itoa(num) + "\n")
		}
		writer.Flush()
		file.Close()

		fmt.Printf("文件 %s 写入完成\n", fileName)
		done <- fileName // 通知排序协程
	}
}

// 排序数据并写入新文件
func sortAndWriteFiles(done <-chan string) {
	for fileName := range done {
		// 读取文件中的数据
		numbers, err := readDataFromFile(fileName)
		if err != nil {
			fmt.Printf("读取文件 %s 失败: %v\n", fileName, err)
			continue
		}

		// 排序数据
		sort.Ints(numbers)

		// 写入新文件
		sortedFileName := fmt.Sprintf("sorted_%s", fileName)
		err = writeSortedDataToFile(sortedFileName, numbers)
		if err != nil {
			fmt.Printf("写入排序文件 %s 失败: %v\n", sortedFileName, err)
			continue
		}

		fmt.Printf("文件 %s 排序完成并写入 %s\n", fileName, sortedFileName)
	}
}

// 读取文件中的数据
func readDataFromFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// 将排序后的数据写入新文件
func writeSortedDataToFile(fileName string, numbers []int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		writer.WriteString(strconv.Itoa(num) + "\n")
	}
	writer.Flush()

	return nil
}
