package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/kongthap-code/kongthuay/pkg/db"
	"github.com/kongthap-code/kongthuay/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	// "github.com/golang-jwt/jwt/v4"
)

func main() {
	r := gin.Default()

	db.Init()

	r.GET("/CreateUser", func(c *gin.Context) {
		user := model.NewUser("game")
		mgm.Coll(user).Create(user)

		c.JSON(200, gin.H{
			"user": user.Name,
		})
	})

	r.GET("/GetUser/:name", func(c *gin.Context) {
		user := &model.User{}
		name,_ := c.Params.Get("name")

		err := mgm.Coll(user).First(bson.M{"name": name}, user)
		if err != nil {
			c.JSON(404, gin.H{"error":err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"user": user.Name,
		})
	})

	r.Run()
}
