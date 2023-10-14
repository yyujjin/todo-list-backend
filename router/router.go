package router

import "github.com/gin-gonic/gin"

// pong 일반 함수로 변경
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Todos(c *gin.Context) {
	todos := []string { // TODO 전역변수로 만들기
		"todo1",
		"todo2", 
		"todo3",
	}

	c.JSON(200, gin.H{
		"todos": todos, 
	})
}