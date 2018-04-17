CREATE TABLE `pachong`.`pc_sub_user` (
  `pc_sub_user_id` INT NOT NULL AUTO_INCREMENT,
  `pc_sub_user_sub` LONGTEXT NULL,
  `pc_sub_user_ids` LONGTEXT NULL,
  PRIMARY KEY (`pc_sub_user_id`))
  COMMENT = '订阅指向用户';
