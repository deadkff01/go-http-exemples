package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Status  string `json:"status"`
}


// user handler
func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userById(w, r, id)
	case r.Method == "GET":
		users(w, r)
	case r.Method == "POST":
		createUser(w, r)
	case r.Method == "DELETE":
		deleteUser(w, r)
	case r.Method == "PATCH":
		updateUser(w, r)	
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func responseJson(w http.ResponseWriter, json[] byte) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func userById(w http.ResponseWriter, r *http.Request, id int) {
	Connect()
	defer Close()

	var u User
	sqlStatement := `select id, name from users where id = $1;`
	DB.QueryRow(sqlStatement, id).Scan(&u.ID, &u.Name)
	fmt.Println(u)
	json, _ := json.Marshal(u)
	responseJson(w, json)
}

func users(w http.ResponseWriter, r *http.Request) {
	Connect()
	defer Close()

	rows, _ := DB.Query("select id, name from users")
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}
	json, _ := json.Marshal(users)
	responseJson(w, json)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	Connect()
	defer Close()

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	insertStatement := `INSERT INTO users (name) VALUES ($1)`
	_, err = DB.Exec(insertStatement, u.Name)

	if err != nil {
	 panic(err)
	}

	response := Response{"ok"}
	json, _ := json.Marshal(response)
	responseJson(w, json)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	Connect()
	defer Close()

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	sqlStatement  := `DELETE FROM users WHERE id = $1`
	_, err = DB.Exec(sqlStatement , u.ID)

	if err != nil {
	 panic(err)
	}
	
	response := Response{"ok"}
	json, _ := json.Marshal(response)
	responseJson(w, json)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	Connect()
	defer Close()

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	sqlStatement := `
	UPDATE users SET name = $2 WHERE id = $1
	RETURNING id, name`
	err = DB.QueryRow(sqlStatement , u.ID, u.Name).Scan(&u.ID, &u.Name)

	if err != nil {
	 panic(err)
	}
	
	json, _ := json.Marshal(u)
	responseJson(w, json)
}