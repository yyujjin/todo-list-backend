package main

import (
	"todo-list-backend/router"
	// "router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/ping", router.Pong)

	// todos -> todos 슬라이스에 있는 값을 전달
	// ['todo1', 'todo2', 'todo3']
	
	r.GET("/todos", router.Todos) 

	
	// post method/ todos -> todo4 요소 추가 후 todos slice response
	// r.POST("/todos", func(c *gin.Context) {
	// 	todos = append (todos,"todo4")
	// 	c.JSON(200, gin.H{
			
	// 		"todos": todos, 

	// 	})
	// })

	// // delete method / todos -> index 0 delete 
	// r.DELETE("/todos", func(c *gin.Context) {
	// 	todos =  todos[1:] 
	// 	c.JSON(200, gin.H{
			
	// 		"todos": todos, 

	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080
}