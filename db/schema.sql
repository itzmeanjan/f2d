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
    client char(42) not null,
    startBlock bigint not null,
    contract char(66),
    topic0 char(66),
    topic1 char(66),
    topic2 char(66),
    topic3 char(66),
    ts timestamp not null
);
