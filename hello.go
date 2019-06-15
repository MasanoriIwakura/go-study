package main

import "fmt"

// グローバルスコープ
var g1 int
var g2 int = 1 // 書けるけど、冗長な宣言は警告出るらしい
var g3 = 2

func main() {
	fmt.Println("Hello World.")

	// ローカルスコープ
	var l1 int
	l1 = 3
	l2 := 4

	var str = "文字列"
	var b1 bool
	b2 := true

	var (
		f1 float32
		f2 = 0.1
	)

	fmt.Printf("g1=%d, g2=%d, g3=%d\n", g1, g2, g3)
	fmt.Printf("l1=%d, l2=%d\n", l1, l2)
	fmt.Printf("str=%s\n", str)
	fmt.Printf("b1=%t, b2=%t\n", b1, b2)
	fmt.Printf("f1=%f, f2=%f\n", f1, f2)
}
