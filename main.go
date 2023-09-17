package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var db *Storage

func main() {
	db = InitDb()
	http.HandleFunc("/createUser", createHandler)
	http.HandleFunc("/getInfoUser", getHandler)
	http.ListenAndServe(":8081", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	//err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err1 := db.Create(user)
	if err1 != nil {
		fmt.Println(err)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	// Получаем параметры запроса из URL-строки
	queryParams := r.URL.Query()
	id, err := strconv.Atoi(queryParams.Get("id"))
	if err != nil {
		panic(err)
	}
	user, err := db.GetName(id)
	if err == sql.ErrNoRows {
		// Если нет строк в результате, обработка ситуации
		fmt.Println("Нет результатов в базе данных")
		http.NotFound(w, r)
		//w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Println(err)
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userJSON))
}
