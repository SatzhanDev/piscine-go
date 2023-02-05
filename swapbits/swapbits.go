package main

import "fmt"

func SwapBits(octet byte) byte {
	return octet<<4 + octet>>4
}

func main() {
	s := SwapBits(20)
	fmt.Println(s)
}
