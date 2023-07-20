package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Message structはレスポンスとして返すメッセージを表します。
type Message struct {
	Text string `json:"message"`
}

// メインのハンドラ関数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// メッセージを作成
	message := Message{Text: "Hello, World!"}

	// レスポンスをJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	// ルータを作成
	r := mux.NewRouter()

	// ハンドラを設定
	r.HandleFunc("/hello", helloHandler).Methods("GET")

	// サーバを開始
	log.Println("APIサーバをポート8080で起動中...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
