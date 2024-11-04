# System Monitor Dashboard
Still a work in progress, Docker is not building correctly at the moment

This project is designed to monitor system metrics such as CPU usage, memory usage, and disk space in real-time. Built using Go for the backend and a simple HTML frontend, this application provides a clear view of system health and performance. The modular structure allows for easy future expansion into a more sophisticated frontend using Next.js.

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
    docker-compose up --build
    ```
    
4. Access the application:

    Frontend: http://localhost:8000
    Backend API: http://localhost:8080/api/stats

## Libraries and Tools Used

### Backend

    Go for the server-side application
    Gopsutil for gathering system metrics
    Gin framework for handling HTTP requests

### Frontend

    Basic HTML and CSS for initial interface
    Development and Deployment
    Docker for containerization and easy deployment
