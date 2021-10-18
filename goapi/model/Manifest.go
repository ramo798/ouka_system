package model

type RelatedIn struct {
	ID     string `json:"@id"`
	Format string `json:"format"`
}

type ThumbnailIn struct {
	ID string `json:"@id"`
}

type MetadataIn struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type SequencesIn struct {
	Canvases []CanvasesIn `json:"canvases"`
	ID       string       `json:"@id"`
	Type     string       `json:"@type"`
	Label    string       `json:"label"`
}

type CanvasesIn struct {
	Type   string             `json:"@type"`
	Label  string             `json:"label"`
	Width  int                `json:"width"`
	Height int                `json:"height"`
	ID     string             `json:"@id"`
	Images []CanvasesImagesIn `json:"images"`
}

type CanvasesImagesIn struct {
	ID         string                   `json:"@id"`
	Motivation string                   `json:"motivation"`
	Resource   CanvasesImagesResourceIn `json:"resource"`
	Type       string                   `json:"@type"`
	On         string                   `json:"on"`
}

type CanvasesImagesResourceIn struct {
	ID      string                          `json:"@id"`
	Format  string                          `json:"format"`
	Type    string                          `json:"@type"`
	Height  int                             `json:"height"`
	Service CanvasesImagesResourceServiceIn `json:"service"`
	Width   int                             `json:"width"`
}

type CanvasesImagesResourceServiceIn struct {
	ID      string `json:"@id"`
	Context string `json:"@context"`
	Profile string `json:"profile"`
}

type Manifest struct {
	Related          []RelatedIn   `json:"related"`
	Attribution      string        `json:"attribution"`
	Within           string        `json:"within"`
	Context          string        `json:"@context"`
	ID               string        `json:"@id"`
	License          string        `json:"license"`
	Thumbnail        ThumbnailIn   `json:"thumbnail"`
	Description      string        `json:"description"`
	Logo             string        `json:"logo"`
	Label            string        `json:"label"`
	Sequences        []SequencesIn `json:"sequences"`
	Type             string        `json:"@type"`
	Metadata         []MetadataIn  `json:"metadata"`
	ViewingDirection string        `json:"viewingDirection"`
}
