# Snippetbox

A sample project from Alex Edward's book – [Let's Go](https://lets-go.alexedwards.net).

## Setup instructions

### 1. Configure Environment Variables

Create a .env file in the root directory and specify the required environment variables. Use the `.env.example` file as a reference:

```sh
cp .env .env.example
```

Update the `.env` file with the appropriate values.

### 2. It is dockerized. To run it install and run docker, then use:

```sh
docker compose up
```

in the root folder. To run it in detached mode use `-d` flag. To run it in watch mode use the `--watch` flag.

For other configurations look in the `docker-compose.yaml` and `init.sql` files.
