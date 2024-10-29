# Eco App

Eco App is a backend application built with Golang, Go Fiber, and PostgreSQL. It provides features to manage users, products, and articles, with JWT authentication, Redis caching, pagination, and a Clean Code architecture (Controller, Service, Repository) to maintain code organization.

## Tech Stack

- **Backend:** Golang with Go Fiber
- **Database:** PostgreSQL
- **Caching:** Redis
- **Authentication:** JWT (JSON Web Tokens)
- **Design Pattern:** Clean Code (Controller, Service, Repository)

## Features

1. **User Management**
   - Authentication with JWT.
   - User registration and login.
   - Admins can view, update, and delete user data.

2. **Product Management**
   - CRUD (Create, Read, Update, Delete) for products.
   - Supports pagination for product lists.
   - Products include attributes such as name, brand, description, category, price, and eco-score certification.

3. **Article Management**
   - CRUD for articles, including title, image, content, and author management.
   - Supports pagination for article lists.

4. **Redis Caching**
   - Utilizes Redis to cache data and speed up access to frequently requested information.

5. **Pagination**
   - Supports pagination for products and articles to optimize data management.

6. **Clean Code Architecture**
   - Separates application logic into Controller, Service, and Repository layers for ease of development and maintenance.

## Endpoint Structure

### Auth Routes
- `POST /register` - Register a new user.
- `POST /login` - Login and obtain a JWT token.

### User Routes (Admin Only)
- `GET /users` - Retrieve a list of all users (admin authentication required).
- `GET /users/:id` - Retrieve user details by ID.
- `PATCH /users/:id` - Update user data by ID.
- `DELETE /users/:id` - Delete user data by ID.

### Product Routes
- `POST /product` - Add a new product (admin authentication required).
- `GET /product` - Retrieve a paginated list of products.
- `GET /product/:id` - Retrieve product details by ID.
- `DELETE /product/:id` - Delete a product by ID (admin authentication required).

### Article Routes
- `POST /article` - Add a new article (admin authentication required).
- `GET /article` - Retrieve a paginated list of articles.
- `GET /article/:id` - Retrieve article details by ID.
- `PATCH /article/:id` - Update an article by ID.
- `DELETE /article/:id` - Delete an article by ID (admin authentication required).

## Data Structure

### Installation and Configuration

#### Clone the Repository

```bash
git clone https://github.com/username/eco-app.git
cd eco-app
```

#### Configure the Database

Create a PostgreSQL database and update the `.env` file with your configuration details:

```makefile
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdatabase
```

#### Configure Redis

Update the `.env` file with your Redis configuration:

```makefile
REDIS_HOST=localhost
REDIS_PORT=6379
```

#### Install Dependencies

```bash
go mod tidy
```

#### Run the Application

```bash
go run main.go
```

## Directory Structure

```bash
├── controllers/       # Controller layer for handling requests
├── services/          # Service layer for business logic
├── repositories/      # Repository layer for database access
├── models/            # Data models representing database structures
├── config/            # Configuration files for database, Redis, and more
└── main.go            # Application entry point
```

## Contribution

If you'd like to contribute to this project, please fork the repository and submit a pull request. Thank you!
