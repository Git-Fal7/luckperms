CREATE TABLE luckperms_user_permissions (
    id SERIAL PRIMARY KEY,
    uuid uuid NOT NULL,
    permission text NOT NULL,
    value int NOT NULL,
    server text NOT NULL,
    world text NOT NULL,
    expiry int NOT NULL,
    contexts text NOT NULL
);

CREATE TABLE luckperms_groups (
  name text PRIMARY KEY NOT NULL
);