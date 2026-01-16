-- Create "auths" table
CREATE TABLE `auths` (
  `uuid` char(36) NOT NULL,
  `jti` char(36) NULL,
  PRIMARY KEY (`uuid`),
  UNIQUE INDEX `uni_auths_jti` (`jti`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "friend_codes" table
CREATE TABLE `friend_codes` (
  `user_uuid` char(36) NOT NULL,
  `code` char(36) NOT NULL,
  PRIMARY KEY (`user_uuid`),
  UNIQUE INDEX `uni_friend_codes_code` (`code`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "friend_requests" table
CREATE TABLE `friend_requests` (
  `request_uuid` char(36) NOT NULL,
  `user_uuid` char(36) NOT NULL,
  `friend_uuid` char(36) NOT NULL,
  `status` varchar(20) NOT NULL,
  `request_at` datetime(3) NOT NULL,
  PRIMARY KEY (`request_uuid`),
  UNIQUE INDEX `idx_friend` (`user_uuid`, `friend_uuid`),
  INDEX `idx_friend_requests_request_uuid` (`request_uuid`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "friendships" table
CREATE TABLE `friendships` (
  `user_uuid1` char(36) NOT NULL,
  `user_uuid2` char(36) NOT NULL,
  `friend_uuid` char(36) NOT NULL,
  `accepted_at` datetime(3) NOT NULL,
  PRIMARY KEY (`user_uuid1`, `user_uuid2`),
  INDEX `idx_friendships_friend_uuid` (`friend_uuid`),
  UNIQUE INDEX `idx_friendships_user_uuid1` (`user_uuid1`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (
  `uuid` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `mail` varchar(191) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`uuid`),
  UNIQUE INDEX `uni_users_mail` (`mail`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
