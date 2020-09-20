package main

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string){

	/*
	errはアンダースコアで受け取らない、Openfileでどのような設定で開くか決めれる（リードライト、クリエイト、アペンド）
	0666パーミッション
	*/
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile) // Stdout:：画面上に出るエラー → Multiwriterでlogfileに書き込む
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 書き込むフラグ：日付、時間、短いファイルネーム=短いファイルURL(/import.goみたいな)（ちなみに/UsersからlogをとれるLlongfileもある）
	log.SetOutput(multiLogFile) // まとめたログファイルを出力
}

func main() {
	LoggingSettings("test.log") // 2009/11/10 23:00:00 log.go:25: Exit open af;eifa;oij;eo: no such file or directory
	_, err := os.Open("af;eifa;oij;eo") // エラーを起こすため存在しないファイルを指定
	if err != nil{
		log.Fatalln("Exit", err) // Fatallnはログを返してコードを終了させる　→ エラーハンドリングなどで使う
	}
}