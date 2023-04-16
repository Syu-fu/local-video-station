# local-video-station

[![golangci-lint](https://github.com/Syu-fu/local-video-station/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Syu-fu/local-video-station/actions/workflows/golangci-lint.yml)
[![typescript-lint](https://github.com/Syu-fu/local-video-station/actions/workflows/typescript-lint.yml/badge.svg)](https://github.com/Syu-fu/local-video-station/actions/workflows/typescript-lint.yml)

## How to use

1. Run the following command on the shell:

```bash
git clone https://github.com/Syu-fu/local-video-station.git
```

2. Set up the .env file.
3. Run the following command on the shell:

```bash
make dev
```

4. Access <http://{IPADDRESS}:{FRONT_PORT}/list> from a Web browser.

## Dependencies

This program depends on the following:

- Docker compose version 3 or higher
- make

## Development Environment

To start the development environment, run the following command on the shell:
`make dev`

# Frontend

| Endpoint      | Description                                              |
| ------------- | -------------------------------------------------------- |
| /list         | Displays the video list page.                            |
| /video/{id}   | Displays the video player page with the specified video. |
| /tag/create   | Displays the tag creation page.                          |
| /video/create | Displays the video creation page.                        |

# API Endpoints

| Endpoint                  | Method | Description                                             |
| ------------------------- | ------ | ------------------------------------------------------- |
| /video                    | POST   | Creates a new video.                                    |
| /video/list?page={number} | GET    | Returns a list of videos for the specified page.        |
| /video/{id}               | GET    | Returns detailed information about the specified video. |
| /video/count              | GET    | Returns the total number of videos.                     |
| /tag                      | POST   | Creates a new tag.                                      |
| /tag/list                 | GET    | Returns a list of all tags.                             |
