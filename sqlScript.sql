CREATE DATABASE shortlink

USE shortlink;

CREATE TABLE link (
    LinkId int PRIMARY KEY AUTO_INCREMENT,
    FullLink text NOT NULL,
    ShortLink text NOT NULL
)