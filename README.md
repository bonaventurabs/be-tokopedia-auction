# be-tokopedia-auction

## Overview

`be-tokopedia-auction` is a Go application built with the Echo framework and PostgreSQL for managing auctions in GoTo ecosystem (for hackathon purpose). It is designed to run in a Docker Compose environment.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Clone the Repository](#clone-the-repository)
  - [Docker Compose](#docker-compose)
  - [Database Migration](#database-migration)
  - [Access Server](#access-server)
- [API Documentation](#api-documentation)
- [Team Member](#team-member)

## Prerequisites

Make sure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/)

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/bonaventurabs/be-tokopedia-auction.git
cd be-tokopedia-auction
```
### Docker Compose
Run the application using Docker Compose:
```bash
docker-compose up --build
```
This will start the Echo server and PostgreSQL in separate containers.
### Database Migration
Run `init.sql` file into the postgres instance.
### Access Server
Access the Go server through:
```bash
http://localhost:8080
```
## API Documentation
Document the API endpoints, request, and response formats here.
#### Health
- <b>Endpoint</b>: `/api/v1/health`
- <b>Method</b>: `GET`
- <b>Description</b>: Check whether API is running or not.

#### Get All Items
- <b>Endpoint</b>: `/api/v1/items`
- <b>Method</b>: `GET`
- <b>Description</b>: Get all items(product) which is provided by platform.

#### Get All Items
- <b>Endpoint</b>: `/api/v1/items/:id`
- <b>Method</b>: `GET`
- <b>Description</b>: Get a specific item.

#### Bid
- <b>Endpoint</b>: `/api/v1/bid`
- <b>Method</b>: `POST`
- <b>Description</b>: Place a bid on a specific item.
- <b>Request</b>:
```form
bid_price: <int>
item_id: <int>
user_id: <int>
```

#### Post Item
- <b>Endpoint</b>: `/api/v1/items`
- <b>Method</b>: `POST`
- <b>Description</b>: Post or create a item (by seller).
- <b>Request</b>:
```form
name       : <string>
description: <string>
start_auct : <string>
end_auct   : <string>
start_price: <int>
start_price: <int>
seller_id  : <int>

```
## Team Member
<b>Group - 13</b>
- Abdul Hakim 
- M. Yuda Sulaiman 
- Bonaventura Bagas S. 