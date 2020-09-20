package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./err_handling")
	if err != nil{
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	// errがshort declaration(短い宣言 := のこと)でinitialize(初期値をセットすること)されているのはconstが初めてinitializeされるから(最低一つが未定義なら使える、定義済のerrは上書き)
	count, err := file.Read(data)
	if err != nil{
		log.Fatalln("Error")
	}
	fmt.Println(const, string(data))

	if err = os.Chdir("test"); err != nil{ // countの時とは異なり、initializeされているerrのみなのでoverride(上書き)の"="になってる
		log.Fatalln("Error")
	}
}