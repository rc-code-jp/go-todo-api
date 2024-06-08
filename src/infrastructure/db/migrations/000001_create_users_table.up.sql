CREATE TABLE `users` (
  `id`              INT(11)      NOT NULL AUTO_INCREMENT,
  `name`            VARCHAR(255) NOT NULL,
  `email`           VARCHAR(255) NOT NULL,
  `hashed_password` VARCHAR(255) NOT NULL,
  `image_file_path` VARCHAR(255) NULL,
  `deleted_at`      DATETIME(3)  NULL,
  `created_at`      DATETIME(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at`      DATETIME(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
)
