package main

import (
	"encoding/json"
	"fmt"
)

type T struct {}

type Person struct {
	Name      string    `json:"name"` // Marshal()のkeyを指定できる
	Age       int		`json:"age,string,omitempty"` // jsonにするときの型も指定できる
	Nicknames []string	`json:"nicknames,omitempty"` // omitemptyは値が空ならMarshal()のときに隠す指定
	T		  *T		`json:"T,omitempty"` // omitemptyはstructだとポインタ型にする必要がある
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!" // structを関数内で作って最後にpに上書きなんてこともできる
	return err
}

func (p Person)MarshalJSON() ([]byte, error) { // カスタマイズしたいなら必ずMarshalJSON()という名前にする！
	a := struct{Name string}{Name: "test"} // 関数内だけで使えるstructを作って
	v, err := json.Marshal(&struct { // jsonで返すやり方もある
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func main()  {
	b := []byte(`{"name":"mike","age":"20","nicknames":["a", "b" ,"c"]}`) // jsonの表現はbyte(リテラル/バッククォート)で表現
	var p Person // bのjsonをstructに入れたいなら、変数宣言
	if err := json.Unmarshal(b, &p); err != nil{ // Unmarshal()でage:"20"をちゃんとintに直してくれる
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames) // mike 20 [a b c]

	v, _ := json.Marshal(p) // jsonに直してネットワークに送りたいとき
	fmt.Println(string(v)) // {"Name":"mike","Age":20,"Nicknames":["a","b","c"]}
}

