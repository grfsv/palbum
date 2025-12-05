-- Create "auths" table
CREATE TABLE `auths` (
  `user_uuid` char(36) NOT NULL,
  `jti` char(36) NULL,
  PRIMARY KEY (`user_uuid`),
  UNIQUE INDEX `uni_auths_jti` (`jti`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (
  `uuid` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `mail` varchar(191) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`uuid`),
  UNIQUE INDEX `idx_users_mail` (`mail`),
  UNIQUE INDEX `uni_users_mail` (`mail`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
