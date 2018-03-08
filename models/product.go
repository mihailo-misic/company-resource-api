package models

import (
	"time"
)

type (
	// Product describes a Product type
	Product struct {
		ID             uint64 `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
		Name           string `json:"name" binding:"required"`
		Type           int    `json:"type"`
		Price          int    `json:"price" binding:"required"`
		Url            string `json:"url"`
		Image          string `json:"image"`
		BreakpointUp   int    `json:"breakpoint-up" gorm:"column:breakpoint_up" binding:"required"`
		BreakpointDown int    `json:"breakpoint-down" gorm:"column:breakpoint_down" binding:"required"`
		Amount         int    `json:"amount"`

		CreatedAt time.Time  `json:"created-at"`
		UpdatedAt time.Time  `json:"updated-at"`
		DeletedAt *time.Time `json:"deleted-at" sql:"index"`
	}

	// ResProduct represents a formatted Product
	ResProduct struct {
		ID             uint64 `json:"id"`
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
