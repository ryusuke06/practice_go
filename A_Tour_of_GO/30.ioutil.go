package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main()  {
	//context, err := ioutil.ReadFile("main.go")
	// os.Openでも構わないが入出力ならioutilが良い　packetを読み込んでpacketを書き出すとか
	//if err != nil{
	//	log.Fatalln(err)
	//}
	//fmt.Println(string(context)) // contextはbyte型で返って来るのでキャスト

	//if err := ioutil.WriteFile("ioutil_temp.go", context, 0666); err != nil{
	//	log.Fatalln(err)
	//} contextの内容を出力したgoファイルを書き出す

	r := bytes.NewBuffer([]byte("abc"))
	// Newbufferで作ったbuffer(ネットからきたデータを一時的に保存する領域)にbyte配列で保存
	context, _ := ioutil.ReadFile(r)
	fmt.Println(string(context)) // abc
}

