-- CreateTable
CREATE TABLE `isan` (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `control_number` VARCHAR(191) NOT NULL,
    `era_name` VARCHAR(191) NOT NULL,
    `japanese_calendar` VARCHAR(191) NOT NULL,
    `western_calendar` VARCHAR(191) NOT NULL,
    `calendar` DATETIME(3) NOT NULL,
    `document_name` VARCHAR(191) NOT NULL,
    `persons_name` VARCHAR(191) NOT NULL,
    `historical_materials` VARCHAR(191) NOT NULL,
    `position_name` VARCHAR(191) NOT NULL,
    `pic_no` VARCHAR(191) NOT NULL,
    `pic_path` VARCHAR(191) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
