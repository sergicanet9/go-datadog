# go-datadog

This project is a backend web application built with the **Go** programming language. The application exposes two simple RESTful endpoints and is instrumented with **Datadog Agent** to provide comprehensive monitoring.

---

## 🚀 Features

* **Go & HTTP**: A modern backend build with Go and its `net/http` standard library.
* **Docker & Docker Compose**: Containerization of the application and the Datadog Agent for easy deployment.
* **Datadog Agent**: Automatic instrumentation to integrate **APM, logs, and infrastructure metrics**.
* **Unit Tests**: The project includes unit tests for the API endpoints.
* **Makefile**: Automation of common tasks like building and running the project.

---

## 🛠️ Prerequisites

Make sure you have the following software installed on your system:

* **Go (1.21+)**
* **Docker & Docker Compose**
* **Git**

---

## 🏁 Getting Started

Follow these steps to get the project up and running on your local machine.

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/sergicanet9/go-datadog.git
    cd go-datadog
    ```

2.  **Configure Datadog:**
    Create a file named `.env` in the root of the project and add your Datadog API key.

    ```bash
    # Datadog API Key
    DATADOG_API_KEY="your_api_key_here"
    ```

3.  **Build and run the application:**
    The `Makefile` will handle building the Docker image, running `docker-compose`, and starting the Go server along with the Datadog Agent.

    ```bash
    make run
    ```

    The application will be available at `http://localhost:8080`.

---

## ✅ Running Tests

You can run the unit tests for the application to ensure that the endpoints are working as expected.

```bash
make test
```

---

## 📦 API Endpoints

| Method | Endpoint | Description |
| :----- | :------- | :---------- |
| `GET` | `/` | Returns a simple message to verify the server is running. |
| `GET` | `/report` | A simple route that simulates a long-running transaction. Great for testing the Datadog instrumentation. |

---

## 📈 Monitoring with Datadog

This project is fully instrumented with the Datadog Agent to send the following data:

* **APM**: Detailed transaction traces, latency, and error rates for the application.
* **Logs**: Application logs collected from standard output, automatically correlated with the APM traces.
* **Infrastructure Metrics**: Key system-level metrics (CPU, memory, disk I/O) from the Docker container and the host.

After running `make run` and making a few requests to the API endpoints, you'll be able to see the data in your Datadog account.

---

