CREATE TABLE `tasks` (
  `id`             INT(11)        NOT NULL AUTO_INCREMENT,
  `user_id`        INT(11)        NOT NULL,
  `task_group_id`  INT(11)        NOT NULL,
  `parent_task_id` INT(11)        NULL,
  `title`          VARCHAR(255)   NOT NULL,
  `date`           DATE           NULL,
  `time`           TIME           NULL,
  `note`           TEXT           NULL,
  `completed_at`   DATETIME(3)    NULL,
  `deleted_at`     DATETIME(3)    NULL,
  `created_at`     DATETIME(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at`     DATETIME(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`)
)
