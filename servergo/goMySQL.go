package main

import (
	// "database/sql"
	"fmt"
	"net/http"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var e error

type Student struct {
	StudentID   int  `gorm:"primary_key" json:"student_id"`
	StudentCode string  `json:"student_code"`
	FirstName    string `json:"first_name"`
	LastName	string `json:"last_name"`
	MajorID      int  `json:"-"`
	Major 		Major `gorm:"foreignkey:major_id;association_foreignkey:major_id"`

}

type Major struct {

	MajorID   int       `gorm:"primary_key" json:"major_id"`
	MajorName string    `json:"major_name"`
	//Students  []Student `gorm:"ForeignKey:major_id" json:"Students"`
}

func main() {
	db, e = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/student_register?charset=utf8&parseTime=True&loc=Local")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(&Student{}, &Major{})
	db.Model(&Student{}).AddForeignKey("major_id", "major(major_id)", "CASCADE", "CASCADE")
	db.Model(&Major{}).AddIndex("index_major_id_name", "major_id", "major_name")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/majors", getMajors)
	router.GET("/students", getStudents)
	router.POST("/student", insertStudent)
	router.Run(":8080")
}

func getMajors(c *gin.Context) {
	var majors []Major
	db.Find(&majors)
	c.JSON(http.StatusOK, majors)
}


func getStudents(c *gin.Context) {
	var student []Student
	db.Joins("JOIN major on student.major_id = major.major_id ").Preload("Major").Find(&student)
	c.JSON(http.StatusOK, student)
}


func insertStudent(c *gin.Context) {
	var student Student

	c.Bind(&student)
	db.Create(&student)
	c.JSON(400, gin.H{"success": student})
	db.Save(&student)
}
