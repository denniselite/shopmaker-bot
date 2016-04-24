package main

import "time"

type TextNotice struct {
	Id      int	`db:"id, primarykey, autoincrement"`
	UserId  int	`db:"user_id"`
	Message string	`db:"message"`
}

type User struct {
	Id   		int	`db:"id, primarykey"`
	Lang 		string	`db:"lang"`
	CreatedAt	time.Time `db:"created_at"`
}