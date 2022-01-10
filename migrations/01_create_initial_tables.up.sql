DROP TABLE IF EXISTS tasks CASCADE;

CREATE TABLE tasks
(
    id              INT              NOT NULL AUTO_INCREMENT,
    title           VARCHAR(100)     NOT NULL,
    description     TEXT             NOT NULL,
    completed       TINYINT(1)       DEFAULT 0,
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP       NULL,
    deleted_at      TIMESTAMP       NULL,
    PRIMARY KEY (id)
);
