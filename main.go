package main

import (
	"github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
	"os"
	 "github.com/joho/godotenv"
	// "net/http"
	"log"
	// "bytes"
	// "encoding/json"
	// // "fmt"
	// "math/rand"
	s "github.com/myrachanto/firsttemp/support"
	con "github.com/myrachanto/firsttemp/controllers"
)
func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}	
	PORT := os.Getenv("PORT")
	//http.Handle("/", http.FileServer(http.Dir("client")))
	s.Configs()
	//r.Routes()
	e := echo.New()
	//e.Use(middleware.Static("/client"))
	e.Static("/", "client/public")
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	
	//echoGroupUseJWT := e.Group("/api/v1")
	//echoGroupUseJWT.Use(middleware.JWT([]byte(config.EncryptionKey)))
	//echoGroupNoJWT := e.Group("api/v1")
	//api v1/users: logged in usrs
	// //////generall routes
	// e.File("/favicon.ico", "client/public/favicon.ico")
	// e.File("/", "public/index.html")

	e.GET("/users/logout", con.Logout)
	///api/v1/users :public
	e.GET("/users", con.GetUsers)
	e.POST("/users/register", con.Register)
	e.POST("/users/login", con.Login)

	// e.GET("/home", c.GetHome)
	////////customers///////////
	e.GET("/customers", con.GetCustomers)
	e.GET("/customer/:id", con.GetCustomer)
	e.POST("/customers", con.CreateCustomers)
	e.POST("/customers/:id", con.UpdateCustomer)
	e.DELETE("/customers/:id", con.DeleteCustomer)
	e.Logger.Fatal(e.Start(PORT)) 
}