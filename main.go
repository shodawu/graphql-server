package main

import (
	"fmt"
	"go-users/entities"
	"go-users/graph"
	"go-users/graph/generated"
	"go-users/graph/model"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMain *gorm.DB

func initDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	var users []model.User

	db.Model(&users).Find(&users)

	if len(users) == 0 {
		users := []model.User{
			{Name: "UserOne", Email: "user.one@example.com", Status: entities.UserConfirmed},
			{Name: "UserTwo", Email: "user.two@example.com", Status: entities.UserConfirmed},
			{Name: "UserThree", Email: "user.three@example.com", Status: entities.UserConfirmed},
			{Name: "UserFour", Email: "user.four@example.com", Status: entities.UserBlocked},
			{Name: "UserFive", Email: "user.five@example.com", Status: entities.UserRegistered},
		}

		for _, usr := range users {
			er := dbMain.Create(&usr).Error
			if er != nil {
				fmt.Println("creating users error ", er)
			}
		}
	}

	return db
}

func main() {
	dbMain = initDB()

	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/users", getUsersHandler)
	r.Run(":8080")

}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: dbMain,
	}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func getUsersHandler(c *gin.Context) {
	var users []model.User
	er := dbMain.Model(&users).Find(&users).Error
	if er != nil {
		c.JSON(http.StatusBadRequest, er)
		return

	}
	c.JSON(http.StatusOK, &users)
}
