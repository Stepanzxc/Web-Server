create table if not exists product
(
    product_id  int          not null auto_increment,
    provider_id int          not null,

    #title-  - рпиши типы данных
    #description -  рпиши типы данных
    #price  - орпиши типы данных
    #brand  - орпиши типы данных
    #category  - орпиши типы данных

    constraint fk_provider_id foreign key (provider_id)
        references provider (provider_id)
        on delete cascade
        on update cascade,
    primary key (`product_id`)

);
