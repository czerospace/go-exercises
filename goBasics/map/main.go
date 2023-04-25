package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 初始化一个空的 map, 用于存储单词和出现的次数
	wordCount := make(map[string]int)

	// 打开文件
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to open file:", err)
	}
	defer file.Close()

	// 创建一个 bufio.Scanner 对象
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {

		// 如果这一行是空白行则略过
		line := scanner.Text()
		if len(line) == 0 { // 如果这一行是空白行则略过
			continue
		}
		
		// 将每一行按空格分割成单词
		words := strings.Split(scanner.Text(), " ")

		// 遍历单词，并统计出出现次数
		for _, word := range words {

			// 将单词转换为小写字母，避免大小写不同的单词被视为不同的单词
			word = strings.ToLower(word)

			// 如果单词已经存在于 map 中，则将计数器加 1
			if _, ok := wordCount[word]; ok {
				wordCount[word]++
			} else {
				// 如果单词不存在于 map 中，则将其添加到 map 中，计数器初始化为 1
				wordCount[word] = 1
			}
		}
	}

	// 输出统计结果
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
