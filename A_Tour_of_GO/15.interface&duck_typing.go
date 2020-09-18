package main

import "fmt"

type Human interface { // interface型で指定したメソッドを必ず持っていてくださいと命令
	Say() string // メソッド名を宣言してるだけ(返り値がstringの)　実装するコードを当てはめる入れ物のようなもの
}

type Person struct {
	Name string
}

type Dog struct{
	Name string
}

func (p *Person) Say() string{
	// PersonにSay()というメソッドを定義
	// もしSay()メソッド内で値を弄って返り値を出したいなら*Personにして&Person{"Mike"}で呼び出す
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human){
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	// mike(Human interface型)の中にSay()を持つstructを入れたら
	// mike.Say()
	// mikeはPerson(struct型)のSay()が使える

	var x Human = &Person{"X"}
	// var dog Dog = Dog{"dog"}

	DriveCar(mike) // Run
	DriveCar(x) // Get out
	// DriveCar(dog) dog
	// メソッドのsay()が実装されてないerr
	// 必ずHumanというinterfaceを使ってsay()というメソッドを実装したものをDriveCar関数で使う
	// こういう風に仕分けることをinterfaceのduck typingという
	// https://blog.mmmcorp.co.jp/blog/2018/10/26/go-duck-typing/
}