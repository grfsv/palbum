-- Create "auth_entities" table
CREATE TABLE `auth_entities` (
  `uuid` char(36) NOT NULL,
  `jti` char(36) NULL,
  PRIMARY KEY (`uuid`),
  UNIQUE INDEX `uni_auth_entities_jti` (`jti`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "user_entities" table
CREATE TABLE `user_entities` (
  `uuid` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `mail` varchar(191) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`uuid`),
  UNIQUE INDEX `uni_user_entities_mail` (`mail`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
