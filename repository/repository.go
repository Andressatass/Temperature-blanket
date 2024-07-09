package repository

type WheaterRepositoryInterface interface {
	GetTemperatureByIdAndDate(id string, date string) (error, []WheatherInfo)
	DeleteByDate(id string, date string) error
	Create(wheatherInfo WheatherInfo) error
}
