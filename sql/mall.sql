use MALL

CREATE TABLE IF NOT EXISTS t_mall (
    id varchar(64),
    name varchar(32),
    avatar varchar(256),
    createdAt int,
    updatedAt int,
    story varchar(1024),
    PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS t_shop (
    id varchar(64),
    brandId varchar(64),
    createdAt int,
    updatedAt int,
    PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS t_brand (
    id varchar(64),
    name varchar(32),
    avatar varchar(256),
    createdAt int,
    updatedAt int,
    story varchar(1024),
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
