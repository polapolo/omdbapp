CREATE TABLE search (
    id INT(255) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    keyword VARCHAR(100) NOT NULL,
    page VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);