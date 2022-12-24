# go-database-access

Demo repository for article on SQL database access in Go.

## Installation

Clone the repository

```bash
git clone https://github.com/woojiahao/go-database-access.git
cd go-database-access/
```

Build the project

```bash
go build
```

Create a new database in PostgreSQL named "gda" and setup the database

```bash
./gda setup
```

## Usage

Run code examples

```bash
./gda example [connect|single|multi|parameterised|null|insert|transaction|struct|return|prepared|conn|timeout]
```

