# Short URL Service

A simple short URL service similar to Bit.ly or TinyURL. This service allows you to convert long URLs into short links and supports click tracking and redirection.

## Features

### 1. Generate Short URL
- The user inputs a long URL (e.g., `https://example.com/some/very/long/url`), and the system generates a short URL (e.g., `https://sho.rt/abc123`).

### 2. URL Redirection
- When a user accesses the short URL, the system automatically redirects the user to the original long URL.

### 3. Click Tracking (Optional)
- The system records how many times the short URL has been clicked, as well as the user's IP address and other statistical information.

## Microservices Architecture

To improve flexibility and scalability, the system is divided into three microservices, each responsible for one feature module:

### 1. **Shorten Service**
Responsible for receiving user requests for long URLs, generating and returning short URLs.

- Provides a RESTful API: `POST /shorten` to accept a long URL and return a short URL.

### 2. **Redirect Service**
Handles user visits to short URLs by looking up the original long URL and performing the redirection.

- Provides a RESTful API: `GET /:shortLink` where users access the short link and are redirected to the original URL.

### 3. **Stats Service (Optional)**
Records click statistics such as the number of clicks and user IP addresses.

- Provides a RESTful API: `GET /stats/:shortLink` to query the click statistics for a specific short link.

## Tech Stack

### 1. **Gin**
- A high-performance web framework for Go, used to build RESTful APIs.

### 2. **GORM**
- An ORM framework for Go that simplifies database operations, making it easier to interact with relational databases.

### 3. **Redis**
- Used to store mappings between short URLs and original long URLs. Redis provides fast read access, making it ideal for high-concurrency services.

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/your-username/short-url-service.git
cd short-url-service
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the services

Start each microservice individually:

- **Shorten Service**: `go run shorten-service/main.go`
- **Redirect Service**: `go run redirect-service/main.go`
- **Stats Service**: `go run stats-service/main.go`

## Usage Example

### 1. Generate a Short URL

```bash
curl -X POST -d '{"long_url": "https://example.com/some/very/long/url"}' \
     -H "Content-Type: application/json" \
     http://localhost:8080/shorten
```

Response:

```json
{
  "short_url": "https://sho.rt/abc123"
}
```

### 2. Redirect to Original URL

When accessing the short URL, the system will redirect to the original long URL:

```bash
curl -L http://localhost:8080/abc123
```

### 3. View Click Stats (Optional)

```bash
curl http://localhost:8080/stats/abc123
```

Response:

```json
{
  "click_count": 42,
  "ip_addresses": ["192.168.1.1", "192.168.1.2"]
}
```
