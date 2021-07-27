CREATE TABLE `artist` (
    `id` INT NOT NULL AUTO_INCREMENT ,
    `name` VARCHAR(200) NOT NULL,
    `album` VARCHAR(200) NOT NULL,
    `image_url` VARCHAR(200) DEFAULT NULL,
    `release_date` date NOT NULL,
    `price` NUMERIC(10,2) NOT NULL,
    `sample_url` VARCHAR(200) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `artist` (`name`, `album`, `image_url`, `release_date`, `price`, `sample_url`) VALUES
('ﾃｲﾗｰｽｳｨﾌﾄ', 'We Are Never Ever Getting Back\r\nTogether', NULL, '1980-01-22', 320.00, NULL);
INSERT INTO `artist` (`name`, `album`, `image_url`, `release_date`, `price`, `sample_url`) VALUES
('MONGOL800', '小さな恋のうた', NULL, '1982-09-02', 320.00, NULL);
