package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main()  {
	//resp, _ := http.Get("http://example.com")
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))  url先のHTMLを表示
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2") // example.comの後につけるurlを生成 GETするときは?をつけるルール
	fmt.Println(base) // http://example.com
	endpoint := base.ResolveReference(reference).String() // アクセスするエンドポイントを生成
	fmt.Println(endpoint) // http://example.com/test?a=1&b=2
	req, _ := http.NewRequest("GET", endpoint, nil)
	// HTTPリクエストがPOSTならbodyにnilではなくbyte型の値を入れる
	// この時点ではstructを作っただけでまだアクセスはしていない
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	// ヘッダーも送信したいなら 例えばヘッダー情報でwebサーバにキャッシュを保存したりしなかったりが要求できる
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q) // map[a:1] b:[2] c:[3&%]]
	fmt.Println(q.Encode())
	// a=1&b=2&c=3%26%25 ampersandは区切り文字として使われているのでエンコーディングすれば3&%は3%26%25に変えてくれる
	req.URL.RawQuery = q.Encode() // 3&%のような物はエンコードする必要があるから

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

