# News API System with JWT Authentication

This project is a **Go-based web application** designed to provide a secure system for consuming news via an API, protected by **JWT authentication**. It includes features like token renewal, session management, and integration with external news APIs. The project is structured to ensure scalability and follows best practices for Go development.

---

## Features

### Authentication:
- **Login and Logout** with JWT-based security.
- **Token Renewal** for active sessions.
- Middleware to **protect private routes**.
- **Blacklist tokens** for secure logout functionality.

### News API:
- Fetch a list of news articles from external sources.
- Retrieve a specific news article by ID.
- Mock integration with free external news APIs.

### Web Interface:
- Pages for **login** and **logout** styled with Bootstrap.
- Protected access to news pages using JWT.

### Testing:
- Unit tests for:
  - JWT authentication.
  - Token renewal.
  - Handlers for the News API (success and error cases).

---

## Installation and Setup

### Prerequisites:
1. **Go**: Ensure Go is installed ([Download Go](https://go.dev/dl/)).
2. **Git**: Make sure Git is installed on your system.
3. (Optional) **Heroku CLI** or other deployment CLI tools if deploying to the cloud.

### Steps:
1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/news-api-system.git
   cd news-api-system
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application locally:
   ```bash
   go run main.go
   ```
4. Access the application at `http://localhost:8080`.

---

## Project Structure

```plaintext
news-api-system/
├── handlers/           # Handlers for APIs and routes
├── middleware/         # Middleware for JWT and route protection
├── models/             # Data models (e.g., News, User)
├── templates/          # HTML templates for web pages
├── tests/              # Unit and integration tests
├── main.go             # Entry point of the application
├── go.mod              # Go modules file
└── Procfile            # For deployment on Heroku
```

---

## External News API Integration

### APIs Used:
1. [GNews API](https://gnews.io/)
2. [NewsAPI](https://newsapi.org/)

### Example API Calls:
#### GNews (using `curl`):
```bash
curl -X GET "https://gnews.io/api/v4/top-headlines?token=YOUR_API_KEY&lang=en"
```
#### NewsAPI (using `curl`):
```bash
curl -X GET "https://newsapi.org/v2/top-headlines?apiKey=YOUR_API_KEY&country=us"
```

---

## Testing

Run the test suite using:
```bash
go test ./...
```
Tests are included for:
- JWT authentication and token renewal.
- API handlers for fetching and filtering news.
- Error handling scenarios for invalid input or tokens.

---

## Deployment

### Free Deployment Options:
1. **Heroku**:
   - Add a `Procfile` with:
     ```
     web: ./news-api-system
     ```
   - Deploy using Git:
     ```bash
     git push heroku master
     ```
2. **Render**:
   - Connect your GitHub repository to Render.
   - Set the build command to `go build` and start command to `./main`.
3. **Railway**:
   - Link Railway to your GitHub repository.
   - Deploy automatically after each commit.
4. **Fly.io**:
   - Initialize and deploy using:
     ```bash
     flyctl launch
     ```

---

## Future Improvements

1. **Database Integration**:
   - Replace in-memory storage with a relational or NoSQL database (e.g., PostgreSQL, MongoDB).
2. **Real-Time News Fetching**:
   - Direct integration with external news APIs for up-to-date information.
3. **Enhanced Security**:
   - Implement stricter rate-limiting and IP blocking.
4. **Continuous Integration**:
   - Add CI/CD pipelines to automate testing and deployments.

---

## License
This project is licensed under the MIT License. Feel free to use, modify, and distribute as needed.

---

## Contributions
Contributions are welcome! Please open an issue or submit a pull request to suggest improvements or fix bugs.

---

### Author
Developed by [felipeazsantos].
