CREATE DATABASE IF NOT EXISTS `run_log_db`;

use run_log_db

-- DROP TABLE IF EXISTS transactions;
-- CREATE TABLE transactions (
--     id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
--     user_id int,
--     amount int,
--     time DATETIME,
--     category int,
--     subcategory int,
--     upload_time DATETIME
-- );

-- TODO: hack, fix this later
SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64),
    username_slug VARCHAR(64),
    firstname VARCHAR(32),
    lastname VARCHAR(32)
);

DROP TABLE IF EXISTS activities;
CREATE TABLE activities (
    -- id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    activity_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    start_time INTEGER,
    total_time REAL,
    total_distance REAL,
    avg_heart_rate INTEGER,
    avg_cadence INTEGER,
    ascent REAL,
    descent REAL,
    records MEDIUMBLOB,
    laps BLOB,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (activity_id, user_id)
);


DELIMITER $$

CREATE TRIGGER activities_key_setter BEFORE INSERT ON activities
FOR EACH ROW BEGIN
    SET new.activity_id = (
        SELECT IFNULL(MAX(activity_id), 0) + 1
        FROM activities
        WHERE user_id = NEW.user_id
    );
END $$

DELIMITER ;

CREATE DATABASE IF NOT EXISTS `other`;

-- TODO: hack, fix this later
SET FOREIGN_KEY_CHECKS=1;