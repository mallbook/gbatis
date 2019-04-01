use MALL

CREATE TABLE IF NOT EXISTS t_mall (
    id char(36) NOT NULL,
    name varchar(64) NOT NULL,
    avatar varchar(256) NOT NULL,
    createdAt int NOT NULL,
    updatedAt int NOT NULL,
    story varchar(1024),
    PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS t_shop (
    id char(36) NOT NULL,
    brandId char(36) NOT NULL,
    createdAt int NOT NULL,
    updatedAt int NOT NULL,
    PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS t_brand (
    id char(36) NOT NULL,
    name varchar(64) NOT NULL,
    avatar varchar(256) NOT NULL,
    createdAt int NOT NULL,
    updatedAt int NOT NULL,
    story varchar(1024),
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
