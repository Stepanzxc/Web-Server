create table if not exists order_information
(
    product_id  int          not null,
    order_id int          not null,
	quantity       int          not null,
    constraint fk_product_id foreign key (product_id)
        references product (product_id),
    constraint fk_order_id foreign key (order_id)
        references `order` (order_id),
    primary key (product_id,order_id)

);
