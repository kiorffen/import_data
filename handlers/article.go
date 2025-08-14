package handlers

import (
	"net/http"
	"import_data/models"
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article models.Article
	
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := article.Save(); err != nil {
		if err.Error() == "article with URL already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Article with this URL already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save article"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Article created successfully",
		"id":      article.ID,
	})
}