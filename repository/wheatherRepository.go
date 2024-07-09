package repository

import "gorm.io/gorm"

type WheatherRepository struct {
	repository WheaterRepositoryInterface
}

func NewWheatherRepository(repo WheaterRepositoryInterface) WheatherRepository {
	return WheatherRepository{
		repository: repo,
	}
}

type WheatherInfo struct {
	gorm.Model
	ID          string
	Date        string
	Temperature float32
}

// func Init() {
// 	db, err := gorm.Open(mysql.Open())
// }
