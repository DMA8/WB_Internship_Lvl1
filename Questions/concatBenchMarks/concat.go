package main

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"
)

var params = []string{"string1", "string2", "string3", "string4"}
// BENCHMARKS!
func main() {
	println(plusConcat(params))
	println(bytesBufferConcat(params))
	println(builderConcat(params))
	println(copyConcat(params))
	println(joinConcat(params))
	fmt.Println(unsafe.Sizeof(struct{}{}))
	a := make(map[int]int)
	a[12312] = 2321312
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[-1] = -1
	for key, value := range a{
		fmt.Println(key, " --- ", value)
	}
	
}

func plusConcat(args []string) string{
	result := ""
	for _, arg := range args {
		result += arg
	}
	return result
}

func bytesBufferConcat(args []string) string {
	buffer := bytes.Buffer{}
	for _, arg := range args {
		buffer.WriteString(arg)
	}
	return buffer.String()
}

func builderConcat(args []string) string {
	builder := strings.Builder{}
	for _, arg := range args {
		builder.WriteString(arg)
	}
	return builder.String()
}

func copyConcat(args []string) string {
	allLength := 0
	func(){
		for _, arg := range args{
			allLength += len(arg)
		}
	}()
	offset := 0
	buffer := make([]byte, allLength)
	for _, arg := range args {
		offset += copy(buffer[offset: ], arg)
	}
	return string(buffer)
}

func joinConcat(args []string) string {
	return strings.Join(args, "")
}
