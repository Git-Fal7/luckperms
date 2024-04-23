CREATE TABLE luckperms_user_permissions (
    id SERIAL PRIMARY KEY,
    uuid varchar(36) NOT NULL,
    permission varchar(200) NOT NULL,
    value boolean NOT NULL,
    server varchar(36) NOT NULL,
    world varchar(64) NOT NULL,
    expiry bigint NOT NULL,
    contexts varchar(200) NOT NULL
);

CREATE TABLE luckperms_groups (
  name varchar(36) PRIMARY KEY NOT NULL
);