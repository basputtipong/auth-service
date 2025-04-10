# Auth-Service

## Build & Run

### Build Docker Container
```bash
make docker-build
```

### Start Docker Container
```bash
make docker-up
```

---

## Local Development

- Run the service locally:
```bash
make run
```

- The service will be available at:
  http://localhost:1300

---

## API Details

### POST /login – Login with userId and passcode

**Request Body:**
```json
{
  "userId": "000018b0e1a211ef95a30242ac180002",
  "passcode": "123456"
}
```

**Response:**
```json
{
  "authToken": "some.jwt.token"
}
```

---

### POST /verify – Verify User Passcode

**Request Header:**
```
Authorization: Bearer {authToken}
```

**Request Body:**
```json
{
  "passcode": "123456"
}
```

**Response:**
```json
{
  "name": "John Doe",
  "greetingMsg": "hello"
}
```

---

### GET /banner – Get User Banner

**Request Header:**
```
Authorization: Bearer {authToken}
```

**Request Body:**
```json
{}
```

**Response:**
```json
{
  "bannerId": "000018cfe1a211ef95a30242ac180002",
  "title": "Want some money?",
  "description": "You can start applying",
  "image": "https://dummyimage.com/54x54/999/fff"
}
```

---

### GET /health – Health Check

**Response:**
```json
{
  "status": "healthy"
}
```

---

## Unit Testing

### Run Unit Tests
```bash
make test-service
```

### Generate Mocks with Mockery

1. Install mockery:
```bash
make mock-install
```

2. Generate mock files:
```bash
mockery --all --dir=internal/core/port --output=internal/core/port/mocks --outpkg=mocks
```

---

See `Makefile` for other useful commands.
