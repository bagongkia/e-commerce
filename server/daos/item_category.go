package daos

import (
	"../models"
)

type ItemCategoryDAO struct{}

func NewItemCategoryDAO() *ItemCategoryDAO {
	return &ItemCategoryDAO{}
}

func (dao *ItemCategoryDAO) Query(db models.DB) ([]models.ItemCategory, error) {
	rows, err := db.Query("SELECT * FROM item_category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itemCategories := []models.ItemCategory{}
	for rows.Next() {
		row := new(models.ItemCategory)
		err := rows.Scan(&row.ItemCategoryId,
			&row.ItemCategoryName,
			&row.CreatedBy,
			&row.CreatedTime,
			&row.ModifiedBy,
			&row.ModifiedTime)
		if err != nil {
			return nil, err
		}
		itemCategories = append(itemCategories, *row)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return itemCategories, nil
}