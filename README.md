# PayX - Payment Gateway API

A robust and scalable payment gateway backend service built with Go, featuring seamless integration with Midtrans payment processing. PayX provides a comprehensive REST API for handling payment transactions, user authentication, and transaction management.

## Overview

PayX is designed to simplify payment processing for modern applications. It offers secure user authentication, flexible payment options including QRIS, and comprehensive transaction tracking capabilities. Built with performance and security in mind, PayX leverages Go's concurrency features and industry-standard security practices.

## Key Features

-   **Secure Authentication**: JWT-based user authentication and authorization
-   **Payment Processing**: Integration with Midtrans for multiple payment methods
-   **QRIS Support**: Quick Response Indonesian Standard payment processing
-   **Transaction Management**: Complete transaction history and status tracking
-   **Webhook Integration**: Real-time payment status notifications
-   **RESTful API**: Clean and intuitive API design
-   **Database Management**: PostgreSQL with GORM for efficient data handling

## Technology Stack

-   **Language**: Go 1.24.4
-   **Web Framework**: Gin
-   **ORM**: GORM
-   **Database**: PostgreSQL
-   **Payment Gateway**: Midtrans
-   **Authentication**: JWT (JSON Web Tokens)
-   **Containerization**: Docker

## Getting Started

### Prerequisites

-   Go 1.24.4 or higher
-   PostgreSQL database
-   Midtrans account (sandbox or production)

### Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/theyashdhiman04/PayX.git
    cd PayX
    ```

2.  **Install dependencies**
    ```bash
    go mod download
    ```

3.  **Configure environment variables**

    Copy `.env-example` to `.env` and fill in your configuration:
    ```bash
    cp .env-example .env
    ```

    Update the following variables in `.env`:
    -   Database connection details
    -   JWT secret key
    -   Midtrans credentials

4.  **Run database migrations**

    The application automatically runs migrations on startup. Ensure your PostgreSQL database is running and accessible.

5.  **Start the server**
    ```bash
    go run main.go
    ```

    The server will start on `http://localhost:8080` (or the port specified in your `.env` file).

## API Endpoints

-   `POST /api/v1/auth/register` - Register a new user
-   `POST /api/v1/auth/login` - Login and receive JWT token
-   `GET /api/v1/profile` - Get authenticated user profile
-   `POST /api/v1/payments/create` - Create a new payment transaction
-   `POST /api/v1/payments/qris` - Create a QRIS payment
-   `GET /api/v1/payments/status/:orderID` - Get transaction status
-   `GET /api/v1/payments/history` - Get user transaction history

## Deployment

The application can be deployed to various platforms including Render, Heroku, or any containerized environment.

## License

This project is open source and available under the MIT License.

## Author

**Yash Dhiman**
