CREATE TABLE `pachong`.`pc_body_msg` (
  `pc_body_msg_id` INT NOT NULL AUTO_INCREMENT,
  `pc_body_msg_user_id` VARCHAR(45) NOT NULL,
  `pc_body_msg_body` LONGTEXT NULL,
  PRIMARY KEY (`pc_body_msg_id`));
