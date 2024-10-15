# Blog Aggregator CLI

Blog Aggregator is a command-line tool for fetching and aggregating RSS feeds. It allows users to register, log in, add RSS feeds, follow other users' feeds, and browse posts from feeds they're following. It supports PostgreSQL for database management.

## Requirements

To run Gator, make sure you have the following installed:

- **Go**: Blog Aggregator is written in Go. You can download Go from [here](https://golang.org/dl/).
- **PostgreSQL**: Blog Aggregator uses PostgreSQL as the database. Download PostgreSQL from [here](https://www.postgresql.org/download/).

## Installation

To install Blog Aggregator, clone the repository and install the CLI using `go install`:

```bash
git clone https://github.com/<your-username>/blog-aggregator.git
cd blog-aggregator
go install