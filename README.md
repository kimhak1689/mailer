# Go Email Sender

This is a simple Go application that provides an API endpoint for sending emails using SMTP. It includes basic authentication for added security. The application uses the Gin framework for routing and Gomail for sending emails. Environment variables are used to configure SMTP settings, authentication, and email content.

## Getting Started

Follow the steps below to set up and run the application locally.

### Prerequisites

- Go installed on your machine
- Docker and Docker Compose (optional, for running in a container)

### Running Locally

1. Clone the repository to your local machine:

```bash
git clone <repository_url>
```

2. Navigate to the project directory:
```
cd go-email-sender
```
3. Create a .env file in the root directory of the project and add the following environment variables:
```
SMTP_USERNAME=
SMTP_PASSWORD=
SMTP_HOST=
SMTP_PORT=
SMTP_FROM_ADDRESS=
BASIC_AUTH_USERNAME=
BASIC_AUTH_PASSWORD=
SMTP_FROM_NAME=
```
4. Build and run the application:
```
go run main.go
```
The application will start on port 8080.

## Running with Docker
1. Create a .env file as described in the previous section.

2. Use Docker Compose to build and run the application:

```
docker-compose up --build
```

## Configuration
The application can be configured using environment variables. Here is a list of available variables:
```
SMTP_USERNAME: SMTP username for authentication
SMTP_PASSWORD: SMTP password for authentication
SMTP_HOST: SMTP server hostname
SMTP_PORT: SMTP server port
SMTP_FROM_ADDRESS: Email address used in the "From" field
BASIC_AUTH_USERNAME: Username for basic authentication
BASIC_AUTH_PASSWORD: Password for basic authentication
SMTP_FROM_NAME: Name associated with the "From" address
```
## API Usage
### Endpoint
```
POST /send_email
```
Request Body
to_address: Recipient's email address
subject: Email subject
body: Email content (HTML supported)
### Example
```
curl -X POST http://localhost:8080/send_email \
-u username:password \
-F "to_address=xxx@gmail.com" \
-F "subject=Test Subject" \
-F "body=This is a test email body."
```
