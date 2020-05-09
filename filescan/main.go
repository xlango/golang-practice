package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	// 确定初始目录
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 当检测到输入时取消遍历
	go func() {
		// 读一个字节
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// 并行遍历文件树的每个根
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 定期打印结果
	// tick := time.Tick(500 * time.Millisecond)
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// 耗尽 fileSizes 以允许已有的  goroutine 结束
			for range fileSizes {
				// 不执行任何操作
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes 关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	// 最终统计
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir 递归地遍历以dir为根目录的整个文件树
// 并在 fileSizes 上发送每个已找到的文件的大小
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema 是一个用于限制目录并发数的计数信号量
var sema = make(chan struct{}, 20)

// dirents 返回 dir 目录中的条目
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // 获取令牌
	case <-done:
		return nil // 取消
	}
	defer func() { <-sema }() // 释放令牌

	// 读取目录
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du：%v\n", err)
		return nil
	}
	defer f.Close()

	// 不做限制，读取所有条目
	entries, err := f.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du：%v\n", err)
	}
	return entries
}

// 创建取消通道
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
