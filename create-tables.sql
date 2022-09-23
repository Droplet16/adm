DROP TABLE IF EXISTS u_role;
CREATE TABLE u_role (
  id             serial       NOT NULL,
  name           VARCHAR(50)  NOT NULL,
  auth_item_name VARCHAR(50)  NOT NULL,
  CONSTRAINT u_role_pk PRIMARY KEY (id)
);

INSERT INTO u_role
  (name, auth_item_name)
VALUES
  ('Пользователь', 'user'),
  ('Администратор', 'admin');

DROP TABLE IF EXISTS u_user;
CREATE TABLE u_user (
  id         serial       NOT NULL,
  login      VARCHAR(128) NOT NULL,
  password   VARCHAR(128) NOT NULL,
  email      VARCHAR(128) NOT NULL,
  u_role_id  int          NOT NULL,
  CONSTRAINT u_user_pk PRIMARY KEY (id)
);

DROP TABLE IF EXISTS u_user_role;
CREATE TABLE u_user_role (
  u_user_id        int          NOT NULL,
  u_role_id        int          NOT NULL,
  CONSTRAINT u_user_role_pk PRIMARY KEY (u_user_id, u_role_id),
  CONSTRAINT u_user_role_u_user_fk FOREIGN KEY (u_user_id) REFERENCES public.u_user(id),
  CONSTRAINT u_user_role_u_role_fk FOREIGN KEY (u_role_id) REFERENCES public.u_role(id)
  );