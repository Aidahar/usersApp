package transport

import (
	"net/http"
	"strconv"

	"github.com/Aidahar/filmsApi/ewrap"
	"github.com/Aidahar/filmsApi/internal/domain"
	"github.com/Aidahar/filmsApi/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	ewrap.LogFatal(err)

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.GetUserById(id)
	ewrap.LogFatal(err)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) AddUser(c echo.Context) error {
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

func (h *Handler) UpdateUser(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	domain.Users[id].Name = u.Name
	return c.JSON(http.StatusOK, domain.Users[id])
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(domain.Users, id)
	return c.NoContent(http.StatusNoContent)
}
