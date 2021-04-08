package example

import "github.com/gin-gonic/gin"

func Default() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    "Hello world",
		})
	})

	r.Run()
}

func RouterTest() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"name": name,
		})
	})

	router.POST("/user")

	router.Run(":8081")
}

func QueryParam() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "chui")
		lastname := c.Query("lastname")

		c.JSON(200, gin.H{
			"firstname": firstname,
			"lastname":  lastname,
		})
	})

	router.Run()
}

func PostParam() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "18")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	router.Run()
}

func GroupRouter() {
	router := gin.Default()

	v1 := router.Group("/v1")

	{
		v1.GET("/login", func(c *gin.Context) {})
		v1.GET("/submit", func(c *gin.Context) {})
		v1.GET("/read", func(c *gin.Context) {})
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/login", func(c *gin.Context) {})
		v2.GET("/submit", func(c *gin.Context) {})
		v2.GET("/read", func(c *gin.Context) {})
	}

	router.Run()
}
