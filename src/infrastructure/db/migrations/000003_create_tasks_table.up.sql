CREATE TABLE `tasks` (
  `id`          INT(11)        NOT NULL AUTO_INCREMENT,
  `user_id`     INT(11)        NOT NULL,
  `category_id` INT(11)        NULL,
  `amount`      DECIMAL(10, 2) NOT NULL,
  `date`        DATE           NOT NULL,
  `time`        TIME           NULL,
  `note`        TEXT           NULL,
  `place`       VARCHAR(255)   NULL,
  `deleted_at`  DATETIME(3)    NULL,
  `created_at`  DATETIME(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at`  DATETIME(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`)
)
