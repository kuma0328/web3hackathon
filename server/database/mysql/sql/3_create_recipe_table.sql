CREATE TABLE `recipes` (
  `id`            int          COLLATE utf8mb4_bin NOT NULL AUTO_INCREMENT,
  `name`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `img_url`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `community_id`  int          COLLATE utf8mb4_bin NOT NULL,
  FOREIGN KEY (community_id) REFERENCES communities(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `recipe_steps` (
  `id`            int          COLLATE utf8mb4_bin NOT NULL AUTO_INCREMENT,
  `number`        int          COLLATE utf8mb4_bin NOT NULL,
  `content`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `recipe_id`     int          COLLATE utf8mb4_bin NOT NULL,
  FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `spices` (
  `id`            int          COLLATE utf8mb4_bin NOT NULL AUTO_INCREMENT,
  `name`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `recipe_id`     int          COLLATE utf8mb4_bin NOT NULL,
  FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
