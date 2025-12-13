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

```bash
go install github.com/aott33/gator@latest
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
gator register <name>    # Create a new user
gator login <name>       # Login as a user
gator users              # List all users
gator addfeed <name> <url>  # Add an RSS feed
gator feeds              # List all feeds
gator follow <url>       # Follow a feed
gator following          # List feeds you follow
gator unfollow <url>     # Unfollow a feed
gator agg <interval>     # Start aggregating feeds (e.g., "1m" for every minute)
gator browse [limit]     # View collected posts
```
