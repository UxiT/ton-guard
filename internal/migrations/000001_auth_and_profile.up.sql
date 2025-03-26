create table profile (
                         uuid uuid primary key not null,
                         telegram_id int not null unique,
                         email varchar(255) not null,
                         password_hash varchar(255),

                         created_at timestamp not null,
                         updated_at timestamp not null
);

create table customer (
                          uuid uuid primary key not null unique,
                          profile_uuid uuid not null references profile(uuid),

                          created_at timestamp not null,
                          updated_at timestamp not null
);

create table account (
    uuid uuid primary key not null,
    external_uuid uuid not null,

    currency varchar(32) not null,
    status varchar(64) not null,
    balance numeric(25, 18) not null,

    customer_uuid uuid not null references customer(uuid),

    created_at timestamp not null,
    updated_at timestamp not null
);
