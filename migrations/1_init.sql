-- +migrate Up
create table "users" (
    "user_id" text primary key,
    "full_name" text,
    "email" text unique,
    "password" text,
    "role" text,
    "created_at" timestamptz not null,
    "updated_at" timestamptz not null
);

create table "repos" (
    "name" text primary key,
    "description" text,
    "url" text unique,
    "color" text,
    "lang" text,
    "fork" text,
    "stars" text,
    "stars_today" text,
    "build_by" text,
    "created_at" timestamptz not null,
    "updated_at" timestamptz not null
);

create table "bookmarks" (
    "bid" text primary key,
    "user_id" text unique,
    "repo_name" text,
    "created_at" timestamptz not null,
    "updated_at" timestamptz not null
);

alter table "bookmarks" add foreign key ("user_id") references "users" ("user_id");
alter table "bookmarks" add foreign key ("repo_name") references "repos" ("name");

-- +migrate Down
drop table bookmarks;
drop table users;
drop table repos;
