CREATE TABLE `pachong`.`user_sub_msg` (
  `user_sub_msg_id` INT NOT NULL AUTO_INCREMENT,
  `user_sub_msg_user_id` VARCHAR(45) NOT NULL,
  `user_sub_msg_user_msg` LONGTEXT NULL,
  `user_sub_msg_user_sub` LONGTEXT NULL,
  `user_sub_msg_is_read` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`user_sub_msg_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COMMENT = '用户订阅信息的存取集合';
