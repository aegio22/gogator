# Gator 🐊

Gator is a **blog aggregator CLI** written in Go. It allows users to register accounts, subscribe to RSS feeds, follow feeds, and continuously aggregate new posts into a PostgreSQL database.

This project is designed as a real-world backend-style CLI application, using PostgreSQL for persistence and modern Go tooling for database access and code generation.

---

## Requirements

Before running Gator, you must have the following installed:

### Go
Gator is written in Go and distributed as a statically compiled binary.

- Go **1.21+ recommended**
- Install from: https://go.dev/dl/

Verify installation:

```bash
go version
```

### PostgreSQL
Gator uses PostgreSQL for data storage.

- Install PostgreSQL locally or via Docker
- Create a database for Gator to use

Verify installation:

```bash
psql --version
```

---

## Tech Stack

- **Go** — CLI application
- **PostgreSQL** — persistent data storage
- **psql** — database interaction and administration
- **sqlc** — generates type-safe Go code from SQL queries

---

## Installation

Install the `gator` CLI using Go:

```bash
go install github.com/<your-github-username>/gogator@latest
```

Ensure your Go bin directory is on your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Verify installation:

```bash
gator help
```

---

## Configuration

Gator uses a config file to store the database connection string and the currently logged-in user.

### Config File Location

```text
~/.gatorconfig.json
```

### Example Config File

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "alice"
}
```

- `db_url` — PostgreSQL connection string
- `current_user_name` — the active user for CLI commands

The database must already exist before running Gator.

---

## Running Gator

### Development Mode

For development only:

```bash
go run .
```

### Production Usage

After installing or building the binary:

```bash
gator <command>
```

Go binaries are statically compiled. Once built or installed, the Go toolchain is **not required** to run Gator.

---

## Commands

### register
Register a new user.

```bash
gator register <username>
```

### login
Log in as an existing user.

```bash
gator login <username>
```

### addfeed
Add a new RSS feed with a name and URL.

```bash
gator addfeed <name> <url>
```

### follow
Follow an existing feed as the current user.

```bash
gator follow <url>
```

### unfollow
Unfollow a feed for the current user.

```bash
gator unfollow <url>
```

### feeds
Show all available feeds.

```bash
gator feeds
```

### following
Show all feeds followed by the current user.

```bash
gator following
```

### users
List all users and indicate which user is currently logged in.

```bash
gator users
```

### agg
Continuously fetch and save new posts from followed feeds.

```bash
gator agg
```

This command is typically run as a long-lived process.

### reset
Delete **all users and posts** from the database.

```bash
gator reset
```

⚠️ This is destructive and cannot be undone.

---

## Database Notes

- SQL schema and queries are written manually
- `sqlc` is used to generate type-safe Go code from SQL
- Database administration and debugging can be done via `psql`

---

## Project Structure

```text
.
├── main.go
├── internal/
├── sql/
├── go.mod
└── README.md
```

---


