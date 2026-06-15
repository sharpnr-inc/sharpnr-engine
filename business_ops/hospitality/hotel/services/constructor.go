package hotelservices

import (
	hotelrepo "sharpnr-engine/business_ops/hospitality/hotel/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
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
