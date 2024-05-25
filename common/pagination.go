package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit  int
	Offset int
	Page   int
}

type MetaPage struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	From        int `json:"from"`
	To          int `json:"to"`
}

type Meta struct {
	Pagination MetaPage `json:"pagination"`
}

func NewMeta(totalItems, perPage, page, offset, itemsCount int) *Meta {
	totalPages := (totalItems + perPage - 1) / perPage

	return &Meta{
		Pagination: MetaPage{
			Total:       totalItems,
			PerPage:     perPage,
			CurrentPage: page,
			LastPage:    totalPages,
			From:        offset + 1,
			To:          offset + itemsCount,
		},
	}
}

func ExtractPaginationParams(ctx *gin.Context, defaultPerPage, defaultPage int) (int, int) {
	perPage := defaultPerPage
	page := defaultPage

	if pp := ctx.Query("per_page"); pp != "" {
		if ppInt, err := strconv.Atoi(pp); err == nil {
			perPage = ppInt
		}
	}

	if p := ctx.Query("page"); p != "" {
		if pInt, err := strconv.Atoi(p); err == nil {
			page = pInt
		}
	}

	return perPage, page
}
