CREATE TABLE `refresh_tokens` (
  `id`           INT(11)      NOT NULL AUTO_INCREMENT,
  `uuid`         VARCHAR(255) NOT NULL,
  `hashed_token` VARCHAR(255) NOT NULL,
  `user_id`      INT(11)      NOT NULL,
  `revoked`      TINYINT(1)   NOT NULL DEFAULT 0,
  `created_at`   DATETIME(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at`   DATETIME(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  UNIQUE KEY `refresh_token_uuid_unique` (`uuid`),
)
