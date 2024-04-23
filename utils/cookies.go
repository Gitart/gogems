package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Write cookies in hours
func CokieWrite(c echo.Context, name, value string, tm time.Duration) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value

	// если не указывать разделитель будет
	// записыываться для каждой страницы отдельно
	// Разделитель для общего использования
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(tm * time.Hour)

	c.SetCookie(cookie)
}

// Read Cokies
func CokieRead(c echo.Context, name string) string {
	cookie, err := c.Cookie(name)

	if err != nil {
		// "ERROR:  read cookies"
		return err.Error()
	}

	// fmt.Println(cookie.Name)
	// fmt.Println(cookie.Value)
	return cookie.Value
}

// Read all cookies
func ReadAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "Read all the cookies")
}

// Work with cookies
func TestingwriteCookie(c echo.Context) error {

	CokieWrite(c, "Iddoc", "12334", 123)
	CokieWrite(c, "names", "Iddoc", 155)

	s := CokieRead(c, "names")

	return c.String(http.StatusOK, s)
}

// Read cookies
func TestingreadCookie(c echo.Context) error {
	s := CokieRead(c, "names")
	return c.String(http.StatusOK, s)
}
