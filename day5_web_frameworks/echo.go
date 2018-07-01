package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// MyContext ctx
type MyContext struct {
	echo.Context
}

// Foo foo
func (c *MyContext) Foo() {
	println("foo")
}

type (
	// User user
	User struct {
		Name  string `json:"name" form:"name" query:"name"`
		Email string `json:"email" form:"email" query:"email"`
	}

	// ValidateUser validate test
	ValidateUser struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		println("custom ctx")
		return func(c echo.Context) error {
			println("build ctx")
			cc := &MyContext{c}
			return h(cc)
		}
	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		println("error handling")
		return func(c echo.Context) error {
			println("err handler internal")
			// Extract the credentials from HTTP request header and perform a security
			// check

			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})
	//TODO 自定义异常处理，发生异常时返回err json
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.POST("/users", login)
	e.POST("/users/_validate", validateLogin)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	cc := c.(*MyContext)
	cc.Foo()
	return c.String(http.StatusOK, "Hello , World!")
}

/**
curl \
  -X POST \
  http://localhost:1323/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe","email":"joe@labstack"}'

curl \
  -X POST \
  http://localhost:1323/users \
  -d 'name=Joe' \
  -d 'email=joe@labstack.com'
*/
func login(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	c.Response().Before(func() {
		println("before response")
	})
	c.Response().After(func() {
		println("after response")
	})
	c.Response().After(func() {
		println("after response 2")
	})
	return c.JSON(http.StatusOK, u)
}

/**
curl \
  -X POST \
  http://localhost:1323/users/_validate \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe","email":"joe@invalid-domain"}'
*/
func validateLogin(c echo.Context) (err error) {
	u := new(ValidateUser)
	if err = c.Bind(u); err != nil {
		return
	}
	if err = c.Validate(u); err != nil {
		// return c.JSON(http.StatusBadRequest, err)
		return
	}
	return c.JSON(http.StatusOK, u)
}
