package data

import (
	"errors"
	"fmt"
	"log"
	"wilayah/domain"

	"gorm.io/gorm"
)

type kotaData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.KotaData {
	return &kotaData{
		db: DB,
	}
}

func (kd *kotaData) InsertKota(newData domain.Kota) (idKota int, err error) {
	cnv := FromModel(newData)
	result := kd.db.Table("kota").Create(&cnv)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (kd *kotaData) GetAll() ([]domain.Kota, error) {
	var kotaData []Kota
	err := kd.db.Find(&kotaData).Error
	if err != nil {
		return []domain.Kota{}, err
	}
	var convert []domain.Kota
	for i := 0; i < len(kotaData); i++ {
		convert = append(convert, kotaData[i].ToModel())
	}
	return convert, nil
}

func (pd *kotaData) Update(id int, updatedData domain.Kota) (idKota int, err error) {
	var cnv = FromModel(updatedData)
	cnv.ID = uint(id)
	result := pd.db.Model(&Kota{}).Where("ID = ?", id).Updates(cnv)
	if result.Error != nil {
		log.Println("Cannot update data", errors.New("error db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return -1, errors.New("failed update data")
	}

	return int(result.RowsAffected), nil
}

func (pd *kotaData) Delete(id int) (idKota int, err error) {
	res := pd.db.Delete(&Kota{}, id)
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
