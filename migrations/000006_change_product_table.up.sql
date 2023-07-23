ALTER TABLE product
ADD category_id int NULL,
ADD CONSTRAINT fk_category_id
FOREIGN KEY(category_id) REFERENCES category(category_id)
        on delete cascade
        on update cascade;
