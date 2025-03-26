create table exchange_rates (
                                uuid uuid primary key not null,
                                iso_code varchar(32) not null,
                                rate numeric(25, 18) not null,
                                related_iso_code varchar(32) not null,
                                meta jsonb,

                                created_at timestamp not null
);
