package handler

import (
	"net/http"
	"watcher/internal/dto"

	"github.com/labstack/echo/v4"
)

type Server struct {
	cacheStore CacheStore
}

type CacheStore interface {
	GetURL(url string) (dto.Info, error)
	GetMax() dto.InfoWithName
	GetMin() dto.InfoWithName
}

func New(cacheStore CacheStore) *Server {
	return &Server{cacheStore: cacheStore}
}
func (n *Server) GetMax(ctx echo.Context) error {
	max := n.cacheStore.GetMax()
	return ctx.JSONPretty(http.StatusOK, max, "  ")
}

func (n *Server) GetMin(ctx echo.Context) error {
	min := n.cacheStore.GetMin()
	return ctx.JSONPretty(http.StatusOK, min, "  ")
}

type GetSiteStatResponse struct {
}

func (n *Server) GetSiteStat(ctx echo.Context) error {
	site := ctx.Param("id")

	if site == "" {
		return ctx.JSONPretty(http.StatusBadRequest, struct {
			message string
		}{
			message: "id is null",
		}, "  ")
	}

	res, err := n.cacheStore.GetURL(site)

	if err != nil {
		return ctx.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	return ctx.JSONPretty(http.StatusOK, res, "  ")
}
