CREATE TABLE segments
(
    id serial primary key,
    slug varchar(255) not null unique
);

CREATE TYPE segment_operation AS ENUM('add', 'delete');

CREATE TABLE users_segments
(
    id serial primary key,
    user_id int not null,
    foreign key(segment_id) references segments on delete cascade,
    operation segment_operation not null,
    updated_at date not null
);