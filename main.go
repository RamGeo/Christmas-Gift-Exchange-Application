package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/gin-gonic/gin"
    "net/http"
)

// Mock gift suggestions
var gifts = []string{
    "Book",
    "Toy",
    "Watch",
    "Perfume",
    "Gift Card",
    "Chocolate",
    "Flowers",
}

// Function to assign a random gift
func assignGift(name string) string {
    rand.Seed(time.Now().UnixNano())
    return gifts[rand.Intn(len(gifts))]
}

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("templates/*")
    r.Static("/static", "./static")

    // Route for the story page
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "story.html", nil)
    })

    // Route for the gift assignment page
    r.GET("/assignment", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    // Route for assigning gifts
    r.POST("/assign", func(c *gin.Context) {
        name := c.PostForm("name")
        assignedGift := assignGift(name)
        c.HTML(http.StatusOK, "result.html", gin.H{"name": name, "gift": assignedGift})
    })

    // Start the server on port 8080 and print a message
    fmt.Println("Server running at http://localhost:8080")
    r.Run(":8080")
}
