package service

import (
	"fmt"
	"goapi/model"
)

func GetAllColumns() []model.Kaouisan {
	db := GormConnect()
	defer db.Close()

	var all []model.Kaouisan

	db.Limit(30).Find(&all).Limit(30)

	for i := 0; i < len(all); i++ {
		fmt.Println(all[i])
	}

	return all
}
