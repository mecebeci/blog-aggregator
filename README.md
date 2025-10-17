# Blog Aggregator
---
This is a command-line **Blog Aggregator** built in Go.  
It allows users to register, add RSS feeds, follow other users’ feeds, and browse posts — all from the terminal.  
The project uses **PostgreSQL** for data persistence and a simple **CLI command system** for interaction.

---

## Requirements

Before running the project, make sure you have:

- **Go 1.23+** installed
    
- **PostgreSQL** installed and running   

---

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/blog-aggregator.git
cd blog-aggregator
```

### 2. Set up the database

Create a PostgreSQL database and run migrations using [Goose](https://github.com/pressly/goose):

```bash
createdb blog_aggregator
goose -dir ./sql/schema postgres "postgres://username:password@localhost:5432/blog_aggregator?sslmode=disable" up
```

---

## Running the App

You don’t need to install a binary — you can simply run commands using:

```bash
go run . <command> [args...]
```

Example:

```bash
go run . register mecebeci
go run . login mecebeci
go run . addfeed "Hacker News" "https://hnrss.org/newest"
go run . feeds
go run . browse
```

---

## Available Commands

|Command|Description|Login Required|
|---|---|---|
|`register <username>`|Register a new user|No|
|`login <username>`|Log in as an existing user|No|
|`reset`|Reset all local configuration (clears logged-in user)|No|
|`users`|List all registered users|No|
|`agg`|Start the feed aggregator (fetches and updates posts)|No|
|`addfeed <name> <url>`|Add a new RSS feed|Yes|
|`feeds`|List all feeds and their owners|No|
|`follow <feed_url>`|Follow a feed by URL|Yes|
|`following`|Show feeds you’re currently following|Yes|
|`unfollow <feed_url>`|Unfollow a feed|Yes|
|`browse`|Browse latest posts from your followed feeds|Yes|

---

## Config File

After logging in or registering, a config file is automatically created here:

```
~/.config/blog-aggregator/config.json
```

Example content:

```json
{
  "current_user_name": "mecebeci"
}
```

This file keeps track of the currently logged-in user.

---
