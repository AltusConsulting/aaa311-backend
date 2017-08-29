package models

type Roles struct {
	Name string `json:"name" binding:"required"`
}
