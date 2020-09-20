package main

import (
	"fmt"
	"os"
)

func foo(){
	// hellofoo worldfoo  ※ deferはこの関数内で最後に実行する命令
	defer fmt.Println("worldfoo")
	fmt.Println("hellofoo")
}

func main() {
	// run success 3 2 1  ※ deferは3 2 1と逆順処理
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	// fileを閉じ理処理などを先に定義できる（忘れない！！）
	file, _ := os.Open("./1.import.go") // 豆知識　アンダースコアで消したとこは通常エラーを返す
	defer file.Close()
	data := make([]byte, 100) // 100byteのバイト配列を作る（スライスだけど文字列をバイトで扱うスライスを慣例的にいう）
	file.Read(data)
	fmt.Println(string(data)) // byte型なので文字列型にキャストする
}
