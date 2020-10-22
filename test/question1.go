package main

const s = "Go101.org"
// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128
var c byte= 1 << len(s[:])
var d byte= c / 128

func main() {
	println(len(s))
	println(len(s[:]))
	println(1 << len(s))
	println(1 << len(s[:]))
	println(a, b,c,d)
}
