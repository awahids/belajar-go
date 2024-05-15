Here's a comprehensive `README.md` file for your Go project. This documentation covers the use of Gin Gonic, GORM, and PostgreSQL, along with instructions for running seeders and using `CompileDaemon`.

```markdown
# Belajar Go Project

This project is a Go-based web application utilizing Gin Gonic for the web framework, GORM for ORM, and PostgreSQL as the database. It includes features such as user and role management with JWT-based authentication and more.

## Project Structure

```
belajar-go/
│
├── cmd/
│     └── main.go
│
├── configs/
│     └── dbConfig.go
│
├── internal/
│     ├── delivery/
│     │     ├── handlers/
│     │     │     └── bookHandler/
│     │     │               └── bookHandler.go
│     │     │
│     │     ├── data/
│     │     │     ├── request/
│     │     │     │     └── bookReq/
│     │     │     │          └── bookRequest.go
│     │     │     │
│     │     │     └── response/
│     │     │          ├── bookRes
│     │     │          │     └── bookResponse.go
│     │     │          └── response.go
│     │     │     
│     │     └── router/ 
│     │           ├── bookRouter/
│     │           │          └── bookRouter.go
│     │           │
│     │           └── router.go
│     │
│     ├── domain/
│     │     ├── models/
│     │     │     └── books.go
│     │     │
│     │     ├── repositories/
│     │     │     ├── bookRepo/
│     │     │     │      └── bookRepo.go
│     │     │     │
│     │     │     └── repoInterface/
│     │     │            └── bookRepoInterface.go
│     │     │
│     │     └── services/
│     │           ├── bookService/
│     │           │      └── bookService.go
│     │           │
│     │           └── serviceInterface/
│     │                  └── bookServiceInterface.go
│     │         
│     └── infrastructure/
│           └── database/
│                  ├── database.go
│                  └── migrations.go
│
├── pkg/
│     ├── utils/
│     │     └── base.go
│     │
│     └── helpers/
│           └── errorPanic.go
│
├── .env.example
├── .gitignore
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL
- `CompileDaemon` for live reloading during development

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/belajar-go.git
   cd belajar-go
   ```

2. Copy the example environment variables file and configure it:
   ```bash
   cp .env.example .env
   ```

3. Install the dependencies:
   ```bash
   go mod download
   ```

### Configuration

Update the `.env` file with your database credentials and other configurations.

### Database Setup

1. Create the database:
   ```bash
   createdb belajar_go
   ```

2. Run the database migrations:
   ```bash
   go run internal/infrastructure/database/migrations.go
   ```

### Running the Application

You can use `CompileDaemon` to automatically rebuild and restart the application when you make changes to the source code.

1. Install `CompileDaemon`:
   ```bash
   go install github.com/githubnemo/CompileDaemon@latest
   ```

2. Start the application:
   ```bash
   CompileDaemon --build="go build cmd/main.go" --command=./main
   ```

### Running Seeders

Seeders are scripts to populate the database with initial data. To run the seeders:

```bash
go run internal/infrastructure/database/seeder.go
```

### API Documentation

#### Endpoints

##### Books

- `GET /books` - Get all books
- `POST /books` - Create a new book
- `GET /books/:id` - Get a book by ID
- `PUT /books/:id` - Update a book by ID
- `DELETE /books/:id` - Delete a book by ID

### Project Structure Explanation

- **cmd/**: Entry point of the application.
- **configs/**: Configuration files.
- **internal/**: Internal application code.
  - **delivery/**: Contains the HTTP handlers and routers.
  - **domain/**: Contains the business logic and models.
  - **infrastructure/**: Contains the database setup and migrations.
- **pkg/**: Utility packages and helpers.
- **.env.example**: Example environment variables file.
- **go.mod**: Go module file.
- **go.sum**: Go module dependencies file.

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

### Notes:

1. Replace `your_project_path` with the actual import paths in your project.
2. Customize the `README.md` content based on your actual application features and endpoints.