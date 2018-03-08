package models

import (
	"time"
)

type (
	// Product describes a Product type
	Product struct {
		ID             uint       `json:"id" gorm:"primary_key"`
		CreatedAt      time.Time  `json:"created-at"`
		UpdatedAt      time.Time  `json:"updated-at"`
		DeletedAt      *time.Time `json:"deleted-at" sql:"index"`
		Name           string     `json:"name"`
		Type           int        `json:"type"`
		Price          int        `json:"price"`
		Url            string     `json:"url"`
		Image          string     `json:"image"`
		BreakpointUp   int        `json:"breakpoint-up" gorm:"column:breakpoint_up"`
		BreakpointDown int        `json:"breakpoint-down" gorm:"column:breakpoint_down"`
		Amount         int        `json:"amount"`
	}
	// ResProduct represents a formatted Product
	ResProduct struct {
		ID             uint   `json:"id"`
		Name           string `json:"name"`
		Type           int    `json:"type"`
		Price          int    `json:"price"`
		Url            string `json:"url"`
		Image          string `json:"image"`
		BreakpointUp   int    `json:"breakpoint-up"`
		BreakpointDown int    `json:"breakpoint-down"`
		Amount         int    `json:"amount"`
	}
)
