create table card (
    uuid uuid primary key not null unique,
    external_uuid uuid not null unique,
    account_uuid uuid references account(uuid),

    created_at timestamp not null,
    updated_at timestamp not null
)