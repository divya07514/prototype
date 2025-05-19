# Message Broker System

## Overview

This project implements a simple message broker system in Go. It allows for publishing messages to specific topics and subscribing to those topics to receive messages. The system is thread-safe and supports multiple subscribers per topic.

## Features

- **Publish/Subscribe Model**: Supports publishing messages to topics and subscribing to topics to receive messages.
- **Thread-Safe**: Ensures safe concurrent access using mutexes.
- **Dynamic Subscription Management**: Allows adding and removing subscribers dynamically.

## Project Structure

- `message_broker/model/message.go`: Defines the `Message` struct used for publishing data.
- `message_broker/subscriber/subscriber.go`: Defines the `Subscriber` struct and methods for managing subscriber channels.
- `message_broker/broker/broker.go`: Implements the core `Broker` functionality, including subscribing, unsubscribing, and publishing messages.

## Usage

### Broker

1. **Create a Broker**:
   ```go
   broker := broker.NewBroker()