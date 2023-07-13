create table if not exists provider
(
    provider_id int not null auto_increment,
    title varchar(255) not null,
    created_at timestamp default now(),
    status bool,
    primary key (`provider_id`)
);
