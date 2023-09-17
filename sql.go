package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

func InitDb() *Storage {
	db, err := sql.Open("mysql", "root:serpent@/tasks")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Соединение с БД")
	}
	storage := Storage{
		db: db,
	}
	return &storage
}

func (s *Storage) GetName(id int) (User, error) {
	var user User
	err := s.db.QueryRow("SELECT id,name,last,sex FROM users WHERE id= ?", id).Scan(&user.Id, &user.Name, &user.LastName, &user.Sex)

	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

// Создание новой задачи
func (s *Storage) Create(u User) error {
	res, err := s.db.Prepare("INSERT INTO `users`(id, name, last, sex) VALUES (NULL, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = res.Exec(u.Name, u.LastName, u.Sex)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
