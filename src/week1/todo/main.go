package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}

var db *gorm.DB

func main() {

	var err error

	db, err = gorm.Open("mysql", "root@tcp(127.0.0.1)/todolist?parseTime=true")

	fmt.Println("error: ", err)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	db.LogMode(true)

	defer db.Close()

	err = db.AutoMigrate(Todo{}).Error
	if err != nil {
		log.Fatal("Failed to migrate db")
	}

	router := gin.Default()

	router.GET("/todos", listTodos)
	router.POST("/todos", createTodo)
	router.Run(":8080")
}

func listTodos(c *gin.Context) {

	var todos []Todo

	err := db.Find(&todos).Error

	if err != nil {
		c.String(500, "Failed to get todo list")
	}

	c.JSON(200, todos)
}

func createTodo(c *gin.Context) {
	var arg struct {
		Title string
	}

	err := c.BindJSON(&arg)
	if err != nil {
		c.String(400, "Invalid param")
	}

	todo := Todo{
		Title: arg.Title,
	}

	err = db.Create(&todo).Error
	if err != nil {
		c.String(500, "Fail to create new todo")
		return
	}

	c.JSON(200, todo)
}
