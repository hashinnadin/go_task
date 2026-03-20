package main

import "github.com/gin-gonic/gin"

var sessions = map[string]string{}

func hy() {

}
func login(c *gin.Context) {

	var user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid"})
		return
	}

	if user.Id == 1 && user.Name == "hashin" {

		sessionId := "abcd123"

		sessions[sessionId] = user.Name

		c.SetCookie("session", sessionId, 3600, "/", "localhost", false, true)

		c.JSON(200, gin.H{
			"message": "Successfully Logged",
		})
		return
	}

}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		sessionId, err := c.Cookie("session")

		if err != nil {
			c.JSON(401, gin.H{"error": "please login"})
			c.Abort()
			return
		}
		_, exists := sessions[sessionId]

		if !exists {
			c.JSON(401, gin.H{"error": "invalid session id"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func logout(c *gin.Context) {

	sessionId, _ := c.Cookie("session")

	delete(sessions, sessionId)

	c.SetCookie("session", "", -1, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"message": "logged out",
	})
}

func Dashboard(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome Dashboard",
	})
}

func main() {

	r := gin.Default()

	r.POST("/login", login)

	r.GET("/dashboard", AuthMiddleware(), Dashboard)

	r.GET("/logout", logout)

	r.Run(":6006")
}
