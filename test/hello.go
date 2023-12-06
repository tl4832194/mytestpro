// package表示定义包名，包名一般为该文件所在的目录名称
package test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 函数名首字母大写表示全局，可供pakage包外部调用，如果首字母小写，那么包外部无法看到该函数，也就无法调用了
func Hello() string {
	return "hello word!"
}

func SoComan() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "\n"
	}
	return s
}

/*
该函数主要用于扫描重复输入
并将重复输入的数据和次数保存在集合中
如果重复输入次数大于1并打印重复的数据以及次数
*/
func Dup() map[string]int {
	// 定义一个空的map集合比如：map[string:int]
	counts := make(map[string]int)
	// 将标准输入存储起来
	input := bufio.NewScanner(os.Stdin)

	// 循环扫描存储的值（即每次输入的值）
	for input.Scan() {
		fmt.Printf("input: %v\n", input.Text())
		counts[input.Text()]++
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("n and line is: %d\t%s\n", n, line)
				fmt.Printf("counts: %v\n", counts)
			}
		}
	}
	return counts
}

// 因为该服务仅给包内调用，所以首字母小写
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fmt.Printf("counts1111: %v\n", counts)
	}
	// NOTE: ignoring potential errors from input.Err()
}

/*
该函数功能和上面Dup()函数功能类似，也是统计重复行的数量
不过该函数是输入一个文件，一行一行的循环读取文件，将文件中的重复行数量统计出来
*/
func Dup1() map[string]int {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Print("请输入内容：")
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// fmt.Printf("arg: %v\n", arg)
			f, err := os.Open(arg)
			fmt.Printf("f: %v\n", f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup1: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
			for line, n := range counts {
				if n > 1 {
					fmt.Printf("n and line is: %d\t%s\n", n, line)
					fmt.Printf("fileName is: %v\n", arg)
				}
			}
		}
	}
	return counts
}

// 和Dup1函数功能一样，只是采用不通的方法读取文件内容，这里是读取所有内容，然后通过split进行分割切片
func Dup2() map[string]int {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2 err: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	return counts
}
