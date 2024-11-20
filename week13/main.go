package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Album struct {
    ID     int64   `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float32 `json:"price"`
}

func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "./recordings.db")
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

func getAlbums(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM album")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch albums"})
        return
    }
    defer rows.Close()

    var albums []Album
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to scan album"})
            return
        }
        albums = append(albums, alb)
    }

    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
    var newAlbum Album
    if err := c.BindJSON(&newAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    stmt, err := db.Prepare("INSERT INTO album(title, artist, price) VALUES (?, ?, ?)")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to prepare statement"})
        return
    }
    defer stmt.Close()

    res, err := stmt.Exec(newAlbum.Title, newAlbum.Artist, newAlbum.Price)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to execute statement"})
        return
    }

    albID, err := res.LastInsertId()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve last insert ID"})
        return
    }
    newAlbum.ID = albID
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
    initDB()
    router := gin.Default()
    router.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    })
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)
    router.Run("localhost:8080")
}