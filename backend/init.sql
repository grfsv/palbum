-- MySQL Initialization Script for palbum Map

-- Ensure utf8mb4 character set and collation
ALTER DATABASE palbum_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Set timezone
SET time_zone = '+09:00';

-- Grant privileges to user
GRANT ALL PRIVILEGES ON palbum_db.* TO 'palbum_user'@'%';
FLUSH PRIVILEGES;

-- Note: マイグレーションツール（golang-migrate）を使用してテーブルを作成します
-- このファイルは初期設定のみを行います
