package data

import (
	"errors"
	"fmt"
	"log"
	"wilayah/domain"

	"gorm.io/gorm"
)

type memData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.MemberData {
	return &memData{
		db: DB,
	}
}

func (md *memData) InsertMember(newData domain.Member) (idMember int, err error) {
	cnv := FromModel(newData)
	result := md.db.Table("members").Create(&cnv)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error from db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (md *memData) GetAll() ([]domain.Member, error) {
	var memberData []Member
	err := md.db.Find(&memberData).Error
	if err != nil {
		return []domain.Member{}, err
	}
	var convert []domain.Member
	for i := 0; i < len(memberData); i++ {
		convert = append(convert, memberData[i].ToModel())
	}
	return convert, nil
}

func (md *memData) Update(id int, updatedData domain.Member) (idMember int, err error) {
	var cnv = FromModel(updatedData)
	cnv.ID = uint(id)
	result := md.db.Model(&Member{}).Where("ID = ?", id).Updates(cnv)
	if result.Error != nil {
		log.Println("Cannot update data", errors.New("error db"))
		return -1, errors.New("data already exist")
	}
	if result.RowsAffected == 0 {
		return -1, errors.New("failed update data")
	}

	return int(result.RowsAffected), nil
}

func (md *memData) Delete(id int) (idMember int, err error) {
	res := md.db.Delete(&Member{}, id)
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
