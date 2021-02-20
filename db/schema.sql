-- This is required for generating UUID in PostgreSQL
create extension pgcrypto;

create table users (
    apiKey char(66) primary key,
    address char(42) not null,
    ts timestamp not null,
    enabled boolean default true
);

create index on users(address);

create table tasks (
    id uuid default gen_random_uuid() primary key,
    client char(66) not null,
    startBlock bigint not null,
    contract char(66),
    topic0 char(66),
    topic1 char(66),
    topic2 char(66),
    topic3 char(66),
    ts timestamp not null,
    foreign key (client) references users(apiKey)
);

create index on tasks(client);
create index on tasks(contract);
create index on tasks(topic0);
create index on tasks(topic1);
create index on tasks(topic2);
create index on tasks(topic3);

create table event_logs (
    origin char(42) not null,
    index integer not null,
    topics text[] not null,
    data bytea,
    txHash char(66) not null,
    blockHash char(66) not null,
    blockNumber bigint not null,
    primary key (blockHash, index),
);

create index on event_logs(origin);
create index on event_logs(txHash);
create index on event_logs(blockNumber);
create index on event_logs using gin(topics);
