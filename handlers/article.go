package handlers

import (
	"import_data/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article models.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": err.Error(),
			"id":      article.ID,
		})
		return
	}

	if err := article.Save(); err != nil {
		if err.Error() == "article with URL already exists" {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Article with this URL already exists",
				"id":      article.ID,
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "Failed to save Article",
			"id":      article.ID,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Article created successfully",
		"id":      article.ID,
	})
}
