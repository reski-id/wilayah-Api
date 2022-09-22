package data

import (
	"errors"
	"fmt"
	"log"
	"wilayah/domain"

	"gorm.io/gorm"
)

type provinsiData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.ProvinsiData {
	return &provinsiData{
		db: DB,
	}
}

func (pd *provinsiData) InsertProvinsi(newProvinsi domain.Provinsi) (idProvinsi int, err error) {
	cnv := FromModel(newProvinsi)
	result := pd.db.Table("provinsis").Create(&cnv)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (pd *provinsiData) GetAll() ([]domain.Provinsi, error) {
	var provinsiData []Provinsi
	err := pd.db.Find(&provinsiData).Error
	if err != nil {
		return []domain.Provinsi{}, err
	}
	var convert []domain.Provinsi
	for i := 0; i < len(provinsiData); i++ {
		convert = append(convert, provinsiData[i].ToModel())
	}
	return convert, nil
}

func (pd *provinsiData) Update(id int, updatedData domain.Provinsi) (idProvinsi int, err error) {
	var cnv = FromModel(updatedData)
	cnv.ID = uint(id)
	result := pd.db.Model(&Provinsi{}).Where("ID = ?", id).Updates(cnv)
	if result.Error != nil {
		log.Println("Cannot update data", errors.New("error db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return -1, errors.New("failed update data")
	}

	return int(result.RowsAffected), nil
}

func (pd *provinsiData) Delete(id int) (idProvinsi int, err error) {
	res := pd.db.Delete(&Provinsi{}, id)
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
