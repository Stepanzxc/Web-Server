create table if not exists product
(
    product_id  int          not null auto_increment,
    provider_id int          not null,
	title       VARCHAR(255) not null,
	description VARCHAR(255) not null,
	price       int          not null,
	brand       VARCHAR(255) not null,
	category    VARCHAR(255) not null,
    constraint fk_provider_id foreign key (provider_id)
        references provider (provider_id)
        on delete cascade
        on update cascade,
    primary key (`product_id`)

);
