[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![project_license][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<br />
<div align="center">
  <h3 align="center">Job Scraper</h3>

  <p align="center">
    A job scraping repo to automate your job searches!
    <br />
  </p>
</div>

## About The Project

There are a few websites that provide custom solutions to scrape a bunch of job related apis. However, most of these
sites/platforms require a subscription of some sorts. The matter becomes worse, when the subscription is paid in dollars
and a lot of countries have individuals with credit cards that do not allow dollar payments.

I decided to build a job scraper pooling data from the following [SOURCES](./platforms/readme.md)

### Prerequisites

| Tool             | Version       |
|------------------|---------------|
| Golang           | ^1.23.4       |
| MySQL            | ^5.7          |
| Chromium-Browser | 133.0.6943.53 |

## Installation

1. Install dependencies in go

   ```
   go mod tidy
   ```

2. Run database migrations in mysql

- Log in to mysql

  ```
  mysql -u <username> -p

  ```

- Enter your password in the empty screen
- Run the command
  ```
  create database job_scraper;
  ```
- Run the migrations file whilst in the mysql shell
  ```
  SOURCE ./sql/job_migration.sql
  ```

3. Run the application

- Create your own .env file from the .env.example

   ```
    cp .env.example .env
   ``` 

- Start the app

   ```
   go run .
   ```

## Running Test

- Run the test within the package

  ```
  go test
  ```
  
- Run all test recursively

  ```
  go test ./...
  ```

- Run all test recursively with coverage

  ```
  go test -coverpkg=./... ./...
  ```

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/doobie-droid/job_scraper.svg?style=for-the-badge

[contributors-url]: https://github.com/doobie-droid/job_scraper/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/doobie-droid/job_scraper.svg?style=for-the-badge

[forks-url]: https://github.com/doobie-droid/job_scraper/network/members

[stars-shield]: https://img.shields.io/github/stars/doobie-droid/job_scraper.svg?style=for-the-badge

[stars-url]: https://github.com/doobie-droid/job_scraper/stargazers

[issues-shield]: https://img.shields.io/github/issues/doobie-droid/job_scraper.svg?style=for-the-badge

[issues-url]: https://github.com/doobie-droid/job_scraper/issues

[license-shield]: https://img.shields.io/github/license/doobie-droid/job_scraper.svg?style=for-the-badge

[license-url]: https://github.com/doobie-droid/job_scraper/blob/master/LICENSE

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555

[linkedin-url]: https://www.linkedin.com/in/leslie-douglas-074a6112a/
