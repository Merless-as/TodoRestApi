CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(256) not null,
    username      varchar(256) not null,
    password_hash varchar(256) not null
);

CREATE TABLE todo_list
(
    id          serial       not null unique,
    title       varchar(256) not null,
    description varchar(256)
);

CREATE TABLE users_lists
(
    id      serial                                          not null unique,
    user_id int references users (id) on delete cascade     not null,
    list_id int references todo_list (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(256) not null,
    description varchar(256),
    done        boolean      not null default false
);

CREATE TABLE lists_items
(
    id      serial                                           not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_list (id) on delete cascade  not null
);