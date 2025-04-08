create table top_up (
    uuid char(255) primary key not null unique,
    profile_uuid uuid not null references profile(uuid),
    amount numeric(25, 18) not null,
    network varchar not null,
    status varchar not null,
    is_closed bool not null default false,

    deleted_at timestamp,
    created_at timestamp not null
);
