DROP DATABASE IF EXISTS activity_app;

CREATE DATABASE activity_app;

USE activity_app;

CREATE TABLE activity (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    day varchar (255),
    time varchar (255),
    description varchar (255)
);
