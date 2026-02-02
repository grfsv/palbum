CREATE TABLE `dtks` (
  `user_uuid` char(36) NOT NULL,
  `date` date NOT NULL,
  `key` binary(16) NOT NULL,
  PRIMARY KEY (`user_uuid`, `date`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
