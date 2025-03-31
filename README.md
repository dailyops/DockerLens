# Docker Image Inspector

A simple Go utility to inspect Docker images and gather useful metadata, including:

- Base Docker Image
- Python Version (if available)
- Pip Version (if available)
- Size of the Docker Image

## Prerequisites

- Docker installed and running
- Go installed (version 1.16+ recommended)

## Installation

1. Clone the repository:

```
git clone https://github.com/yourusername/docker-image-inspector.git
cd docker-image-inspector
```

2. Build the utility:

```
go build -o docker-image-inspector
```

## Usage

Run the utility from the terminal:

```
./docker-image-inspector
```

You will be prompted to enter the Docker image name. Example:

```
Enter Docker image name: python:3.9-slim

Docker Image Inspector
Property              Value
Image                 python:3.9-slim
Size                  300MB
Base Image            debian:buster-slim
Python Version        Python 3.9.7
Pip Version           pip 21.2.4
```

## Features

- Automatically detects Python and Pip versions if available inside the image.
- Displays the base image, size, and other useful metadata.
- Lightweight and easy to use.

## Contributing

Feel free to open issues or submit pull requests for new features and improvements!

## License

This project is licensed under the MIT License.
