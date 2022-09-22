package data

import (
	"errors"
	"fmt"
	"log"
	"wilayah/domain"

	"gorm.io/gorm"
)

type kelData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.KelurahanData {
	return &kelData{
		db: DB,
	}
}

func (kd *kelData) InsertKelurahan(newData domain.Kelurahan) (idKelurahan int, err error) {
	cnv := FromModel(newData)
	result := kd.db.Table("kelurahans").Create(&cnv)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (kd *kelData) GetAll() ([]domain.Kelurahan, error) {
	var kelurahanData []Kelurahan
	err := kd.db.Find(&kelurahanData).Error
	if err != nil {
		return []domain.Kelurahan{}, err
	}
	var convert []domain.Kelurahan
	for i := 0; i < len(kelurahanData); i++ {
		convert = append(convert, kelurahanData[i].ToModel())
	}
	return convert, nil
}

func (kd *kelData) Update(id int, updatedData domain.Kelurahan) (idKelurahan int, err error) {
	var cnv = FromModel(updatedData)
	cnv.ID = uint(id)
	result := kd.db.Model(&Kelurahan{}).Where("ID = ?", id).Updates(cnv)
	if result.Error != nil {
		log.Println("Cannot update data", errors.New("error db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return -1, errors.New("failed update data")
	}

	return int(result.RowsAffected), nil
}

func (kd *kelData) Delete(id int) (idKelurahan int, err error) {
	res := kd.db.Delete(&Kelurahan{}, id)
	if res.Error != nil {
		log.Println("cannot delete data", res.Error.Error())
		return 0, res.Error
	}
	if res.RowsAffected < 1 {
		log.Println("no data deleted", res.Error.Error())
		return 0, fmt.Errorf("failed to delete data")
	}
	return int(res.RowsAffected), nil
}
