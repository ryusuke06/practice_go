/*
あなたは英語の文法チェック・修正システムの作成を担当しています。 このシステムを実現するには、英単語を複数形に変換する機能が必要です。

単語の複数形への変換は、次のルールに従い行われます。

・末尾が s, sh, ch, o, x のいずれかである英単語の末尾に es を付ける
・末尾が f, fe のいずれかである英単語の末尾の f, fe を除き、末尾に ves を付ける
・末尾の1文字が y で、末尾から2文字目が a, i, u, e, o のいずれでもない英単語の末尾の y を除き、末尾に ies を付ける
・上のいずれの条件にも当てはまらない英単語の末尾には s を付ける

入力された英単語を複数形に変換するプログラムを作成してください。

※必ずしも実在の英単語が入力されるわけではありません。単純に文字列を上記のルールに沿って変換するプログラムを作成してください。

条件
すべてのテストケースで以下の条件を満たします。

・1 ≦ N ≦ 10
・a_i は a～z のアルファベット小文字の並べた文字列
・2 ≦ (a_i の長さ) ≦ 20
・架空の英単語が与えられることがあります

N は変換しなければならない英単語の数を表します。
a_i は複数形に変換しなければならない英単語を表します。
*/

package main
import (
    "bufio"
    "fmt"
    "os"
    )

var sc = bufio.NewScanner(os.Stdin) // 改行含む標準出力 https://qiita.com/tnoda_/items/b503a72eac82862d30c6

func nextLine() string{
    sc.Scan()
    return sc.Text()
}

func trans(s string)string {
    if s[len(s)-1:len(s)]=="s"||s[len(s)-2:len(s)]=="ch"||s[len(s)-2:len(s)]=="sh"||s[len(s)-1:len(s)]=="o"||s[len(s)-1:len(s)]=="x"{
        s += "es"
    }else if s[len(s)-1:len(s)]=="f"{
        s = s[0:len(s)-1] + "ves"
    }else if s[len(s)-2:len(s)]=="fe"{
        s = s[0:len(s)-2] + "ves"
    }else if s[len(s)-1:len(s)]=="y"{
        if s[len(s)-2:len(s)-1]=="a"||s[len(s)-2:len(s)-1]=="i"||s[len(s)-2:len(s)-1]=="u"||s[len(s)-2:len(s)-1]=="e"||s[len(s)-2:len(s)-1]=="o"{
        	// 正規表現を使ってもっとスマートにできそう
            s += "s"
        }else{
            s = s[0:len(s)-1] + "ies"
        }
    }else{
        s += "s"
    }
    return s
}

func main(){
	var n int
	fmt.Scan(&n)

	for i:=1;i<=n;i++{
		fmt.Println(trans(nextLine()))
	}
}