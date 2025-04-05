# GO TEMPLATE

This is a template for building applications in Go using the clean architecture principles. The project is designed to be an example of modern and high-quality app. It can be starting point for your next Go application, ensuring that your codebase is maintainable and scalable.

The template contains the functionality of working with the profile and authorization

This project uses <a href="https://github.com/gosuit">GoSuit</a> for flexible and concise development.

## Getting Started

### Run
You can run app local:

```zsh
make run
``` 

Or in docker:

```zsh
make docker
```

### Configuration

To configure the configuration type, you must set the ENVIRONMENT variable ("LOCAL" or "DOCKER"). Default is "LOCAL".

If "LOCAL" is selected, you should configure the .env file for environment variables (example in .env.example).

If the ENVIRONMENT variable is not "LOCAL", the environment variables will be taken from the machine environment.

### Test

You can run unit test with:

```zsh
make unit-test
```

And e2e tests (./tests directory):

```zsh
make e2e-test
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.