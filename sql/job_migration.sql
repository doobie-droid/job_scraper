
--
-- MySQL database dump
--

-- Dumped from database version 5.7 (Ubuntu 22.)
CREATE TABLE jobs (
    id INT NOT NULL AUTO_INCREMENT
    url VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    INDEX idx_slug (slug)
)