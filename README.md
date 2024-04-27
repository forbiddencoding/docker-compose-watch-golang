# Docker Compose File Watch with Golang
[![Docker](https://img.shields.io/badge/-Docker-informational?style=for-the-badge&color=27272A&logo=docker)](LICENSE.md)
[![Golang](https://img.shields.io/badge/-Golang-informational?style=for-the-badge&color=27272A&logo=go)](LICENSE.md)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge&color=27272A)](LICENSE.md)

> To use this feature, you need to be on Docker Compose version 2.22 or higher.
> To check your version, run `docker compose version`.

## Introduction

This repository contains a simple Golang application which demonstrates the use of the **Docker Compose Watch** feature.

This project exposes a simple Golang HTTP server which returns the time of the incoming request. Simply run the project
and open `http://localhost:8080` in your browser to see the current time.

Modify the `main.go` file and save it to see the file watch feature in action.

## What is "Docker Compose Watch"?

Docker Compose Watch is a file watch command which automates the update process for running Docker Compose
services as developers edit and save their work. By monitoring specified files and directories on the host machine,
Docker detects changes automatically and performs corresponding actions within the service container.

If you want to learn more, check out the [official documentation](https://docs.docker.com/compose/file-watch/).

## Usage

1. `docker compose watch` to start the Compose project
2. (optional) In another terminal, run `docker compose logs -f -t` to see the console output of the services

Alternatively you can utilize the `Makefile` with the following commands:

| Command      | Description                              |
|--------------|------------------------------------------|
| `make up`    | Starts the compose project in watch mode |
| `make down`  | Stops the compose project                |
| `make prune` | Removes all dangling images              |
| `make logs`  | Shows the logs of the running service    |

## Caveats

Since in this example, the application is created and run inside the container, every change will trigger a rebuild of
the image. This means that a lot of dangling images will be created over time. To remove them, simply run the
`make prune` command.

If some might wonder why the code is not mounted as a volume, it's because the watch feature does not monitor mounted
volume paths.

Hopefully, the Docker team will add a way to prevent this in the future.

## Alternatives

If you want a stable and tested way for live-reloading your Golang Docker container, I highly recommend taking a look at
[air](https://github.com/cosmtrek/air). It's a great tool which I use in my own projects.

## License

This project is licensed under the terms of the MIT License. Feel free to use, modify, and distribute the code
as per the terms of the license. See the [LICENSE](LICENSE.md) file for more details.

---

<p align="center">Made in the Black Forest with 🦊 in mind.</p>