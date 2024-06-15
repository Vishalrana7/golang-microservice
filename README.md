# Cat-Fact Microservice with Logging

## Overview

This project is a Go-based microservice that provides cat facts. The service includes logging functionality and is containerized using Docker. The project uses GitHub Actions for continuous integration, ensuring code quality through linting and testing, and automatically building and pushing Docker images to DockerHub.

## Project Description

The Cat-Fact Microservice is a lightweight, efficient service built with Go, designed to provide random cat facts upon request. This project was created with a focus on simplicity and scalability, ensuring that it can handle multiple requests efficiently while maintaining a minimal resource footprint. The microservice includes robust logging capabilities to monitor its operation and facilitate debugging. Containerization with Docker makes the deployment process seamless, allowing the service to run consistently across different environments. By leveraging GitHub Actions for continuous integration, this project ensures that all code changes are automatically tested and validated, maintaining high code quality and reliability. The integration of Docker Compose further simplifies the orchestration and management of the application and its dependencies, streamlining the development and deployment workflow.

## Features

- **Cat Facts**: Provides random cat facts.
- **Logging**: Integrated logging for monitoring and debugging.
- **Docker**: Containerized application for easy deployment.
- **CI/CD**: Automated workflow for linting, testing, and Docker image management using GitHub Actions.

## Getting Started

### Prerequisites

- Go 1.22 or later
- Docker
- Docker Compose
- GitHub Account (for accessing repository and setting up GitHub Actions)

### Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/Vishalrana7/golang-microservice.git
    cd golang-microservice
    ```

2. **Build and run the Docker container**:
    ```sh
    docker compose up --build
    ```

### Usage

Once the service is up and running, you can access it at `http://localhost:3000`. The service will provide random cat facts.

### Project Structure

- **main.go**: Entry point of the application.
- **handlers.go**: Contains HTTP handlers.
- **logger.go**: Logging setup and middleware.
- **Dockerfile**: Instructions to build the Docker image.
- **docker-compose.yml**: Docker Compose configuration.

## Docker

### Building the Docker Image

To build the Docker image manually, run:
```sh
docker build -t visualzero007/my-go-repo:latest .
