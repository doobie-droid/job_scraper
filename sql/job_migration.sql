--
-- MySQL database dump
--
-- Dumped from database version 5.7 (Ubuntu 22.)
USE job_scraper;

-- Ensure database uses utf8mb4 so that complex characters and emojis from job titles are valid
ALTER DATABASE job_scraper CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- TODO
-- ADD DOCKER TO PROJECT FOR EASY SETUP AND UPDATE README FILE WHEN YOU ADD DOCKER
CREATE TABLE jobs (
    id INT NOT NULL AUTO_INCREMENT,
    url VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE,
    title VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)