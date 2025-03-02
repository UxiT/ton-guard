create table profile (
                         uuid uuid primary key not null,
                         telegram_id int not null unique,
                         email varchar(255) not null,
                         password_hash varchar(255),

                         created_at timestamp with time zone not null,
                         updated_at timestamp with time zone not null
);

create table account (
                         uuid uuid primary key not null,
                         external_uuid uuid not null,
                         currency varchar(32) not null,
                         status varchar(64) not null,
                         balance numeric(25, 18) not null,

                         created_at timestamp with time zone not null,
                         updated_at timestamp with time zone not null
);
