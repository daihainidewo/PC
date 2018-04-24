CREATE TABLE `pachong`.`user` (
  `usernum` INT NOT NULL AUTO_INCREMENT,
  `userid` VARCHAR(45) NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `userpasswd` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`usernum`));

ALTER TABLE `pachong`.`user`
ADD UNIQUE INDEX `userid_UNIQUE` (`userid` ASC),
ADD UNIQUE INDEX `username_UNIQUE` (`username` ASC);
