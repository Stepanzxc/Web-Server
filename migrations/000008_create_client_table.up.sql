create table if not exists client
(
    client_id  int          not null auto_increment,
	address       VARCHAR(255) not null,
    primary key (`client_id`)

);