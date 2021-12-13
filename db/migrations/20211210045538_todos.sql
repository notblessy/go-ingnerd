-- migrate:up
create table todos (
  id integer primary key AUTO_INCREMENT not null,
  name varchar(255) not null,
  description varchar(255),
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

-- migrate:down
drop table todos;
