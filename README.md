# LeetCode Badge Generator

This project is a simple web application that allows users to generate a custom LeetCode badge URL by entering their LeetCode username. The badge displays the user's statistics and can be easily shared.

## Example
[![LcBadge](http://64.176.66.30:8080/api/slug/ZORRIX/badge.svg)](http://64.176.66.30:8080)

## Features

- **Username Input:** Users can enter their LeetCode username to generate a badge URL.
- **Caching of user data:** data is stored in the postgres database, and updated if statistics were saved more than 15 minutes ago.
- **Copy URL Button:** A button to easily copy the generated badge URL.
## Prerequisites

Before you begin, ensure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Golang](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Make](https://www.gnu.org/software/make/)

## Starting project locally

```bash
make env-up
make app-build
make migrate-psql-up
make app-start
````

## Usage
Once the application is running, navigate to http://localhost:8080 in your web browser. Enter your LeetCode username in the input field, click "Get Badge Link," and you will be provided with a URL to your custom badge. You can also copy the URL using the "Copy URL" button.

## Author
Created by Andrew Pavlov.


