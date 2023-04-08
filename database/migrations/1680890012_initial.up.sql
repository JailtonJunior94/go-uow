CREATE TABLE categories (
    id varchar(36) NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    description varchar(500)
);

CREATE TABLE courses (
    id varchar(36) NOT NULL PRIMARY KEY,
    category_id varchar(36) NOT NULL,
    name varchar(255) NOT NULL,
    description varchar(500),
    price decimal(10,2) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);