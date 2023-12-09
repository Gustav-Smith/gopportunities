package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utlizando as configurações Default do gin
	r := gin.Default()
	// Definindo uma Rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando a nossa api
	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
