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
	StudentID   int    `gorm:"primary_key" json:"student_id"`
	StudentCode string `json:"student_code"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MajorID     int    `json:"major_id"`
	Major       Major  `gorm:"foreignkey:major_id;association_foreignkey:major_id"`

	//Major Major `gorm:"ForeignKey:major_id" json:"major"`
	//CustId      int  `json:"student_id"`

}

type Major struct {
	MajorID   int    `gorm:"primary_key" json:"major_id"`
	MajorName string `json:"major_name"`
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
	router.GET("/major", getMajors)
	router.GET("/major/:id", getMajor)
	router.GET("/student", getStudents)
	router.GET("/studentbymajor/:id", getStudentByMajor)
	router.GET("/student/:id", getStudent)
	router.POST("/student", insertStudent)
	router.PUT("/student/:id", UpdateStudent)
	router.DELETE("/student/:id", DeleteStudent)
	router.Run()
}

func getMajors(c *gin.Context) {
	var majors []Major
	db.Find(&majors)
	c.JSON(http.StatusOK, majors)
}

func getStudents(c *gin.Context) {
	var student []Student

	if err := db.Joins("JOIN major on student.major_id = major.major_id ").Preload("Major").Find(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}



func getStudentByMajor(c *gin.Context) {
	var student []Student
	id := c.Param("id")
	if err := db.Joins("JOIN major on student.major_id = major.major_id ").Preload("Major").Where("student.major_id = ?",id).Find(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}


func getMajor(c *gin.Context) {
	id := c.Param("id")
	var major Major

	if err := db.Find(&major, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, major)
}


func getStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student

	if err := db.Joins("JOIN major on student.major_id = major.major_id ").Preload("Major").Find(&student, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, student)
}

func insertStudent(c *gin.Context) {
	var student Student

	// c.Bind(&student)
	// db.Create(&student)
	// c.JSON(201, gin.H{"success": student})
	// db.Save(&student)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := db.Save(&student).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student

	if err := db.Find(&student, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := db.Save(&student).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, student)
}


func DeleteStudent(c *gin.Context){
	id := c.Param("id")
	var student Student

	if err := db.Find(&student, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if err := db.Delete(&student).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
