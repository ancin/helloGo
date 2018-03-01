package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/***
*分块读取
*
****/

func processBlock(line []byte) {
	os.Stdout.Write(line)
}

func readBlock(filePath string, bufSize int, hookfn func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := make([]byte, bufSize) //一次读取多少个字节
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		hookfn(buf[:n]) // n 是成功读取字节数
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

	return nil
}

func readFile() {
	fmt.Println("start read file")
	readBlock("output.txt", 10000, processBlock)
}
