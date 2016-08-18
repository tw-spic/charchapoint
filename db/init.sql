CREATE USER charchapoint_user WITH PASSWORD '******';

CREATE DATABASE charchapoint;

\c charchapoint;

CREATE TABLE zones(
   id BIGSERIAL PRIMARY KEY,
   lat DECIMAL NOT NULL,
   long DECIMAL NOT NULL,
   radius DECIMAL NOT NULL
);

-- CREATE USER charchapoint_user WITH PASSWORD '******';

REVOKE CONNECT ON DATABASE charchapoint FROM PUBLIC;

GRANT CONNECT ON DATABASE charchapoint TO charchapoint_user;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO charchapoint_user;
GRANT SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA public TO charchapoint_user;