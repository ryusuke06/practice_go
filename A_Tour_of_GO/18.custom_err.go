package main

import (
	"fmt"
)

/*
Goのプログラムは、エラーの状態を error 値で表現します。
error 型は fmt.Stringer に似た組み込みのインタフェースです:

type error interface {
    Error() string
}
*/

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	// なぜポインタにする必要があるか？ → 生成元が異なる同じ文言でもポインタであれば比較できるため
	//（エラーメッセージの比較はシステム上でも使われている → カスタムエラーメッセージで混乱させないため）
	return fmt.Sprintf("User not found: %v", e.Username)
	// User not found: mike
	// stringerを使い自分なりのエラーメッセージをかける

}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}