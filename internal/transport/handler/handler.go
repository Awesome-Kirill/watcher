package handler

import (
	"net/http"

	_ "watcher/docs"
	"watcher/internal/dto"
	"watcher/internal/stat"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	cacheStore CacheStore
	stat       stat.Stat
}

type CacheStore interface {
	GetURL(url string) (dto.Info, error)
	GetMax() dto.InfoWithName
	GetMin() dto.InfoWithName
}

func New(cacheStore CacheStore, adminKey string) *echo.Echo {
	server := &Server{
		cacheStore: cacheStore,
		stat:       stat.Stat{},
	}
	e := echo.New()

	e.GET("/stat/min", server.GetMin)
	e.GET("/stat/max", server.GetMax)
	e.GET("/stat/:id/site", server.GetSiteStat)

	e.GET("/admin/stat", server.GetStat, middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == adminKey, nil
	}))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}

type GetMaxResponse struct {
	Name    string  `json:"name"`
	Seconds float64 `json:"seconds"`
}

// GetMax godoc
// @Summary     Return most slow site
// @Description Return most slow site
// @Produce     json
// @Success     200       {object} GetMaxResponse
// @Router      /stat/max [get]
func (n *Server) GetMax(ctx echo.Context) error {
	defer n.stat.IncMax()
	max := n.cacheStore.GetMax()

	response := GetMaxResponse{
		Name:    max.Name,
		Seconds: max.ResponseTime.Seconds(),
	}
	return ctx.JSONPretty(http.StatusOK, response, "  ")
}

type GetMinResponse struct {
	Name    string  `json:"name"`
	Seconds float64 `json:"seconds"`
}

// GetMin godoc
// @Summary     Return most fasts site
// @Description Return most fasts site
// @Produce     json
// @Success     200       {object} GetMinResponse
// @Router      /stat/min [get]
func (n *Server) GetMin(ctx echo.Context) error {
	defer n.stat.IncMin()

	min := n.cacheStore.GetMin()

	response := GetMinResponse{
		Name:    min.Name,
		Seconds: min.ResponseTime.Seconds(),
	}
	return ctx.JSONPretty(http.StatusOK, response, "  ")
}

// todo
type GetStatResponse struct {
	Min uint64 `json:"/stat/min"`
	Max uint64 `json:"/stat/max"`
	URL uint64 `json:"/stat/{id}/site"`
}

// GetStat godoc
// @Summary     Return statistic
// @Tags        Administration
// @Description Return most fasts site
// @Security ApiKeyAuth
// @Produce     json
// @Success     200       {object} GetMinResponse
// @Router      /admin/stat [get]
func (n *Server) GetStat(ctx echo.Context) error {
	stats := n.stat.Stat()
	response := GetStatResponse{
		Min: stats.Min,
		Max: stats.Max,
		URL: stats.URL,
	}
	return ctx.JSONPretty(http.StatusOK, response, "  ")
}

type GetSiteStatResponse struct {
	Name    string  `json:"name"`
	Seconds float64 `json:"seconds"`
	IsAlive bool    `json:"IsAlive"`
}

// GetSiteStat godoc
// @Summary     Return most fasts site
// @Description Return most fasts site
// @Param id   path string true "Site ID"
// @Produce     json
// @Success     200       {object} GetSiteStatResponse
// @Router      /stat/{id}/site [get]
func (n *Server) GetSiteStat(ctx echo.Context) error {
	defer n.stat.IncURL()

	site := ctx.Param("id")

	if site == "" {
		return ctx.JSONPretty(http.StatusBadRequest, struct {
			message string
		}{
			message: "id is null",
		}, "  ")
	}

	data, err := n.cacheStore.GetURL(site)

	if err != nil {
		return ctx.JSONPretty(http.StatusBadRequest, GetSiteStatResponse{}, "  ")
	}
	response := GetSiteStatResponse{
		Name:    site,
		Seconds: data.ResponseTime.Seconds(),
		IsAlive: data.IsAlive,
	}
	return ctx.JSONPretty(http.StatusOK, response, "  ")
}
