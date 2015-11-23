package main

import "fmt"

func main(){
	fmt.Printf("",shouldEscape(' '))
	fmt.Printf("",shouldEscape('?'))
	fmt.Printf("",shouldEscape('='))
	fmt.Printf("",shouldEscape('&'))
	fmt.Printf("",shouldEscape('1'))
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}
