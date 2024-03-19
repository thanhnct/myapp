package productmysql

import (
	"context"
	productdomain "myapp/module/product/domain"
)

func (repo MysqlRepository) CreateProduct(ctx context.Context, prod *productdomain.ProductCreateDTO) error {
	if err := repo.db.Table(prod.TableName()).Create(&prod).Error; err != nil {
		return err
	}
	return nil
}
