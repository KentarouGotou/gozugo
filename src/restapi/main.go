package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/users", handleGetUserList)
	http.HandleFunc("/users/", handleRequest)

	server.ListenAndServe()
}

// リクエストを処理するハンドラ
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error

	// レスポンスヘッダーを設定
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// リクエストメソッドによって処理を分岐
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// userの一覧を取得するハンドラ
func handleGetUserList(w http.ResponseWriter, r *http.Request) {
	var err error

	// GetUsers()でユーザーの一覧を取得
	users, err := GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// json.MarshalIndent()でusersをjsonに変換
	output, err := json.MarshalIndent(&users, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// レスポンスヘッダーを設定
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Write(output)
}

// GETで指定したidのuserを取得するハンドラ
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// Base は，path の最後の要素を返します。 末尾のスラッシュは，最後の要素を抽出する前に削除されます。
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	// GetUser()でidに該当するuserを取得
	user, err := GetUser(id)
	if err != nil {
		return
	}

	// json.MarshalIndent()でuserをjsonに変換
	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	// レスポンスヘッダーを設定
	w.WriteHeader(200)
	w.Write(output)
	return
}

// POSTでuserを新規作成するハンドラ
func handleUser(w http.ResponseWriter, r *http.Request) (err error) {

	contentLength := r.ContentLength
	contentBody := make([]byte, contentLength)
	r.Body.Read(contentBody)

	var user User
	err = json.Unmarshal(contentBody, &user)
	if err != nil {
		return
	}

	err = user.CreateUser()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

// PUTで指定したidのpostを更新するハンドラ
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	user, err := GetUser(id)
	if err != nil {
		return
	}

	contentLength := r.ContentLength
	contentBody := make([]byte, contentLength)
	r.Body.Read(contentBody)

	err = json.Unmarshal(contentBody, &user)
	if err != nil {
		return
	}

	err = user.UpdateUser()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

// DELETEで指定したidのpostを削除するハンドラ
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := retrieve(id)
	if err != nil {
		return
	}

	err = post.delete()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}
