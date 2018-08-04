package main

import (
	"github.com/labstack/echo"
	"github.com/zmb3/spotify"
	"net/http"
)

type Handler struct {
	client *spotify.Client
}

func NewHandler(client *spotify.Client) Handler {
	return Handler{
		client: client,
	}
}

func (h *Handler) Current(c echo.Context) error {
	playing, err := h.client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, playing)
}
