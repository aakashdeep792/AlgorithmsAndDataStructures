package model

import "time"

const (
	SMALL int = iota
	MEDIUM
	LARGE
)

type Product struct {
	ID         int       `json:"id"`
	Name       string    `json:"Name"`
	Decription string    `json:"Description"`
	CreatedAt  time.Time `json:"CreatedAt"`
	Size       int       `json:"size"`
}
