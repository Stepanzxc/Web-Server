create table if not exists client
(
    client_id  int          not null auto_increment,
	name       VARCHAR(255) not null,
	surname VARCHAR(255) not null,
	address       VARCHAR(255) not null,
	phone_number    VARCHAR(255) not null,
    primary key (`client_id`)

);