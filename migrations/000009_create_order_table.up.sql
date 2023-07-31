create table if not exists `order`
(
    order_id  int          not null auto_increment,
    client_id int          not null,
	price       int          not null,
    created_at DATE not null,
    constraint fk_client_id foreign key (client_id)
        references client (client_id)
        on delete cascade
        on update cascade,
    primary key (`order_id`)

);