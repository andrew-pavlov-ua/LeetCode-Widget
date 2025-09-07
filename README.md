# LeetCode Badge Generator

This project is a web application that allows users to generate a custom LeetCode badge URL by entering their LeetCode username. The badge displays the user's lc statistics lc profile visits and can be easily shared.

### Project web page: https://lc.andrewpavlov.org/

## Example
[![LeetCode Badge](https://lc.andrewpavlov.org/api/slug/MURASAME_/badge.svg)](https://lc.andrewpavlov.org/redirect-page/murasame_)

## Features

- **Username Input:** Users can enter their LeetCode username to generate a badge URL.
- **Caching of user data:** data is stored in the postgres database, and updated if statistics were saved more than 15 minutes ago.
- **Copy URL Button:** A button to easily copy the generated badge Markdown and HTML.
- **Counting your lc acc's visits:** every time someone clicks on a badge, app is counting it and saves to db by time period.
## Prerequisites

Before you begin, ensure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Golang](https://go.dev/doc/install)
- [Make](https://www.gnu.org/software/make/)

## Starting project locally

```bash
make env-up
make app-build
make migrate-psql-up
make app-start
````

## Usage
Once the application is running, navigate to http://localhost in your web browser. Enter your LeetCode username in the input field, click "Get Badge Link," and you will be provided with a URL to your custom badge. You can also copy the URL using the "Copy URL" button.

## (Optional) Static, serverless fallback

If you ever need a zero-backend version (e.g., Netlify only), you can generate a simple badge via shields.io and link it to your profile:

## Author
Created by Andrew Pavlov.

