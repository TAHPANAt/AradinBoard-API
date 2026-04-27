package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// โครงสร้างข้อมูลของกระทู้
type Thread struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// จำลองฐานข้อมูลแบบ In-memory
var threads = []Thread{
	{ID: "1", Title: "ยินดีต้อนรับสู่ AradinBoard", Content: "กระทู้แรกของระบบ แวะมาทักทายกันได้ครับ!"},
}

func main() {
	r := gin.Default()

	// Endpoint 1: สำหรับให้ Prometheus เช็คว่าแอปยังรอดอยู่ไหม (Health Check)
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "AradinBoard is running!"})
	})

	// Endpoint 2: ดึงข้อมูลกระทู้ทั้งหมด
	r.GET("/threads", func(c *gin.Context) {
		c.JSON(http.StatusOK, threads)
	})

	// Endpoint 3: สร้างกระทู้ใหม่
	r.POST("/threads", func(c *gin.Context) {
		var newThread Thread
		if err := c.BindJSON(&newThread); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		threads = append(threads, newThread)
		c.JSON(http.StatusCreated, newThread)
	})

	// รันเซิร์ฟเวอร์ที่พอร์ต 8080
	r.Run(":8080")
}