# Go Voucher API

A RESTful API for managing brands, vouchers, and redemptions, built with Go using the Gin framework and GORM for database interaction.

## Features

- Manage **Brands**
- Create, retrieve, and filter **Vouchers**
- Handle **Redemptions** and **Transactions**
- PostgreSQL database integration
- Input validation using `go-playground/validator`

## Prerequisites

- [Go](https://go.dev/) (1.18+)
- [PostgreSQL](https://www.postgresql.org/)
- [Git](https://git-scm.com/)

## Installation

1. Clone the repository:

2. Install dependencies:



```go mod tidy
  ```


3. Set up the database:
- Create a PostgreSQL database (e.g., `testgo`).
- Update the connection string in the `initDB` function inside `main.go`:

  ```go
  dsn := "host=localhost user=postgres password=password dbname=testgo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  ```

4. Run the server

    ```go run main.go```

The server will start at `http://localhost:8080`.

## API Endpoints

### Brand Endpoints

#### Create a Brand
- **POST** `/brand`
- **Request Body**:
```json
{
 "name": "Brand Name"
}

#### Create a Voucher
- **POST** `/voucher`
- **Request Body**:
```json
{
  "brand_id": 1,
  "name": "Voucher 10K",
  "cost_in_point": 10000
}

#### Get a Voucher
- **GET** `/voucher/brand?id={id}`

#### Create Redemption
- **POST** `/transaction/redemption`
- **Request Body**:
```json
{
  "vouchers": [
    {
      "id": 1,
      "cost_in_point": 10000
    }
  ]
}

#### Get Redemption
- **GET** `/transaction/redemption?transactionId={id}`




