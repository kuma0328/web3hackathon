INSERT INTO recipes (`name`, `img_url`, `community_id`) VALUES ("ramen", "https://hoge.png", "1");

INSERT INTO recipe_steps (`number`, `content`, `recipe_id`) VALUES ("1", "step1 : hoge", "1");
INSERT INTO recipe_steps (`number`, `content`, `recipe_id`) VALUES ("2", "step2 : fuga", "1");
INSERT INTO recipe_steps (`number`, `content`, `recipe_id`) VALUES ("3", "step13: piyo", "1");

INSERT INTO spices (`name`, `recipe_id`) VALUES ("sato", "1");
INSERT INTO spices (`name`, `recipe_id`) VALUES ("sio", "1");
INSERT INTO spices (`name`, `recipe_id`) VALUES ("su", "1");
INSERT INTO spices (`name`, `recipe_id`) VALUES ("shoyu", "1");
INSERT INTO spices (`name`, `recipe_id`) VALUES ("miso", "1");
