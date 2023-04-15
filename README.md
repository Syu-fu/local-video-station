# local-video-station

[![golangci-lint](https://github.com/Syu-fu/local-video-station/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Syu-fu/local-video-station/actions/workflows/golangci-lint.yml)
[![typescript-lint](https://github.com/Syu-fu/local-video-station/actions/workflows/typescript-lint.yml/badge.svg)](https://github.com/Syu-fu/local-video-station/actions/workflows/typescript-lint.yml)

## Development Environment

To start the development environment, run the following command on the shell:
`make dev`

## Dependencies

This program depends on the following:

- Docker compose version 3 or higher
- make

# Frontend

| Endpoint    | Description                                              |
| ----------- | -------------------------------------------------------- |
| /list       | Displays the video list page.                            |
| /video/{id} | Displays the video player page with the specified video. |
| /tag        | Displays the tag creation page.                          |

# API Endpoints

| Endpoint                  | Method | Description                                             |
| ------------------------- | ------ | ------------------------------------------------------- |
| /video/list?page={number} | GET    | Returns a list of videos for the specified page.        |
| /video/count              | GET    | Returns the total number of videos.                     |
| /video/{id}               | GET    | Returns detailed information about the specified video. |
| /tag                      | POST   | Creates a new tag.                                      |
