# POS System Project

This project is a Point-of-Sale (POS) system designed for practice and use in a store. It is built using Golang with hexagonal architecture for the backend and PostgreSQL as the database. Below is a guideline to help you practice coding and build the system step by step.

## Features
1. **Inventory Management**: Add items, check stock.
2. **Point-of-Sale**: Create orders.
3. **Invoice Management**: Save invoices as PDF or other formats.

## Updated Folder Structure

```
wholesale-pos/
├── backend/            # Backend code (Golang with hexagonal architecture)
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   ├── configs/
│   ├── pkg/
│   └── go.mod
├── frontend/           # Frontend code (React or Vue.js)
│   ├── public/
│   ├── src/
│   ├── package.json
│   └── README.md
├── docs/               # Documentation
└── README.md           # Project guideline
```

## Why Separate Backend and Frontend?
1. **Independent Development**: You can work on the backend and frontend independently.
2. **Scalability**: Easier to scale and maintain as the project grows.
3. **Deployment**: You can deploy the backend and frontend to different servers or services if needed.

## Getting Started

### Prerequisites
1. Install [Golang](https://golang.org/dl/).
2. Install [PostgreSQL](https://www.postgresql.org/download/).
3. Install a tool for database migrations (e.g., [golang-migrate](https://github.com/golang-migrate/migrate)).

### Step-by-Step Guide

1. **Set Up the Project**:
   - Initialize a Go module: `go mod init wholesale-pos`.
   - Create the folder structure as shown above.

2. **Database Setup**:
   - Create a PostgreSQL database.
   - Write migration files in the `backend/migrations/` folder.

3. **Domain Layer**:
   - Define entities like `Item`, `Order`, and `Invoice` in the `backend/internal/domain/` folder.

4. **Application Layer**:
   - Implement use cases like `AddItem`, `CreateOrder`, and `GenerateInvoice` in the `backend/internal/application/` folder.

5. **Adapters**:
   - Create adapters for database interaction, REST API, etc., in the `backend/internal/adapters/` folder.

6. **Ports**:
   - Define interfaces for the adapters in the `backend/internal/ports/` folder.

7. **Configuration**:
   - Add configuration files (e.g., `config.yaml`) in the `backend/configs/` folder.

8. **Frontend Recommendation**:
   - Use [React](https://reactjs.org/) or [Vue.js](https://vuejs.org/) for the frontend. Both are modern and widely used frameworks.

9. **Testing**:
   - Write unit tests for each layer.

10. **Run the Application**:
    - Use `go run` to start the application.

## Additional Notes
- Follow the principles of hexagonal architecture to keep the code modular and testable.
- Use environment variables for sensitive configurations like database credentials.

Happy coding!