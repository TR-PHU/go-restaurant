create table restaurants
(
    id                  int auto_increment
        primary key,
    owner_id            int                                 null,
    name                varchar(255)                        not null,
    address             varchar(255)                        not null,
    city_id             int                                 null,
    lat                 double                              null,
    lng                 double                              null,
    cover               json                                null,
    logo                json                                null,
    shipping_fee_per_km double    default 0                 null,
    status              int       default 1                 not null,
    created_at          timestamp default CURRENT_TIMESTAMP null,
    updated_at          timestamp default CURRENT_TIMESTAMP null
);

create index city_id
    on restaurants (city_id);

create index owner_id
    on restaurants (owner_id);

create index status
    on restaurants (status);