package models

import "time"

type ItemCategory struct {
	ItemCategoryId	 int		`json:"itemCategoryId"`
	ItemCategoryName string		`json:"itemCategoryName"`
	CreatedBy	 string		`json:"createdBy"`
	CreatedTime	 time.Time	`json:"createdTime"`
	ModifiedBy	 string		`json:"modifiedBy"`
	ModifiedTime	 time.Time	`json:"modifiedTime"`
}