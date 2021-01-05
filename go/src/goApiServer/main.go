package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 익명함수
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"time"
)

type (
	user struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}
)

func createUser(c echo.Context) error {

	db, err := sql.Open("mysql", "root:qwd@tcp(127.0.0.1:3306/testdb")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)

	var seq int
	err = db.QueryRow("SELECT count(id) FROM user ").Scan(&seq)
	if err != nil{
		panic(err)
	}
	u := &user{
		ID: seq + 1,
	}

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusCreated, err)
		return err
	}
	db.Exec("INSERT INTO user VALUES (?,?)",u.ID,u.Name)

	defer db.Close()
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	db, err := sql.Open("mysql", "root:qwd@tcp(127.0.0.1:3306/testdb")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	id, _ := strconv.Atoi(c.Param("id"))
	db.QueryRow()

	defer db.Close()
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!@#!@#!#@#")
	})
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1343"))
}
