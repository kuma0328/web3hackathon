CREATE TABLE `users` (
  `id`        int          COLLATE utf8mb4_bin NOT NULL AUTO_INCREMENT,
  `name`      varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `password`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `mail`      varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (id),
  unique (mail)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;