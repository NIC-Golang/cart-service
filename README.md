# Cart Service
Cart Service is a microservice for managing shopping carts in an online store. The project is built on Go using Redis for data storage and caching.

## Main Features
- Add products to the cart.
- Remove products from the cart.
- Update product quantity in the cart.
- View the contents of the cart.
- Clear the cart.
- Caching cart data using Redis.

## Requirements:
- Golang 1.23+
- Docker & Docker Compose
- Redis (runs in a Docker container)  

## Setup Instructions  

Clone the repository and set up the environment:  

```sh
git clone https://github.com/NIC-Golang/cart-service.git
cd cart-service
cp .env.example .env
```
## Available Makefile Commands:
```sh
make help       # Show available commands  
make install    # Install dependencies (go mod tidy)  
make run        # Run the server  
make stop       # Stop the server  
make restart    # Restart the server  
make compile    # Compile the application  
make clean      # Clean the build cache  
```

## Running the service:
Make sure Docker Engine is running, then execute:
```
docker compose build
docker compose up
```
To stop the service:
```
docker compose down
```
