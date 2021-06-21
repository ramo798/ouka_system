use my_database;

CREATE TABLE kaouisan (
    id INT AUTO_INCREMENT NOT NULL,
    control_number VARCHAR(50),
    era_name VARCHAR(10),
    japanese_calendar VARCHAR(50),
    western_calendar VARCHAR(50),
    calendar datetime,
    document_name VARCHAR(50),
    persons_name VARCHAR(50),
    historical_materials VARCHAR(50),
    position_name VARCHAR(50),
    pic_no VARCHAR(50),
    pic_path VARCHAR(255),
    PRIMARY KEY (id)
);


-- INSERT INTO kaouisan (control_number,era_name,japanese_calendar,western_calendar,document_name,persons_name,historical_materials,position_name,pic_no,pic_path) VALUE ('00000002','建武','建武３年３月２０日',13360030200,'藤原数門軍忠状証判','上野頼兼','武雄神社文書','承了（花押）','00000002.bmp','kaou\kaouis\kaoi1\0603082\00000002.bmp');
-- INSERT INTO kaouisan (control_number,era_name,japanese_calendar,western_calendar,document_name,persons_name,historical_materials,position_name,pic_no,pic_path) VALUE ('00000003','建武２','建武３23年３月２０日',13360030202,'藤原数門軍忠状証判1111','上野頼兼33','2121武雄神社文書','承了（2121212花押）','00000022.bmp','kaou\kaouis\kaoi1\0603082\00000002.bmp');
-- INSERT INTO kaouisan (control_number,era_name,japanese_calendar,western_calendar,document_name,persons_name,historical_materials,position_name,pic_no,pic_path) VALUE ('00000004','建武３','建武３2323年３月２０日',13360030202,'藤原数門軍忠状証判2222','上野22頼兼','武雄1212神社文書','承212了（花押）','00000012.bmp','kaou\kaouis\kaoi1\0603082\00000002.bmp');