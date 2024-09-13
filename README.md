![Yae Miko](yae-miko.jpg)  

# Yae Miko

A simple API Gateway built with round-robin and least connection load balancing algorithms. This gateway includes Prometheus for monitoring and Grafana for visualization. 

## Features

- **Round-robin load balancing**: Distributes requests evenly across backend services.
- **Least connection load balancing**: Routes requests to the backend service with the fewest active connections.
- **Prometheus integration**: Exposes metrics for monitoring.
- **Grafana integration**: Visualizes Prometheus metrics.
- **Health checks**: Regular health checks for backend services.

## Stack

- **API Gateway**: Custom gateway written in [Golang].
- **Load Balancer Algorithms**: Round-robin and least connection.
- **Monitoring**: 
  - **Prometheus**: Metrics collection from the API gateway.
  - **Grafana**: Visualization of metrics.
- **Containerization**: Docker and Docker Compose.

## Prerequisites

- **Docker**: Make sure you have Docker installed. [Install Docker](https://docs.docker.com/get-docker/).
- **Docker Compose**: Ensure Docker Compose is installed. [Install Docker Compose](https://docs.docker.com/compose/install/).

## Project Structure

```bash
.
├── docker-compose.yml          # Docker Compose configuration file
├── etc/prometheus.yml          # Prometheus configuration\
└── README.md                   # Project documentation
```

## Setup and Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/zakirkun/yae-miko.git
   cd yae-miko
   ```

2. **Update configuration files:**

   - Adjust any settings inside `etc/prometheus.yml` if needed.

3. **Start services using Docker Compose:**

   ```bash
   docker-compose up --build
   ```

   This will start:
   - The API Gateway with round-robin and least connection load balancing
   - Prometheus for monitoring the gateway
   - Grafana for visualizing metrics

4. **Access services:**
   - API Gateway: `http://localhost:3333`
   - Prometheus: `http://localhost:9090`
   - Grafana: `http://localhost:3000`

## Load Balancing Algorithms

- **Round-robin**: Distributes incoming requests in a circular order to the available services.
- **Least connection**: Routes the request to the server with the least number of active connections.

## Monitoring

### Prometheus

Prometheus is used to scrape metrics from the API gateway. Metrics are available at the `/metrics` endpoint of the gateway.

- **Prometheus URL**: `http://localhost:9090`
- Default scrape interval: 15s
- Configuration file: `etc/prometheus.yml`

### Grafana

Grafana is set up to visualize Prometheus metrics. Pre-configured dashboards can be customized according to your needs.

- **Grafana URL**: `http://localhost:3000`
- Default credentials: `admin/admin`
- You can import dashboards using the Grafana UI or add custom ones.

## Configuration

### Docker Compose Configuration

The `docker-compose.yml` defines three services:

- **yae_miko_cluster**: The main gateway application.
- **prometheus**: Prometheus service for collecting metrics.
- **grafana**: Grafana service for visualizing the metrics.

### API Gateway Configuration

- Configure the load balancing algorithms (round-robin, least connection) in the gateway's configuration file.

## Example Request

You can test the API gateway with a simple `curl` command:

```bash
curl -X GET http://localhost:3333/
```

## Dashboards

Grafana dashboards are set up to visualize:

- Request rates
- Latency per backend
- Active connections (least connection)
- Gateway response codes

## Contributing

Feel free to submit issues or contribute by creating pull requests. Contributions are always welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.