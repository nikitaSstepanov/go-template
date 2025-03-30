# GO TEMPLATE

This is a template for building applications in Go using the Clean Architecture principles. The project is designed to be a starting point for your next Go application, ensuring that your codebase is maintainable, scalable, and testable.

This project leverages <a href="https://github.com/gosuit">GoSuit</a> for rapid development and enhanced productivity.

## Getting Started

▎Prerequisites

• Go 1.17 or higher

• Git

▎Clone the Repository

git clone https://github.com/yourusername/GoCleanArchitectureTemplate.git
cd GoCleanArchitectureTemplate
▎Install Dependencies

Make sure to install the required dependencies using Go Modules:

go mod tidy
▎Build the Application

To build the application, run:

go build -o myapp ./cmd/myapp
▎Run the Application

You can run the application using:

./myapp
By default, the application will start on port 8080. You can change the port by setting the PORT environment variable.

▎Running Tests

To run the tests, simply execute:

go test ./...
▎Project Structure

The project is organized into several key directories:

• cmd/: Contains the entry point for the application.

• internal/: Contains the core business logic and domain models.

• pkg/: Contains shared libraries and utilities.

• configs/: Configuration files.

• migrations/: Database migrations.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.