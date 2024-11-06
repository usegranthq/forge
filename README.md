# UseGrant Backend

Welcome to the UseGrant Backend repository! This document will guide you through the prerequisites, development setup, and running migrations for this project.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: Make sure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/doc/install). This project requires Go version 1.23.1 or higher.
- **Docker**: Docker is required to run the database and other services in containers. You can download Docker from the [official Docker website](https://docs.docker.com/get-docker/).

## Development Setup

Follow these steps to set up the development environment:

#### Clone the repository:

```sh
git clone https://github.com/usegranthq/backend.git
cd backend
```

#### Install dependencies:

```sh
go mod tidy
```

This command will download and install the necessary Go modules for the project.

#### Start the development server:

```sh
make dev
```

This command will start the development server with live reloading enabled, using Air.

### Running Migrations

Database migrations are essential for managing changes to the database schema over time. Follow these steps to run migrations:

#### Generate the schema:

```sh
make generate
```

This command will generate the necessary code for the database schema based on the Ent schema definitions.

#### Create a new migration:

```sh
make schema name=<migration_name>
```

Replace `<migration_name>` with the name of your migration. This command will create a new migration file with the specified name.

#### Apply the migration:

```sh
make migrate name=<migration_name>
```

Replace `<migration_name>` with the name of your migration. This command will apply the migration to the database.

Happy hacking!

### Deploying

The deployment process is automated via GitHub Actions on every push to the default branch. Checkout [deploy.yml](.github/workflows/deploy.yml) for more details.

```sh
make deploy
```

Deployment happens via SSH. Checkout the [Makefile](./Makefile) for more details.
