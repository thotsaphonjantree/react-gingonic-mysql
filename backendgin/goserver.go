package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student struct {
	StudentID   int    `gorm:"primary_key" json:"student_id"`
	StudentCode string `json:"student_code"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MajorID     int    `json:"major_id"`
	Major       Major  `gorm:"foreignkey:major_id;association_foreignkey:major_id"`
}

type Major struct {
	MajorID   int    `gorm:"primary_key" json:"major_id"`
	MajorName string `json:"major_name"`
}

type Handler struct {
	DB *gorm.DB
}

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	h := Handler{}
	h.Initialize()
	r.GET("/major", h.getAllMajor)
	r.GET("/student", h.getAllStudent)
	r.GET("/major/:id", h.getMajor)
	r.GET("/studentbymajor/:id", h.getStudentByMajor)
	r.GET("/student/:id", h.getStudent)
	r.POST("/student", h.insertStudent)
	r.PUT("/student/:id", h.UpdateStudent)
	r.DELETE("/student/:id", h.DeleteStudent)
	return r
}

func (h *Handler) Initialize() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/student_register?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&Student{}, &Major{})
	db.Model(&Student{}).AddForeignKey("major_id", "major(major_id)", "CASCADE", "CASCADE")
	db.Model(&Major{}).AddIndex("index_major_id_name", "major_id", "major_name")

	h.DB = db
}

func (h *Handler) getAllMajor(c *gin.Context) {
	var majors []Major
	h.DB.Find(&majors)
	c.JSON(http.StatusOK, majors)
}

func (h *Handler) getAllStudent(c *gin.Context) {
	var student []Student
	if err := h.DB.Joins("JOIN major on student.major_id = major.major_id ").Preload("Major").Find(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}

func (h *Handler) getMajor(c *gin.Context) {
	id := c.Param("id")
	var major Major

	if err := h.DB.Find(&major, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, major)
}

func (h *Handler) getStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student

	if err := h.DB.Joins("JOIN major on student.major_id = major.major_id ").Preload("Major").Find(&student, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *Handler) insertStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&student).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student

	if err := h.DB.Find(&student, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&student).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student

	if err := h.DB.Find(&student, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if err := h.DB.Delete(&student).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) getStudentByMajor(c *gin.Context) {
	var student []Student
	id := c.Param("id")
	if err := h.DB.Joins("JOIN major on student.major_id = major.major_id ").Preload("Major").Where("student.major_id = ?", id).Find(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}
