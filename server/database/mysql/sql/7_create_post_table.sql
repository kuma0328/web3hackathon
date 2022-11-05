CREATE TABLE `posts` (
  `id`            int          COLLATE utf8mb4_bin NOT NULL AUTO_INCREMENT,
  `content`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `img`           LONGBLOB     NOT NULL,
  `community_id`  int          COLLATE utf8mb4_bin NOT NULL,
  `user_id`       int          COLLATE utf8mb4_bin NOT NULL,
  -- `like`          int          COLLATE utf8mb4_bin NOT NULL AUTO_INCREMENT,
  FOREIGN KEY (community_id) REFERENCES communities(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;