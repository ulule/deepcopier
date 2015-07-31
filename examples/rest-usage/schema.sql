CREATE TABLE account (
    id serial PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (355) UNIQUE NOT NULL,
    date_joined TIMESTAMP NOT NULL
);

insert into account (username, first_name, last_name, password, email, date_joined) values ('thoas', 'Florent', 'Messa', '8d56e93bcc8d63a171b5630282264341', 'foo@bar.com', '2015-07-31 15:10:10');
insert into account (username, first_name, last_name, password, email, date_joined) values ('gilles', 'Gilles', 'Fabio', '8d56e93bcc8d63a171b5630282264341', 'bar@foo.com', '2015-07-31 16:10:10');
