# Gator

A CLI RSS feed aggregator written in Go. Gator allows you to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Prerequisites

You need the following installed:

- [Go](https://golang.org/dl/) (1.25 or later)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/aott33/gator.git
cd gator
```

2. Build the application:
```bash
go build -o gator
```

3. Optionally, install it to your `$GOPATH/bin`:
```bash
go install
```

## Configuration

Create a config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

Replace `username`, `password`, and `gator` with your PostgreSQL credentials and database name.

## Commands

```bash
gator register <name>       # Create a new user
gator login <name>          # Login as a user
gator users                 # List all users
gator addfeed <name> <url>  # Add an RSS feed
gator feeds                 # List all feeds
gator follow <url>          # Follow a feed
gator following             # List feeds you follow
gator unfollow <url>        # Unfollow a feed
gator agg <interval>        # Start aggregating feeds (e.g., "1m" for every minute)
gator browse [limit]        # View collected posts - limits to 2 by default
```

## Extending the Project

The core functionality is complete, but here are some optional ideas I would like to implement:

- Add sorting and filtering options to the browse command
- Add pagination to the browse command
- Add concurrency to the agg command so that it can fetch more frequently
- Add a search command that allows for fuzzy searching of posts
- Add bookmarking or liking posts
- Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
- Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
- Write a service manager that keeps the agg command running in the background and restarts it if it crashes
