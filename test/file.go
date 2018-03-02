/**
User:       wliangde
CreateTime: 18/3/1 下午12:22
**/
package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//++++++++++++++++++遍历目录，统计文件大小
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err %v\n", err)
		return nil
	}
	return entries
}

func WalkDir(dir string, fileSizes chan<- int64) {
	for _, file := range dirents(dir) {
		if file.IsDir() {
			newDir := filepath.Join(dir, file.Name()) //遍历子目录下的文件
			WalkDir(newDir, fileSizes)
		} else {
			fileSizes <- file.Size()
			//fmt.Println("文件", dir, file.Name(), "大小", file.Size())
		}
	}
}

func GoWalk(dir string) {
	startTime := time.Now()

	fileSizes := make(chan int64, 1000)
	go func() {
		WalkDir(dir, fileSizes)
		close(fileSizes)
	}()

	fileCnt := 0
	var totalFileSize int64 = 0
	for s := range fileSizes {
		fileCnt++
		totalFileSize += s
	}

	//++++++++++++++++++

	fmt.Println("遍历目录", dir, "遍历文件个数", fileCnt, "总大小", totalFileSize, "总耗时", time.Now().Sub(startTime))
}
