package main

import (
	"fmt"
	"regexp"
	"sort"
	"time"
)

func timepkg()  {
	t := time.Now()
	fmt.Println(t) // 2020-09-21 10:15:37.726678 +0900 JST m=+0.000129316
	fmt.Println(t.Format(time.RFC3339)) // 2020-09-21T10:15:37+09:00
}

func regexpkg()  {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple") // true 正規表現と比較
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e") // 何度も正規表現を使う場合、最適化しておく
	ms := r.MatchString("apple")
	fmt.Println(ms) // true やってることは同じでもrはこの後も使いまわせる

	// s := "/view/test" のような形をマッチさせるには？
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test") // r2に当てはまるか→当てはまる
	fmt.Println(fs) // /view/test
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss) // [/view/test view test]
	// 正規表現の最も左の文字列にマッチした文字列を返す（今回は/view）スライスでバラした状態でも持ってこれる
	fmt.Println(fss, fss[0], fss[1], fss[2]) // /view/test view test スライスから出せる
}

func sortpkg() {
	i := []int{5,3,4,2,6}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name }) // アルファベット順
	// sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age }) 年齢順
	// forループを二重に入れてi,jで回しているイメージ レス関数で小さいものを前に
	fmt.Println(i, s, p)
	// [2 3 4 5 6] [a d f] [{Bob 50} {Mike 30} {Nancy 20} {Vera 40}]
}

func main()  {
	timepkg()
	regexpkg()
	sortpkg()
}