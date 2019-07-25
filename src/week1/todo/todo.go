package todo

import (
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
	db, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/todolist?parseTime=true")
	if err != nil {
		log.Fatal("failed to connect db")
	}
	db.LogMode(true)

	defer db.Close()
	err = db.AutoMigrate(Todo{}).Error
	if err != nil {
		log.Fatal("failed to migrate table todos")
	}

	router := gin.Default()

	router.GET("/todos", listTodos)
	router.POST("/todos", createTodo)
	router.Run(":8088")
}

func listTodos(c *gin.Context) {
	var todos []Todo
	err := db.Find(&todos).Error

	if err != nil {
		c.String(500, "failed to list todolist")
		return
	}

	c.JSON(200, todos)
}

func createTodo(c *gin.Context) {
	var argument struct {
		Title string
	}

	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	todo := Todo{
		Title: argument.Title,
	}

	err = db.Create(&todo).Error
	if err != nil {
		c.String(500, "failed to create new todo")
		return
	}

	c.JSON(200, todo)
}
