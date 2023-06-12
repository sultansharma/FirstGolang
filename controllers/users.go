package controllers

import (
	_ "encoding/json"
	"fmt"
	"main/config"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Users"))
}

type User struct {
	Name string
	Id   int
}

var db = config.GetMySQLDB()


func UserLogin(w http.ResponseWriter, r *http.Request) {

	tokenString, err := config.CreateJWT()

	user := r.FormValue("username")
	var version string
	if err != nil {
		w.Write([]byte("Token creating error"))
	}
	var u User
	err2 := db.QueryRow("SELECT first_name, id FROM users where id=1").Scan(&u.Name,&u.Id)
	if err2 != nil {
		w.Write([]byte(fmt.Sprintf("Error %s", err2)))
	}
	w.Write([]byte(fmt.Sprintf("Hff %s ,%s", tokenString, version)))
	w.Write([]byte(fmt.Sprintf("First Name : %s Username %s", u.Name , user)))
}

func CheckUser(w http.ResponseWriter, r *http.Request) {
	//var count string
	var count int64
	user := r.FormValue("username")
	rows, err := db.Query("SELECT COUNT(*) FROM users WHERE email=?",user);
	if(err!=nil) {
		w.Write([]byte(fmt.Sprintf("Error %s", err)))
	} else {
		for rows.Next() {
			rows.Scan(&count)
		}
	}
	w.Write([]byte(fmt.Sprintf("Hff %d",count)))
}

func Home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(fmt.Sprintf("Helloh")))

}

type UserData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
