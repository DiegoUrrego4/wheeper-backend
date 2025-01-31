# Wheeper - Backend
Wheeper is Colombia's pioneering supermarket price comparison platform. Our name combines "Where is cheaper?" into "Wheeper" - reflecting our mission to help consumers find the best prices across Colombian supermarkets.

## Overview
Wheeper empowers consumers to make informed purchasing decisions by comparing prices across multiple supermarkets in real-time. Our platform simplifies the shopping experience, helping users save both time and money.

## Features
- RESTful API for price comparisons
- Real-time supermarket data processing
- Secure authentication system
- PostgreSQL database integration
- Docker containerization

## Technical Requirements
- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL 12.22
- Git

## Installation
### Development Setup
1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/wheeper-backend.git
   cd wheeper-backend
2. Install project dependencies ```go mod tidy```
3. Configure environment variables
   - Copy the environment template file ```cp .env-template .env```
   - Update the ```.env``` file with your configuration
4. Start the PostgreSQL database: ```docker-compose up -d```
5. Run the application: 
```bash
go run cmd/main.go
```

## Usage


## Contributing
- Fork the repository
- Create your feature branch (git checkout -b feature/AmazingFeature)
- Commit your changes (git commit -m 'Add some AmazingFeature')
- Push to the branch (git push origin feature/AmazingFeature)
- Open a Pull Request
