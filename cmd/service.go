package cmd

import (
	"github.com/maxzycon/SIB-Golang-Assigment-2/pkg/dto"
	"github.com/maxzycon/SIB-Golang-Assigment-2/pkg/model"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func (s *Service) GetAll() (resp []*dto.OrderResponse, err error) {
	data := []*model.Order{}
	if err = s.db.Preload("Item").Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.OrderResponse, 0)
	for _, row := range data {
		tempParent := &dto.OrderResponse{
			ID:                row.ID,
			CreatedAt:         row.CreatedAt,
			UpdatedAt:         row.UpdatedAt,
			CustomerName:      row.CustomerName,
			OrderItemResponse: make([]dto.OrderItemResponse, 0),
		}

		for _, v := range row.Item {
			tempParent.OrderItemResponse = append(tempParent.OrderItemResponse, dto.OrderItemResponse{
				ID:          v.ID,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				ItemCode:    v.ItemCode,
				Description: v.Description,
				Qty:         v.Quantity,
				OrderID:     v.OrderID,
			})
		}

		resp = append(resp, tempParent)
	}

	return
}

func (s *Service) DeleteOrder(id uint) (err error) {
	data := &model.Order{Model: gorm.Model{
		ID: id,
	}}
	if err = s.db.First(data).Error; err != nil {
		return
	}

	if err = s.db.Delete(&model.Order{Model: gorm.Model{
		ID: id,
	}}).Error; err != nil {
		return
	}

	return
}

func (s *Service) CreateOrder(payload *dto.PaloadOrder) (err error) {
	data := &model.Order{Model: gorm.Model{
		CreatedAt: payload.OrderAt,
	}, CustomerName: payload.CustomerName}

	for _, v := range payload.Items {
		data.Item = append(data.Item, model.Item{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Qty,
		})
	}

	if err = s.db.Create(data).Error; err != nil {
		return
	}

	return
}

func (s *Service) UpdateOrder(id uint, payload *dto.PaloadOrder) (resp *dto.OrderResponse, err error) {
	// --- delete
	if err = s.db.Delete(&model.Order{Model: gorm.Model{
		ID: id,
	}}).Error; err != nil {
		return
	}

	// --- new create
	data := &model.Order{Model: gorm.Model{
		CreatedAt: payload.OrderAt,
	}, CustomerName: payload.CustomerName}

	for _, v := range payload.Items {
		data.Item = append(data.Item, model.Item{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Qty,
		})
	}

	if err = s.db.Create(data).Error; err != nil {
		return
	}

	row := &model.Order{Model: gorm.Model{
		ID: data.ID,
	}}
	if err = s.db.Preload("Item").First(row).Error; err != nil {
		return
	}

	// --- get new
	resp = &dto.OrderResponse{
		ID:                row.ID,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
		CustomerName:      row.CustomerName,
		OrderItemResponse: make([]dto.OrderItemResponse, 0),
	}

	for _, v := range row.Item {
		resp.OrderItemResponse = append(resp.OrderItemResponse, dto.OrderItemResponse{
			ID:          v.ID,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Qty:         v.Quantity,
			OrderID:     v.OrderID,
		})
	}

	return
}
