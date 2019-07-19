package dlldb

import (
	"github.com/jinzhu/gorm"

	dgorm "github.com/dllgo/dlldb/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

func NewGormRepository() *GormRepository {
	return &GormRepository{
		DB: dgorm.MustDB(),
	}
}
func (gormr *GormRepository) One(model interface{}, query interface{}, args ...interface{}) error {
	if err := gormr.DB.Where(query, args...).First(model).Error; err != nil {
		return err
	}
	return nil
}
func (gormr *GormRepository) List(model interface{}, order string, limit, offset int, query interface{}, args ...interface{}) error {
	db := gormr.DB.Where(query, args...)
	if len(order) != 0 {
		db = db.Order(order)
	}
	if offset > 0 {
		db = db.Offset(offset)
	}
	if limit > 0 {
		db = db.Limit(limit)
	}
	if err := db.Find(model).Error; err != nil {
		return err
	}
	return nil
}

// Update All Fields
func (gormr *GormRepository) Save(model interface{}) error {
	if err := gormr.DB.Save(model).Error; err != nil {
		return err
	}
	return nil
}

// Update selected Fields, if attrs is an object, it will ignore default value field; if attrs is map, it will ignore unchanged field.
func (gormr *GormRepository) Update(model interface{}, attrs interface{}, query interface{}, args ...interface{}) error {
	if err := gormr.DB.Model(model).Where(query, args...).Update(attrs).Error; err != nil {
		return err
	}
	return nil
}
func (gormr *GormRepository) Count(model interface{}, query interface{}, args ...interface{}) int {
	var count int
	if err := gormr.DB.Model(model).Where(query, args...).Count(&count).Error; err != nil {
		return 0
	}
	return count
}
func (gormr *GormRepository) Delete(model interface{}, query interface{}, args ...interface{}) error {
	if err := gormr.DB.Where(query, args...).Delete(model).Error; err != nil {
		return err
	}
	return nil
}
func (gormr *GormRepository) ExecSql(model interface{}, sql string, args ...interface{}) error {
	if err := gormr.DB.Raw(sql, args...).Scan(model).Error; err != nil {
		return err
	}
	return nil
}
