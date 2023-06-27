package handler

import (
	"net/http"
	"watcher/internal/dto"

	"github.com/labstack/echo/v4"
)

type Name struct {
	cacheStore CacheStore
}

type CacheStore interface {
	GetURL(url string) (dto.Info, error)
	GetMax() dto.InfoWithName
	GetMin() dto.InfoWithName
}

func New(cacheStore CacheStore) *Name {
	return &Name{cacheStore: cacheStore}
}
func (n *Name) GetMax(ctx echo.Context) error {
	max := n.cacheStore.GetMax()
	return ctx.JSONPretty(http.StatusOK, max, "  ")
}

func (n *Name) GetMin(ctx echo.Context) error {
	min := n.cacheStore.GetMin()
	return ctx.JSONPretty(http.StatusOK, min, "  ")
}

type GetSiteStatResponse struct {
}

func (n *Name) GetSiteStat(ctx echo.Context) error {
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
