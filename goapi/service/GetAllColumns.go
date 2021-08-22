package service

import (
	"fmt"
	"goapi/model"
)

func GetAllColumns() []model.Kaouisan {
	db := gormConnect()
	defer db.Close()

	var all []model.Kaouisan

	db.Find(&all).Limit(30)

	for i := 0; i < len(all); i++ {
		fmt.Println(all[i])
	}

	return all
}
