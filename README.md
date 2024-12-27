# Beego Go Channel

This project is a Beego-based web application designed for a basic channel management system. It includes features such as controller management, routing, and simple views for user interaction.
## Description
This project is based on the Cat API, which provides access to a vast collection of cat images, facts, and breeds. The application allows users to interact with the API to display random cat images, search for cats by breed, and view information about different breeds.

### Key Functionalities:
Fetch and display random cat images.
Search for cats based on breed.
View breed details with images and descriptions.
This project uses Beego for routing and handling API requests, creating a simple and engaging interface for cat lovers!

## Features
- Beego-based web framework
- Basic routing and controller logic
- Static file handling
- Testing setup for controllers and routes

  ### set up Beego from scratch, follow these steps:

1. Install Go (if not already installed)
Download and install Go from the official site: https://golang.org/dl/.

2. Install Beego
Run the following command to install Beego:
```bash
go install github.com/beego/beego/v2@latest
```
3. Create Your Beego Project
Use the Beego command to create a new project:

```bash

beego new myproject
cd myproject
```
3. Run the application
```bash
go run main.go
```
or
```bash
bee run
```
The app will be available at http://localhost:8080.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/samiya1859/beegogochannel.git
   cd beegogochannel
   ```
2. Install dependencies:
   ```bash
   go mod tidy
```
3. Run the application
```bash
go run main.go
```
or
```bash
bee run
```
The app will be available at http://localhost:8080.

## Testing

To run tests, execute the following:
```bash
go test ./tests
```
## License
This project is open-source and available under the MIT License.
