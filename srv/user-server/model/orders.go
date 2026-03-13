package model

import (
	__ "github.com/yuhang-jieke/orderai/srv/proto"
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(30);comment:订单名称"`
	Num   int     `gorm:"type:int(11);comment:订单数量"`
	Price float64 `gorm:"type:decimal(10,2);comment:订单金额"`
}

func (o *Orders) OrderAdd(db *gorm.DB) error {
	return db.Create(&o).Error
}

func (o *Orders) UpdateId(db *gorm.DB, in *__.UpdateOrdersReq) error {
	return db.Model(&Orders{}).Where("id=?", in.Id).Update("price", in.Price).Error
}

func (o *Orders) DelId(db *gorm.DB, in *__.DelOrdersReq) interface{} {
	return db.Where("id=?", in.Id).Delete(&o).Error
}

func (o *Orders) GetId(db *gorm.DB, in *__.GetOrdersByIdReq) (Orders, error) {
	var list Orders
	err := db.Model(&Orders{}).Where("id=?", in.Id).First(&list).Error
	return list, err
}

func (o *Orders) Search(db *gorm.DB, in *__.SearchOrdersReq) ([]Orders, error) {
	var list []Orders
	if in.Page <= 0 || in.Page > 3 {
		in.Page = 1
	}
	if in.Size <= 0 || in.Size > 3 {
		in.Size = 1
	}
	tx := db.Model(Orders{})
	if in.Name != "" {
		tx = tx.Where("name like ?", "%"+in.Name+"%")
	}
	if in.Id > 0 {
		tx = tx.Where("id=?", in.Id)
	}
	if in.MinPrice > 0 && in.MaxPrice > 0 && in.MaxPrice > in.MinPrice {
		tx = tx.Where("price between ? and ?", in.MinPrice, in.MaxPrice)
	}
	offset := (in.Page - 1) * in.Size
	err := tx.Offset(int(offset)).Limit(int(in.Size)).Find(&list).Error
	return list, err
}
