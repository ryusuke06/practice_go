package main

import "fmt"

func thirdPartyConnectDB(){
	panic("Unable to connect database!") // パニックを起こし、引数の文字列を返す(パニックを起こすべきではないが)
}

func save() {
	defer func(){
		s := recover() // panicでシステムを制終了させないように
		fmt.Println(s) // Unable to connect database!
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?") // OK?
}