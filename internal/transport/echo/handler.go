package transport

import (
	"net/http"
	"strconv"

	"github.com/Aidahar/filmsApi/internal/domain"
	"github.com/Aidahar/filmsApi/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service service.UserService
}

func NewHandler(service service.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/users", h.GetAllUsers)
	e.GET("/users/:id", h.GetUser)
	e.POST("/users", h.AddUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)

	return e
}

func (h *Handler) GetAllUsers(c echo.Context) error {
	users, err := h.service.GetAllUsers()
	logrus.Fatalf("can`t take all users", err)

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.GetUserById(id)
	logrus.Fatalf("can`t take user", err)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) AddUser(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(&u); err != nil {
		logrus.Fatalf("wrong data", err)
	}
	if err := h.service.AddUser(*u); err != nil {
		logrus.Fatalf("cant add user", err)
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		logrus.Fatalf("wrong data", err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Fatalf("wrong id", err)
	}
	if err := h.service.UpdateUser(id, *u); err != nil {
		logrus.Fatalf("cant update user", err)
	}
	return c.JSON(http.StatusOK, u.ID)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Fatalf("wrong id", err)
	}
	if err := h.service.DeleteUser(id); err != nil {
		logrus.Fatalf("cant delete user", err)
	}
	return c.NoContent(http.StatusNoContent)
}
