package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// User struct
type User struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Note struct
type Note struct {
	NoteId   int       `json:"note_id"`
	Position [6]string `json:"position"`
	Ties     [6]string `json:"ties"`
	Effects  [6]string `json:"effects"`
	Duration int       `json:"duration"`
}

// Tab struct
type Tab struct {
	TabId  int    `json:"tab_id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Notes  []Note `json:"notes"`
	Tempo  int    `json:"tempo"`
}

// var tabs = []Tab{
// 	{
// 		Id: 1,
// 		Artist: "Beatles",
// 		Title: "Let it be",
// 		Notes: []Note{
// 			{
// 				Position: [6]string{"-", "1", "0", "2", "3", "-"},
// 				Ties: [6]string{"-", "-", "-", "-", "-", "-"},
// 				Effects: [6]string{"-", "vib", "-", "-", "-", "-"},
// 				Duration: 960
// 			},
// 			{
// 				Position: [6]string{"-", "1", "-", "-", "-", "-"},
// 				Ties: [6]string{"-", "t", "-", "-", "-", "-"},
// 				Effects: [6]string{"-", "vib", "-", "-", "-", "-"},
// 				Duration: 960
// 			}
// 			},
// 		Tempo: 120
// 	},
// 	{
// 		Id: 1,
// 		Artist: "Beatles",
// 		Title: "Let it be",
// 		 Notes: []Note{
// 			{
// 				Position: [6]string{"-", "-", "-", "7", "-", "-"},
// 				Ties: [6]string{"-", "-", "-", "-", "-", "-"},
// 				Effects: [6]string{"-", "-", "-", "bn2", "-", "-"},
// 				Duration: 480
// 			},
// 			{
// 				Position: [6]string{"-", "-", "-", "9", "10", "-"},
// 				Ties: [6]string{"-", "-", "-", "t", "-", "-"},
// 				Effects: [6]string{"-", "-", "-", "-", "-", "-"},
// 				Duration: 480
// 			}
// 		},
// 		Tempo: 120
// 	},
// }

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./example.sqlite")
	if err != nil {
		panic(err)
	}
}

// Users list
func GetUsers() (users []User, err error) {
	// Query the database for all users
	rows, err := Db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.UserId, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return
}

// User by id
func GetUser(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from users where user_id = ?", id).Scan(&user.UserId, &user.Username, &user.Password)
	return
}

// User by username
func GetUserByUsername(username string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select * from users where username = ?", username).Scan(&user.UserId, &user.Username, &user.Password)
	return
}

// Create user
func (user *User) CreateUser() (err error) {
	statement := "insert into users (username, password) values (?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password)
	return
}
