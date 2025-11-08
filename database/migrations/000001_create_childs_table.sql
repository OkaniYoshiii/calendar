-- +goose up
CREATE TABLE childs (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL,
    `birthday` DATE NOT NULL,
    `nickname` VARCHAR(50)
);

-- +goose down
DROP TABLE childs;