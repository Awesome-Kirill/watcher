package dto

import "time"

type Info struct {
	IsAlive      bool
	ResponseTime time.Duration
}

type InfoWithName struct {
	Info
	Name string
}
