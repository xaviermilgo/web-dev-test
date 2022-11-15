package processors

import "sync"

type DriverRanking struct {
	// Fill in your properties here
}

func (*DriverRanking) String() string {
	// Implement this function
	return ""
}

type HotelRanking struct {
	// Fill in your properties here
}

func (*HotelRanking) String() string {
	// Implement this function
	return ""
}

type ProcessorInterface interface {
	StartProcessing() error
	GetTopRankedDriver() *DriverRanking
	GetTopRankedHotel() *HotelRanking
}

func CreateProcessorFromData(data *TripsData, wg *sync.WaitGroup) ProcessorInterface  {
	// @todo Initialize your processor here
	return nil
}
