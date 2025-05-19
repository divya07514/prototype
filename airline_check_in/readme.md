# Airline Check-In System

## Overview

The Airline Check-In System is a prototype application that simulates seat booking for airline passengers. It demonstrates various approaches to handle seat allocation, including handling concurrency and database locking mechanisms.

## Features

- **Seat Layout Management**: Visual representation of seat availability.
- **User Management**: Fetch user details from the database.
- **Seat Booking**: Multiple approaches to seat booking, including random allocation and concurrency handling.
- **Database Integration**: Uses MySQL for storing user and seat data.

## Directory Structure

- `service/`: Contains core logic for seat layout management, user and seat retrieval, and utility functions.
- `approach_0/`: Basic implementation where users manually select seats without conflict detection.
- `approach_1/`: Random seat allocation for a single user.
- `approach_2/`: Simulates 120 users booking seats simultaneously without concurrency control.
- `approach_3/`: Simulates 120 users booking seats with first-available seat selection but no database locking.
- `approach_4/`: Adds database row locking using `FOR UPDATE` to handle concurrency.
- `approach_5/`: Improves concurrency handling with `FOR UPDATE SKIP LOCKED` for fairness.
- `seats.md`: SQL script for creating and populating the `seats` table.
- `users.md`: SQL script for creating and populating the `users` table.

## Prerequisites

- Go 1.20 or later
- MySQL database
- `github.com/go-sql-driver/mysql` Go package

## Setup

1. Clone the repository and navigate to the project directory.
2. Set up the database:
    - Use the SQL scripts in `seats.md` and `users.md` to create and populate the `seats` and `users` tables.
3. Update the database connection string in the `dbConn` function in each `main.go` file to match your MySQL setup.

## How to Run

Each approach can be run independently. Navigate to the respective directory and execute the `main.go` file.

### Example: Running `approach_0`

```bash
cd approach_0
go run main.go <user_id>