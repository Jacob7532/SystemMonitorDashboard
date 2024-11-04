# System Monitor Dashboard

This project is designed to monitor system metrics such as CPU usage, memory usage, disk space, and GPU temperature in real-time. Built using Go for the backend and a simple HTML frontend, this application provides a clear view of system health and performance. The modular structure allows for easy future expansion into a more sophisticated frontend using Next.js.

## Setup Instructions

### Prerequisites

- Docker
- Go

### Running the Application

1. Clone the repository:
   ```
   git clone https://github.com/jacob7532/SystemMonitorDashboard.git
   cd system-monitor-dashboard
   ```

2. Build and start the Docker containers for the backend:
    ```
    cd backend
    docker build -t system-monitor-backend .
    docker run -d -p 8080:8080 --name system-monitor-backend system-monitor-backend
    ```

3. Serve the frontend using a simple HTTP server (Python):
    ```
    cd ../frontend
    python3 -m http.server 8000
    ```

4. Access the application:

    Frontend: http://localhost:8000

## Libraries and Tools Used

### Backend

    Go for the server-side application
    Gopsutil for gathering system metrics
    Gin framework for handling HTTP requests

### Frontend

    Basic HTML and CSS for initial interface
    Development and Deployment
    Docker for containerization and easy deployment
