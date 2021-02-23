-- This is required for generating UUID in PostgreSQL
create extension pgcrypto;

create table users (
    apikey char(66) primary key,
    address char(42) not null,
    ts timestamp not null,
    enabled boolean default true
);

create index on users(address);

create table tasks (
    id uuid default gen_random_uuid() primary key,
    client char(66) not null,
    startblock bigint not null,
    contract char(66),
    topic0 char(66),
    topic1 char(66),
    topic2 char(66),
    topic3 char(66),
    ts timestamp not null,
    enabled boolean default true,
    foreign key (client) references users(apikey)
);

create index on tasks(client);
create index on tasks(startblock);
create index on tasks(contract);
create index on tasks(topic0);
create index on tasks(topic1);
create index on tasks(topic2);
create index on tasks(topic3);
create index on tasks(enabled);

create table event_logs (
    origin char(42) not null,
    index integer not null,
    topics text[] not null,
    data bytea,
    txhash char(66) not null,
    blockhash char(66) not null,
    blocknumber bigint not null,
    primary key (blockhash, index)
);

create index on event_logs(origin);
create index on event_logs(txhash);
create index on event_logs(blocknumber);
create index on event_logs using gin(topics);

create table task_results (
    index integer not null,
    blockhash char(66) not null,
    id uuid not null,
    primary key(blockhash, index, id),
    foreign key (index) references event_logs(index) on delete cascade,
    foreign key (blockhash) references event_logs(blockhash) on delete cascade,
    foreign key (id) references tasks(id)
);
