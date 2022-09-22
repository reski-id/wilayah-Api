package data

import (
	"errors"
	"fmt"
	"log"
	"wilayah/domain"

	"gorm.io/gorm"
)

type kecData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.KecamatanData {
	return &kecData{
		db: DB,
	}
}

func (kd *kecData) InsertKecamatan(newData domain.Kecamatan) (idKecamatan int, err error) {
	cnv := FromModel(newData)
	result := kd.db.Table("kecamatan").Create(&cnv)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (kd *kecData) GetAll() ([]domain.Kecamatan, error) {
	var kecamatanData []Kecamatan
	err := kd.db.Find(&kecamatanData).Error
	if err != nil {
		return []domain.Kecamatan{}, err
	}
	var convert []domain.Kecamatan
	for i := 0; i < len(kecamatanData); i++ {
		convert = append(convert, kecamatanData[i].ToModel())
	}
	return convert, nil
}

func (kd *kecData) Update(id int, updatedData domain.Kecamatan) (idKecamatan int, err error) {
	var cnv = FromModel(updatedData)
	cnv.ID = uint(id)
	result := kd.db.Model(&Kecamatan{}).Where("ID = ?", id).Updates(cnv)
	if result.Error != nil {
		log.Println("Cannot update data", errors.New("error db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return -1, errors.New("failed update data")
	}

	return int(result.RowsAffected), nil
}

func (kd *kecData) Delete(id int) (idKecamatan int, err error) {
	res := kd.db.Delete(&Kecamatan{}, id)
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
