package main

import (
	"net/http"

	"strconv"

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
	// JSONを返す
	c.IndentedJSON(http.StatusOK, todos)
}

// todosを新規作成
func postTodo(c *gin.Context) {
	var newTodo todo

	// HTTPリクエストのJSONデータをnewTodoにバインドして、エラーがあればerrにnilを格納
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// 新しいTodoを追加
	todos = append(todos, newTodo)
	// 追加したTodoを返す
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// 指定したIDのTodoを取得
func getTodoById(c *gin.Context) {
	// c.Param("id")でパラメータのkeyを取得
	// strconv.Atoiでそれをidに格納, エラーがあればerrにnilを格納
	id, err := strconv.Atoi(c.Param("id"))
	// エラーがあればreturn
	if err != nil {
		return
	}

	// todosの数だけループ
	for _, t := range todos {
		// もしtodosのIDが指定したIDと一致したら
		if t.ID == id {
			// そのTodoを返す
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	// 一致するIDがなければ404を返す
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// 指定したIDのTodoを更新
func patchTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	// todoを定義
	var todo todo
	// idをtodo.IDに格納
	todo.ID = id
	if err = c.BindJSON(&todo); err != nil {
		return
	}

	// iにindex, tにtodoを格納してループ
	for i, t := range todos {
		// もしtodosのIDが指定したIDと一致したら
		if t.ID == id {
			// そのTodoを更新
			todos[i] = todo
			// 更新したTodoを返す
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// 指定したIDのTodoを削除
func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	for i, t := range todos {
		if t.ID == id {
			// i番目の要素を削除したスライスをtodosに格納
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "todo(" + strconv.Itoa(id) + ") is deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	// ルーターの初期化
	router := gin.Default()
	// /todosにアクセスしたときにgetTodosを呼び出す
	router.GET("/todos", getTodos)
	// /todos/:idにアクセスしたときにgetTodoByIdを呼び出す
	router.GET("/todos/:id", getTodoById)
	// /todosにPOSTしたときにpostTodoを呼び出す
	router.POST("/todos", postTodo)
	// /todos/:idにPATCHしたときにpatchTodoを呼び出す
	router.PATCH("/todos/:id", patchTodo)
	// /todos/:idにDELETEしたときにdeleteTodoを呼び出す
	router.DELETE("/todos/:id", deleteTodo)

	// サーバーを起動
	router.Run("localhost:8080")
}
