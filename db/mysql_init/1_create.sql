use my_database;

CREATE TABLE kaouisan (
    id INT(10) AUTO_INCREMENT NOT NULL,
    control_number int(8),
    era_name VARCHAR(10),
    japanese_calendar VARCHAR(16),
    western_calendar int(12),
    calendar datetime,
    document_name VARCHAR(50),
    persons_name VARCHAR(50),
    historical_materials VARCHAR(50),
    position_name VARCHAR(50),
    pic_no VARCHAR(50),
    pic_path VARCHAR(255),
    PRIMARY KEY (id)
);