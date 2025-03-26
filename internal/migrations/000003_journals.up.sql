
create table batch (
                       uuid uuid primary key not null,
                       first_id uuid not null,
                       status varchar(32) not null,
                       type varchar(32) not null,
                       description varchar,
                       meta jsonb,

                       created_at timestamp not null
);

create table batch_final (
                             uuid uuid primary key not null,
                             status varchar(32) not null,
                             type varchar(32) not null,
                             description varchar,
                             meta jsonb,

                             created_at timestamp not null,
                             updated_at timestamp not null
);

create table journal_entry (
                               uuid uuid primary key not null,
                               type varchar(32) not null,
                               amount numeric(25,18) not null,
                               description varchar,
                               meta jsonb,
                               account_id uuid references account(uuid),
                               batch_id uuid references batch(uuid),

                               created_at timestamp not null
);