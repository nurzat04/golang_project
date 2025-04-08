Stats Service
A lightweight microservice for tracking the number of clicks on short URLs — part of a URL shortening system (like Bit.ly).

Features：
Track how many times a short URL has been visited.
Expose an HTTP API to get the click count.
Uses Redis for fast, in-memory counting.

Tech Stack：
Go + Gin (web framework)
Redis (storage)
GORM (optional for DB usage)
