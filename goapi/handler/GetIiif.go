package handler

import (
	"goapi/model"

	"github.com/gin-gonic/gin"
)

func GetIiif(c *gin.Context) {
	// var res model.Manifest
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

	SequencesIn := []model.SequencesIn{
		{
			Canvases: []model.CanvasesIn{
				{
					Type:   "sc:Canvas",
					Label:  "Page 1",
					Width:  5760,
					Height: 3840,
					ID:     "http://codh.rois.ac.jp/pmjt/iiif/200017283/canvas/00001",
					Images: []model.CanvasesImagesIn{
						{
							ID:         "http://codh.rois.ac.jp/pmjt/iiif/200017283/annotation/00001",
							Motivation: "sc:painting",
							Resource: model.CanvasesImagesResourceIn{
								ID:     "http://codh.rois.ac.jp/pmjt/iiif/200017283/200017283_00001.tif/full/full/0/default.jpg",
								Format: "image/jpeg",
								Type:   "dctypes:Image",
								Height: 3840,
								Service: model.CanvasesImagesResourceServiceIn{
									ID:      "http://codh.rois.ac.jp/pmjt/iiif/200017283/200017283_00001.tif",
									Context: "http://iiif.io/api/image/2/context.json",
									Profile: "http://iiif.io/api/image/2/level2.json",
								},
								Width: 5760,
							},
							Type: "oa:Annotation",
							On:   "http://codh.rois.ac.jp/pmjt/iiif/200017283/canvas/00001",
						},
					},
				},
			},
			ID:    "http://codh.rois.ac.jp/pmjt/iiif/200017283/sequence/all",
			Type:  "sc:Sequence",
			Label: "ページ一覧",
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
