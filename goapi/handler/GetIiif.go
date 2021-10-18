package handler

import (
	"goapi/model"
	"goapi/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIiif(c *gin.Context) {
	// var res model.Manifest
	text := c.Query("text")
	subject := c.DefaultQuery("subject", "person")
	sort := c.DefaultQuery("sort", "ASC")
	data := service.GetQuery(text, subject, sort)

	ReratedIn := []model.RelatedIn{
		{
			ID:     "http://codh.rois.ac.jp/pmjt/book/200017283/",
			Format: "text/html",
		},
	}

	MetadataIn := []model.MetadataIn{
		{
			Label: "DC.subject",
			Value: "",
		},
		{
			Label: "DC.identifier",
			Value: "200017283",
		},
		{
			Label: "巻之一～七",
			Value: "volume",
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

		canvase := model.CanvasesIn{
			Type:   "sc:Canvas",
			Label:  i.Persons_name + "_" + i.Historical_materials,
			Width:  wi,
			Height: he,
			ID:     "http://codh.rois.ac.jp/pmjt/iiif/200017283/canvas/00001",
			Images: []model.CanvasesImagesIn{
				{
					ID:         "http://hoge.ac.jp/" + i.Pic_path,
					Motivation: "sc:painting",
					Resource: model.CanvasesImagesResourceIn{
						ID:     "http://hoge.ac.jp/" + i.Pic_path,
						Format: "image/bmp",
						Type:   "oa:Annotation",
						Height: he,
						Service: model.CanvasesImagesResourceServiceIn{
							ID:      "http://codh.rois.ac.jp/pmjt/iiif/200017283/200017283_00002.tif",
							Context: "http://iiif.io/api/image/2/context.json",
							Profile: "http://iiif.io/api/image/2/level2.json",
						},
						Width: wi,
					},
					Type: "oa:Annotation",
					On:   "http://hoge.ac.jp/" + i.Pic_path,
				},
			},
		}
		canvaseList = append(canvaseList, canvase)

	}
	SequencesIn := []model.SequencesIn{
		{
			Canvases: canvaseList,
			ID:       "http://codh.rois.ac.jp/pmjt/iiif/200017283/sequence/all",
			Type:     "sc:Sequence",
			Label:    "ページ一覧",
		},
	}

	res := model.Manifest{
		Related:     ReratedIn,
		Attribution: "<span>花押彙纂画像データ</span>",
		Within:      "http://codh.rois.ac.jp/pmjt/book/",
		Context:     "http://iiif.io/api/presentation/2/context.json",
		ID:          "http://codh.rois.ac.jp/pmjt/book/200017283/manifest",
		License:     "http://creativecommons.org/licenses/by-sa/4.0/",
		Thumbnail: model.ThumbnailIn{
			ID: "http://codh.rois.ac.jp/pmjt/iiif/200017283/200017283_00001.tif",
		},
		Description:      "",
		Logo:             "http://codh.rois.ac.jp/img/codh_logo_tiny.png",
		Label:            "花押彙纂",
		Sequences:        SequencesIn,
		Type:             "sc:Manifest",
		Metadata:         MetadataIn,
		ViewingDirection: "right-to-left",
	}
	c.JSON(200, res)

}
