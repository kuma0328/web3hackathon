CREATE TABLE `communities` (
  `id`      serial       COLLATE utf8mb4_bin NOT NULL,
  `name`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `img_url`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `content`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;