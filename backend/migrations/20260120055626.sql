-- Modify "friendships" table
ALTER TABLE `friendships` MODIFY COLUMN `friendship_uuid` char (36) NULL, DROP PRIMARY KEY, ADD PRIMARY KEY (`user_uuid`, `friend_uuid`);
