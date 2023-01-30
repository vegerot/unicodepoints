package main

import (
	"fmt"
	"os"
	"strconv"
	"unicodepoints/unicodepoints"
)

func main() {

	codepoint, err := strconv.ParseInt(os.Args[2], 16, 32)
	if err != nil {
		panic(err)
	}
	var output interface{}
	operation := os.Args[1]
	switch operation {
	case "--bytes":
		output = unicodepoints.CodePointToBytes(uint32(codepoint))
	case "--string":
		output = unicodepoints.CodePointToString(uint32(codepoint))
	}
	if codepoint < 0 {
		panic("Codepoint must be positive")
	}
	fmt.Println(output)

}
