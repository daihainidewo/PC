CREATE TABLE `pachong`.`user_sub` (
  `user_sub_id` INT NOT NULL AUTO_INCREMENT,
  `user_sub_user_id` VARCHAR(45) NULL,
  `user_sub_sub_msg` LONGTEXT NULL,
  PRIMARY KEY (`user_sub_id`));

ALTER TABLE `pachong`.`user_sub`
  CHANGE COLUMN `user_sub_user_id` `user_sub_user_id` VARCHAR(45) NOT NULL ,
  ADD UNIQUE INDEX `user_sub_user_id_UNIQUE` (`user_sub_user_id` ASC);
