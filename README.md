# ETH Fee Calculator

An application to calculate ETH fees

## Running

- **Option 1**: to run the application asap you can use docker-compose, as the images have already been built and uploaded to docker hub:
```shell
    make docker_run
```

- **Option 2**: to run the application using go:
```shell
    make run
```

## Documentation

- Swagger as the specification to document the API
- After running the application, you can check the webpage at:
    - http://localhost:8080/swagger/index.html

## Project's structure

```shell script
📦eth-fee-calculator
 ┣ 📂cmd
 ┃ ┗ 📂rest
 ┃ ┃ ┗ 📜main.go
 ┣ 📂docs
 ┣ 📂internal
 ┃ ┣ 📂application
 ┃ ┣ 📂domain
 ┃ ┣ 📂infra
 ┃ ┃ ┣ 📂config
 ┃ ┃ ┗ 📂repository
 ┃ ┗ 📂port
 ┃ ┃ ┗ 📂rest
 ┃ ┃ ┃ ┣ 📂handler
 ┃ ┃ ┃ ┣ 📂presenter
 ┗ ┗ ┗ ┗ 📜api.go
```

## Code quality
- **golangci-lint**: 0 lint errors
- **unit testing**: 97.1% of code coverage

## Decisions

- **Design**: Clean architecture
- **Unit testing**: the libraries onsi/gomega and golang/mock are being used because their facility and popularity
  - TODO:
    - The health check could be better tested, adding test for success path
- **Documentation**: Swagger GUI using swaggo and hosted on https://{host}/swagger
- **Errors**: the logging is catching all errors and for the client can vary on response:
    - if it is an internal error only, respond with a status code 500 without details
    - if it is a client error, respond with a status code 4xx
- **Performance**: this dataset is too large to be calculated every time it is requested. For the purpose of this project, it is fine, but it would make more sense to pre-calculate the data and consume from this prepared dataset. 
- **Logging**: used logrus as it is quite complete and popular:
  - TODO:
    - standardization of the logs, like putting them in a json format
    - track more data about request
