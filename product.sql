CREATE TABLE IF NOT EXISTS `product` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `price` INT(11) NOT NULL DEFAULT 0,
  `stock` INT(9) NOT NULL DEFAULT 0,
  `created_at` DATETIME DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  PRIMARY KEY(`id`)
)

INSERT INTO product(name,price,stock,created_at,updated_at) values("Kopi Janji jiwa", 28000, 12, NOW(), NOW())