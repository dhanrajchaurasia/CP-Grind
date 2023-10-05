# Project Name: RESTful API with Golang (Fiber) and PostgreSQL

## Table of Contents
- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license) 
## Introduction
This project is a RESTful API built using the Golang web framework Fiber and PostgreSQL as the database. It provides a foundation for building and deploying RESTful web services. You can use this project as a starting point for creating your own APIs or as a learning resource for Golang and PostgreSQL integration.

### Features
- Integration with Codeforces and AtCoder APIs: This project leverages the Codeforces and AtCoder APIs to retrieve contest data, user information, and other related data. It fetches data from these platforms, processes it, and exposes it through RESTful APIs.
- Data Modification and Presentation: The project not only fetches data from Codeforces and AtCoder but also allows you to modify and present the data in a customized way through its RESTful endpoints.
- Flexible and Extendable: You can easily extend the functionality of this project to integrate with other APIs or add additional features to suit your specific needs.

You can use this project to build applications that interact with competitive programming platforms, analyze contest statistics, or create personalized dashboards for users participating in coding competitions.

## Prerequisites
Before you begin, ensure you have met the following requirements:

- [Golang](https://golang.org/dl/): Install Golang on your system.
- [PostgreSQL](https://www.postgresql.org/download/): Install PostgreSQL and create a database for this project.
- [Git](https://git-scm.com/downloads): Install Git for version control.

## Installation
Follow these steps to set up and run the project:

1. Clone the repository:
   ```bash
   git clone https://github.com/ContriHUB/CP-Grind.git
   ```

2. Change to the project directory:
   ```bash
   cd CP-Grind
   ```

3. Install project dependencies using Go Modules:
   ```bash
   go mod tidy
   ```

4. Set up your environment variables. Create a `.env` file in the project root and add the following configuration, replacing the values with your own:

   ```ini
   # Database Configuration
   DATABASE_URL = "host=localhost user=<Database Username> password=<Database Password> dbname=<Database Name> port=<Database Port>"
   # Fiber Configuration
   PORT=3000
   ```

5. Migrate the database to create the necessary tables:
   ```bash
   go run migrate.go
   ```

6. Start the application:
   ```bash
   go run main.go
   ```

## Configuration
You can configure the application by modifying the `.env` file with your own settings. Additionally, you can customize other aspects of the application by editing the relevant code files.

## Usage
After successfully setting up the project, you can start using the RESTful API. You can use tools like [Postman](https://www.postman.com/downloads/) or `curl` to interact with the API.

## Contributing
We welcome contributions to this project. If you have any bug fixes, enhancements, or new features to propose, please open a pull request. For major changes, please open an issue first to discuss the changes you want to make.

## License
This project is licensed under the [MIT License](LICENSE).
