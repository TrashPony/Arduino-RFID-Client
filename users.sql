CREATE TABLE users (
  id       SERIAL PRIMARY KEY,
  uuid     VARCHAR(64),
  name     VARCHAR(64)
);

CREATE TABLE log (
  id       SERIAL PRIMARY KEY,
  uuid     VARCHAR(64),
  name     VARCHAR(64),
  event    VARCHAR(64),
  time     timestamp
)