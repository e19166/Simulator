# Simulator

A Go-based event management simulator with PostgreSQL integration. This application provides a RESTful API to create and manage event records.

## Features

- **Database Integration**: Connects to a PostgreSQL database using Go's `database/sql` package.
- **Event Management**: Allows the creation of event records via a RESTful API endpoint.
- **Routing**: Uses Gorilla Mux for handling HTTP routes.

## Getting Started

### Prerequisites

- Go 1.18 or later
- PostgreSQL
- Gorilla Mux (`github.com/gorilla/mux`)
- PostgreSQL Driver (`github.com/lib/pq`)

### Installation

1. **Clone the Repository:**

```sh
   git clone https://github.com/yourusername/Simulator.git
   cd Simulator
```

2.**Install Dependencies:**

```sh
  go mod tidy
```

3. **Set Up the Database:**

 * Ensure PostgreSQL is installed and running.
 * Create a database and a **projects** table with the appropriate schema.

4. **Configure Database Connection:**

 * Update the connection string in **db/init.go** to match your PostgreSQL setup.

5. **Run the Application:**

```sh
  go run main.go
```
  The server will start on port 8080.

### API Endpoints

* **POST /projects**

  Create a new event.

  **Request Body Example:**

  ```json
  {
  "Metadata": {
    "Name": "Event Alpha",
    "Labels": {
      "Environment": "production",
      "Type": "web"
    },
    "Annotations": {
      "MonitoringType": "full"
    },
    "DeletionTimestamp": "2024-01-01T00:00:00Z",
    "Reason": "Initial creation",
    "Message": "This is a test event"
  },
  "ApiVersion": "v1",
  "Kind": "project",
  "InvolvedObject": {
    "Name": "Object1",
    "Uuid": "123e4567-e89b-12d3-a456-426614174000",
    "Version": "v1"
  },
  "Action": "create",
  "EventTime": "2024-01-01T00:00:00Z",
  "Source": {
    "Component": "component1",
    "Host": "host1"
  },
  "Count": 1,
  "Outcome": "success",
  "CurrentStatus": "active",
  "CorrelationID": "correlation-123",
  "UserIdpId": "user-123",
  "OrgUuid": "550e8400-e29b-41d4-a716-446655440000",
  "Series": {
    "FirstTimestamp": "2024-01-01T00:00:00Z",
    "LastTimestamp": "2024-01-01T00:00:00Z",
    "State": "running"
  }}
  ```
  

