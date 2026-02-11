package hotelservices

import (
	dbtx "sharpnr-engine/core/dbTx"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
)

type hotelService struct {
	hotelRepo hotelrepo.HotelRepositoryInterface
	txManager dbtx.TxManager
}

func NewHotelService(
	hotelRepo hotelrepo.HotelRepositoryInterface,
	txManager dbtx.TxManager,

) *hotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
		txManager: txManager,
	}
}
