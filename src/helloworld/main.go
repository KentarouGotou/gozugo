package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todoを表す構造体を定義
type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Todoの一覧
var todos = []todo{
	{ID: 1, Title: "タイトルA", Author: "Taro"},
	{ID: 2, Title: "タイトルB", Author: "Jiro"},
	{ID: 3, Title: "タイトルC", Author: "Takuya"},
}

// todosからjsonを作成する関数
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	// ルーターの初期化
	router := gin.Default()
	// /todosにアクセスしたときにgetTodosを呼び出す
	router.GET("/todos", getTodos)

	// サーバーを起動
	router.Run("localhost:8080")
}
