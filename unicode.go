package main

import "fmt"

func main()  {
	for pos, char := range "日本\x80語" { // \x80 是个非法的UTF-8编码
		fmt.Printf("字符 %#U 始于字节位置 %d\n", char, pos)
	}
}