package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	users_table = `CREATE TABLE IF NOT EXISTS users (
		uuid TEXT PRIMARY KEY NOT NULL,
		name CHAR(50) NOT NULL,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(50) NOT NULL UNIQUE, 
		password VARCHAR(50) NOT NULL,
		token TEXT,
		expiretime
	);`
	post_table = `CREATE TABLE IF NOT EXISTS post (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		uuid TEXT NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author VARCHAR(50) NOT NULL, 
		createdat VARCHAR(50) NOT NULL,
		categories BIGINT UNSIGNED,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		FOREIGN KEY (uuid) REFERENCES users(uuid) ON DELETE CASCADE,
		FOREIGN KEY (author) REFERENCES users(username) ON DELETE CASCADE
	);`
	comments_table = `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postID INTEGER,
		content TEXT NOT NULL,
		author VARCHAR(50) NOT NULL, 
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		createdat VARCHAR(50) NOT NULL,
		FOREIGN KEY (postID) REFERENCES post(id) ON DELETE CASCADE,
		FOREIGN KEY (author) REFERENCES users(username) ON DELETE CASCADE
	);`

	likePostTable = `CREATE TABLE IF NOT EXISTS likePost (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userID TEXT,
		postID INTEGER DEFAULT 0,
		status INTEGER DEFAULT 0,
		FOREIGN KEY (userID) REFERENCES users(uuid) ON DELETE CASCADE,
		FOREIGN KEY (postID) REFERENCES post(id) ON DELETE CASCADE
		);`

	likeCommentsTable = `CREATE TABLE IF NOT EXISTS likeComments(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userID TEXT,
		commentsID INTEGER DEFAULT 0,
		status INTEGER DEFAULT 0,
		FOREIGN KEY (userID) REFERENCES users(uuid) ON DELETE CASCADE,
		FOREIGN KEY (commentsID) REFERENCES comments(id) ON DELETE CASCADE
		);`
)

// Создание таблицы пользователя
func CreatTables(db *sql.DB) error {
	allTables := []string{users_table, post_table, comments_table, likePostTable, likeCommentsTable}
	for _, v := range allTables {
		stmt, err := db.Prepare(v)
		if err != nil {
			return fmt.Errorf("Create table: %w", err)
		}
		_, err = stmt.Exec()
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("Create table: %w", err)
		}
	}
	log.Println("All table created successfully!")
	return nil
}
