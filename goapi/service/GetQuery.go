package service

import (
	"goapi/model"
)

func GetQuery(text string, subject string, sort string) []model.Kaouisan {
	db := GormConnect()
	defer db.Close()

	var res []model.Kaouisan

	if subject == "person" {
		db.Where("persons_name like ?", "%"+text+"%").Order("calendar").Find(&res)
	} else if subject == "document" {
		db.Where("document_name like ?", "%"+text+"%").Order("calendar").Find(&res)
	} else if subject == "era" {
		db.Where("era_name like ?", "%"+text+"%").Order("calendar").Find(&res)
	}

	return res
}
