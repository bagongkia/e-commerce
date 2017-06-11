package services

import (
	"../models"
)

type itemCategoryDAO interface {
	Query(db models.DB)([]models.ItemCategory, error)
}

type ItemCategoryService struct {
	dao itemCategoryDAO
}

func NewItemCategoryService(dao itemCategoryDAO) *ItemCategoryService {
	return &ItemCategoryService{dao}
}

func (s *ItemCategoryService) Query(db models.DB) ([]models.ItemCategory, error) {
	return s.dao.Query(db)
}
