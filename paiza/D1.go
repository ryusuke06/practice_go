/*
N 週間後は何日後かを計算してください。

末尾に改行を入れ、余計な文字、空行を含んではいけません。
*/

package main
import "fmt"
func main(){
    var N int
    fmt.Scan(&N)
    n := N*7
    fmt.Println(n)
}