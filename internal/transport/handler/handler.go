package handler

import (
	"net/http"
	"watcher/internal/cache"

	"github.com/labstack/echo/v4"
)

type Name struct {
	cacheStore CacheStore
}

type CacheStore interface {
	GetUrl(url string) (cache.Info, error)
	GetMax() cache.InfoWithName
	GetMin() cache.InfoWithName
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

	res, err := n.cacheStore.GetUrl(site)

	if err != nil {
		return ctx.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	return ctx.JSONPretty(http.StatusOK, res, "  ")
}
