create table refresh_token (
    uuid char(255) primary key not null unique,
    profile_uuid uuid not null references profile(uuid),

    expires_at timestamp not null,
    deleted_at timestamp,
    created_at timestamp not null
);
