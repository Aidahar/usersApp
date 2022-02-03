package transport

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Aidahar/filmsApi/ewrap"
	"github.com/Aidahar/filmsApi/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Users interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}

type Handler struct {
	usersRepos Users
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/users", h.GetAllUsers)
	//	e.GET("/users/:id", GetUser)
	//	e.POST("/users", AddUser)
	//	e.PUT("/users/:id", UpdateUser)
	//	e.DELETE("/users/:id", DeleteUser)

	return e
}

func (h *Handler) GetAllUsers(c echo.Context) error {
	users, err := h.usersRepos.GetAllUsers(context.TODO())
	ewrap.LogFatal(err)

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, domain.Users[id])
}

func AddUser(c echo.Context) error {
	u := &domain.User{
		ID: domain.Seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	domain.Users[u.ID] = u
	domain.Seq++
	return c.JSON(http.StatusCreated, u)
}

func UpdateUser(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	domain.Users[id].Name = u.Name
	return c.JSON(http.StatusOK, domain.Users[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(domain.Users, id)
	return c.NoContent(http.StatusNoContent)
}
