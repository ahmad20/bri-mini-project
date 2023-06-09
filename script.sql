DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS customers;

CREATE TABLE accounts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    approval_status ENUM('waiting', 'rejected', 'approved') DEFAULT 'waiting',
    approved_by INT DEFAULT NULL,
    status ENUM('active', 'inactive') DEFAULT 'inactive'
);

CREATE TABLE customers (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    registered_by INT
);


INSERT INTO accounts (username, password, role, approval_status, approved_by, status)
VALUES ('john123', 'password123', 'admin', 'approved', 1, 'active');

INSERT INTO accounts (username, password, role, approval_status, approved_by, status)
VALUES ('jane456', 'password456', 'admin', 'rejected', 1, 'inactive');

INSERT INTO accounts (username, password, role, approval_status, approved_by, status)
VALUES ('alex789', 'password789', 'superadmin', 'approved', 2, 'active');

INSERT INTO customers (email, first_name, last_name, avatar, registered_by)
VALUES ('john@example.com', 'John', 'Doe', 'avatar1.jpg', 1);

INSERT INTO customers (email, first_name, last_name, avatar, registered_by)
VALUES ('jane@example.com', 'Jane', 'Smith', 'avatar2.jpg', 1);

INSERT INTO customers (email, first_name, last_name, avatar, registered_by)
VALUES ('alex@example.com', 'Alex', 'Johnson', 'avatar3.jpg', 2);
