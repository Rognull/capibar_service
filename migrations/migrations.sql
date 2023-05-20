-- Active: 1675774118937@@127.0.0.1@5433
create table vote
(
    Id_voter INTEGER not null
            primary key,
    Id_candidate INTEGER,
    Id_election INTEGER
);



create unique index users_id_uindex
	on users (id);
