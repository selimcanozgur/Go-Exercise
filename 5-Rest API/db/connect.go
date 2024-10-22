package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB


func Connect() {

	var err error

	err = godotenv.Load()
	
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	DB_URI := os.Getenv("DB_URI")


	DB, err = sql.Open("mysql", DB_URI)

	if err != nil {
		panic("Veritabanı bağlantısı sağlanamadı")
	}
	

	DB.SetMaxOpenConns(10) // Aynı anda açık olabilecek maksimum veritabanı bağlantısı
	DB.SetMaxIdleConns(5) // Bağlantı havuzunda kullanılmadan bekletilebilecek maksimum bağlantı sayısı

	if err := DB.Ping(); err != nil {
		log.Fatal("Veritabanı bağlantısı doğrulanamadı:", err)
	}

	createTables()
}


func createTables() {
	createEventTables := `
	CREATE TABLE IF NOT EXISTS events (
	  id INT AUTO_INCREMENT PRIMARY KEY,
	  name VARCHAR(255) NOT NULL,
	  description TEXT NOT NULL,
	  location VARCHAR(255) NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INT
	)`

	_, err := DB.Exec(createEventTables)
	if err != nil {
		log.Fatal("Tablo oluşturulamadı:", err)
	}
}