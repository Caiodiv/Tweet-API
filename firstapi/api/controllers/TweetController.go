package controllers

import (
	"net/http"

	entities "api/api/entities"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (c *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.tweets)
}

func (c *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	c.tweets = append(c.tweets, *tweet)

	ctx.JSON(http.StatusOK, tweet)
}

func (c *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, tweet := range c.tweets {
		if tweet.ID == id {
			c.tweets = append(c.tweets[0:idx], c.tweets[idx+1:]...)
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})
}
