package main

import (
	"fmt"
	"math/rand"
)
//パッケージ名はインポートパスの最後の要素と同じ名前になる　math/randならrandステートメントで始まるファイル群で構成
//おすすめはしないがmathとだけインポートできる

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

//rand.Intn(n) nまでのランダムな値を返す
//GOでは最初の名前が大文字始まる名前は外部パッケージから参照できる公開（エクスポート）された名前
//stringとintの結合でstring型を返す