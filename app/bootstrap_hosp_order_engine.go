package app

import (
	hospitality_order_engine_repository "sharpnr-engine/modules/hospitality/order_engine/repository"
	hospitality_order_engine_service "sharpnr-engine/modules/hospitality/order_engine/service"

	"gorm.io/gorm"
)

func bootstrapHospitalityOrderEngine(db *gorm.DB) hospitality_order_engine_service.HospitalityOrderEngineServiceInterface {
	repo := hospitality_order_engine_repository.NewHospitalityOrderEngine(db)
	return hospitality_order_engine_service.NewHospitalityOrderEngineService(repo)
}
