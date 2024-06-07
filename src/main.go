package main

import (
	"fmt"
	"os"
	"xampp/htdocs/go/graphql/handler/graphql"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}

func main() {
	fmt.Println("HelloWorld!")
	db, err := getDb()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Println(db)
	r := gin.Default()
	grp := r.Group("/test")
	{
		grp.GET("/version", graphql.Version)
	}
	port := ":3010"
	r.Run(port)
}

func getDb() (*Db, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/graphql_sample?charset=utf8mb4&parseTime=True&loc=Local"
	dbR, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db := &Db{}
	db.db = dbR
	return db, err
}
