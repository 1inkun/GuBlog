package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabaseConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	err = InitDatabase(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDatabase(db *sql.DB) error {
	query, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Printf("在读取数据库schema文件时出出错:%s", err)
		return err
	}
	_, err = db.Exec(string(query))
	if err != nil {
		log.Printf("在导入数据库结构时出出错:%s", err)
		return err
	}
	return nil
}
