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

| Endpoint | Method | Description                   |
| -------- | ------ | ----------------------------- |
| /list    | GET    | Displays the video list page. |

# API Endpoints

| Endpoint                  | Method | Description                                      |
| ------------------------- | ------ | ------------------------------------------------ |
| /video/list?page={number} | GET    | Returns a list of videos for the specified page. |
| /video/count              | GET    | Returns the total number of videos.              |
