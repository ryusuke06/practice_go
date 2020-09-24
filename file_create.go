package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct{
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // perm 0600は読み書きのパーミッション（アクセス権限）
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil,err
	}
	return &Page{Title: title, Body: body}, nil
}

func main()  {
	p1 := &Page{Title: "test", Body: []byte("This is a sample Page")} // DBに書き込む
	p1.save()

	p2, _ := loadPage(p1.Title) // 書き込んだファイルを読み取る
	fmt.Println(string(p2.Body))
}