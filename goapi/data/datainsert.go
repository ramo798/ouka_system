package data

import (
	"encoding/csv"
	"goapi/model"
	"goapi/service"
	"io"
	"log"
	"os"
	"time"
)

func Datatuika() {
	f, err := os.Open("足利花押彙纂画像データ.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	db := service.GormConnect()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(record[7])
		// fmt.Println(record[10])
		// fmt.Println(record[11])
		var tmp model.Kaouisan
		// tmp.height = 100
		db.Table("kaouisan").Where("Pic_no = ?", record[7]).First(&tmp)
		// fmt.Println(tmp)
		tmp.Height = record[11]
		tmp.Width = record[10]
		tmp.Calendar = time.Now()
		db.Table("kaouisan").Save(&tmp)
	}
}
