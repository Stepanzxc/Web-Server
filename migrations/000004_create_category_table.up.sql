create table if not exists category
(
    category_id  int          not null auto_increment,
	title       VARCHAR(255) not null,
    primary key (`category_id`)

);