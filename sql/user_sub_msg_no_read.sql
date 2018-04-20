CREATE TABLE `pachong`.`user_sub_msg_no_read` (
  `user_sub_msg_no_read_id`     INT         NOT NULL AUTO_INCREMENT,
  `user_sub_msg_no_read_userid` VARCHAR(45) NOT NULL,
  `user_sub_msg_no_read_msg`    LONGTEXT    NULL,
  PRIMARY KEY (`user_sub_msg_no_read_id`),
  UNIQUE INDEX `user_sub_msg_no_read_userid_UNIQUE` (`user_sub_msg_no_read_userid` ASC)
);

ALTER TABLE `pachong`.`user_sub_msg_no_read`
  ADD COLUMN `user_sub_msg_readed_msg` LONGTEXT NULL
  AFTER `user_sub_msg_no_read_msg`, RENAME TO `pachong`.`user_sub_msg_read`;

ALTER TABLE `pachong`.`user_sub_msg_read`
  CHANGE COLUMN `user_sub_msg_no_read_id` `user_sub_msg_read_id` INT(11) NOT NULL AUTO_INCREMENT,
  CHANGE COLUMN `user_sub_msg_no_read_userid` `user_sub_msg_read_userid` VARCHAR(45) NOT NULL;

ALTER TABLE `pachong`.`user_sub_msg_read`
  CHANGE COLUMN `user_sub_msg_no_read_msg` `user_sub_msg_no_read_msg` LONGTEXT NULL,
  CHANGE COLUMN `user_sub_msg_readed_msg` `user_sub_msg_readed_msg` LONGTEXT NULL;
