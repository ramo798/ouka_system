package model

import (
	"time"
)

type Kaouisan struct {
	Id                   string
	Control_number       string
	Era_name             string
	Japanese_calendar    string
	Western_calendar     string
	Calendar             time.Time
	Document_name        string
	Persons_name         string
	Historical_materials string
	Position_name        string
	Pic_no               string
	Pic_path             string
	Width                string
	Height               string
}

// CREATE TABLE kaouisan (
//     id INT AUTO_INCREMENT NOT NULL,
//     control_number VARCHAR(50),
//     era_name VARCHAR(10),
//     japanese_calendar VARCHAR(50),
//     western_calendar VARCHAR(50),
//     calendar datetime,
//     document_name VARCHAR(50),
//     persons_name VARCHAR(50),
//     historical_materials VARCHAR(50),
//     position_name VARCHAR(50),
//     pic_no VARCHAR(50),
//     pic_path VARCHAR(255),
//     PRIMARY KEY (id)
// );
