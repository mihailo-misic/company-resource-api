package models

import "github.com/jinzhu/gorm"

type (
	// Product describes a Product type
	Product struct {
		gorm.Model
		Name           string `json:"name"`
		Type           int    `json:"type"`
		Price          int    `json:"price"`
		Url            string `json:"url"`
		Image          string `json:"image"`
		BreakpointUp   int    `gorm:"column:breakpoint_up"json:"breakpointUp"`
		BreakpointDown int    `gorm:"column:breakpoint_down"json:"breakpointDown"`
		Amount         int    `json:"amount"`
	}
	// ResProduct represents a formatted Product
	ResProduct struct {
		ID             uint   `json:"id"`
		Name           string `json:"name"`
		Type           int    `json:"type"`
		Price          int    `json:"price"`
		Url            string `json:"url"`
		Image          string `json:"image"`
		BreakpointUp   int    `json:"breakpointUp"`
		BreakpointDown int    `json:"breakpointDown"`
		Amount         int    `json:"amount"`
	}
)
