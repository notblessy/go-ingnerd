-- migrate:up
create table todos (
  id integer,
  name varchar(255) not null,
  description varchar(255)
);

-- migrate:down
drop table todos;
