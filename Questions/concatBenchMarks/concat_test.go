package main

import (
	//"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

// 209 sec!!!! don't try
// func Benchmark_plusConcat(b *testing.B) {
// 	holyBytes, err := ioutil.ReadFile("bible.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	words := strings.Split(string(holyBytes), " ")
// 	b.ResetTimer()
// 	for n:= 0; n < b.N; n++ {
// 		plusConcat(words)
// 	}
// 	b.StopTimer()
// }

func Benchmark_bytesBufferConcat(b *testing.B) {
	holyBytes, err := ioutil.ReadFile("bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(holyBytes), " ")
	b.ResetTimer()
	for n:= 0; n < b.N; n++ {
		bytesBufferConcat(words)
	}
	b.StopTimer()
}

func Benchmark_builderConcat(b *testing.B) {
	holyBytes, err := ioutil.ReadFile("bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(holyBytes), " ")
	b.ResetTimer()
	for n:= 0; n < b.N; n++ {
		builderConcat(words)
	}
	b.StopTimer()
}

func Benchmark_copyConcat(b *testing.B) {
	holyBytes, err := ioutil.ReadFile("bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(holyBytes), " ")
	b.ResetTimer()
	for n:= 0; n < b.N; n++ {
		copyConcat(words)
	}
	b.StopTimer()
}

func Benchmark_joinConcat(b *testing.B) {
	holyBytes, err := ioutil.ReadFile("bible.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(holyBytes), " ")
	b.ResetTimer()
	for n:= 0; n < b.N; n++ {
		joinConcat(words)
	}
	b.StopTimer()
}
