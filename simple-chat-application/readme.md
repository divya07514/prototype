# Simple Chat Application

## Overview
This project is a simple chat application built using Go. It uses WebSocket for real-time communication and Redis for pub/sub messaging. The application supports multiple clients and channels for message exchange.

## Features
- Real-time communication using WebSocket.
- Pub/Sub messaging with Redis.
- Multi-channel support.
- REST API for sending messages.

## Project Structure
- `simple-chat-application/main.go`: Entry point for the chat application.
- `simple-chat-application/pub_sub/redis.go`: Redis pub/sub implementation.

## Prerequisites
- Go 1.19 or later
- Redis server

## Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-folder>

### Screenshots


![client-1.png](client-1.png)![client-2.png](client-2.png)