package handler

import (
	"goapi/model"
	"goapi/service"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetIiif(c *gin.Context) {
	// var res model.Manifest
	text := c.Query("text")
	subject := c.DefaultQuery("subject", "person")
	sort := c.DefaultQuery("sort", "ASC")
	data := service.GetQuery(text, subject, sort)

	// baseUrl := "http://localhost:4455"
	baseUrl := os.Getenv("BASE_URL")

	textEncoded := url.QueryEscape(text)

	ReratedIn := []model.RelatedIn{
		{
			ID:     baseUrl + "/s/" + textEncoded + "/manifest",
			Format: "text/html",
		},
	}

	MetadataIn := []model.MetadataIn{
		{
			Label: "Search term",
			Value: text,
		},
		{
			Label: "DCTERMS.alternative",
			Value: "かおういさん",
		},
	}

	var canvaseList []model.CanvasesIn

	for _, i := range data {
		var wi int
		wi, _ = strconv.Atoi(i.Width)
		var he int
		he, _ = strconv.Atoi(i.Height)

		picPath := strings.Replace(i.Pic_no, ".bmp", "", -1)

		canvase := model.CanvasesIn{
			Type:   "sc:Canvas",
			Label:  i.Persons_name + "_" + i.Historical_materials,
			Width:  wi,
			Height: he,
			ID:     baseUrl + "/i/" + textEncoded + "/canvas/" + i.Id,
			Images: []model.CanvasesImagesIn{
				{
					ID:         baseUrl + "/d/" + picPath,
					Motivation: "sc:painting",
					Resource: model.CanvasesImagesResourceIn{
						ID:     baseUrl + "/d/" + picPath,
						Format: "image/jpg",
						Type:   "oa:Annotation",
						Height: he,
						Service: model.CanvasesImagesResourceServiceIn{
							ID:      baseUrl + "/i/" + textEncoded + "/200017283_00002.tif",
							Context: "http://iiif.io/api/image/2/context.json",
							Profile: "http://iiif.io/api/image/2/level2.json",
						},
						Width: wi,
					},
					Type: "oa:Annotation",
					On:   baseUrl + "/d/" + picPath,
				},
			},
		}
		canvaseList = append(canvaseList, canvase)

	}
	SequencesIn := []model.SequencesIn{
		{
			Canvases: canvaseList,
			ID:       baseUrl + "/i/" + textEncoded + "/sequence/all",
			Type:     "sc:Sequence",
			Label:    text + "の検索結果",
		},
	}

	res := model.Manifest{
		Related:     ReratedIn,
		Attribution: "<span>花押彙纂画像データ</span>",
		Within:      baseUrl + "/s/",
		Context:     "http://iiif.io/api/presentation/2/context.json",
		ID:          baseUrl + "/s/",
		License:     "http://creativecommons.org/licenses/by-sa/4.0/",
		Thumbnail: model.ThumbnailIn{
			ID: baseUrl + "/i/" + textEncoded + "/200017283_00001.tif",
		},
		Description:      "",
		Logo:             baseUrl + "/o/logo.png",
		Label:            "花押彙纂",
		Sequences:        SequencesIn,
		Type:             "sc:Manifest",
		Metadata:         MetadataIn,
		ViewingDirection: "right-to-left",
	}
	c.JSON(200, res)

}
