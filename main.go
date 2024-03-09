package main

import (
	"database/sql"
	"fmt"
	"log"
        "os"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

type AudioFile struct {
	ID       int
	UserID   int
	PhraseID int
	FilePath string
}

func setupDatabase() {

        dbURI := os.Getenv("DATABASE_URL")
        log.Printf("Setting up DB. URI is %s\n", dbURI)
	// Set up a PostgreSQL database connection
	var err error
        log.Printf("Opening sql connection with DBURI %s\n", dbURI)
	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
        log.Printf("Executing create table with DBURI %s\n", dbURI)
	// Create the "audio_files" table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS audio_files (
			id SERIAL PRIMARY KEY,
			user_id INTEGER,
			phrase_id INTEGER,
			file_path TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	setupDatabase()

	r := gin.Default()

	r.POST("/audio/user/:user_id/phrase/:phrase_id", uploadAudio)
	r.GET("/audio/user/:user_id/phrase/:phrase_id/:audio_format", downloadAudio)

	err := r.Run(":80")
	if err != nil {
		log.Fatal(err)
	}
}

func uploadAudio(c *gin.Context) {
	userID := c.Param("user_id")
	phraseID := c.Param("phrase_id")

	file, err := c.FormFile("audio_file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	// Save the file to the server
	filePath := fmt.Sprintf("uploads/%s_%s.wav", userID, phraseID)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	// Save the file path to the database
	_, err = db.Exec("INSERT INTO audio_files (user_id, phrase_id, file_path) VALUES ($1, $2, $3)", userID, phraseID, filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "File uploaded successfully")
}

func downloadAudio(c *gin.Context) {
	userID := c.Param("user_id")
	phraseID := c.Param("phrase_id")
	audioFormat := c.Param("audio_format")

	var filePath string
	err := db.QueryRow("SELECT file_path FROM audio_files WHERE user_id = $1 AND phrase_id = $2", userID, phraseID).Scan(&filePath)
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	if audioFormat != "m4a" {
		c.String(http.StatusBadRequest, "Invalid audio format")
		return
	}

	c.File(filePath)
}
