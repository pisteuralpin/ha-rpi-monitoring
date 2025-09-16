# Home Assistant Raspberry Pi Monitoring
A simple Go application to monitor Raspberry Pi devices and report their status to Home Assistant via MQTT.

## Features
- [x] CPU Temperature
- [ ] Fan status
- [ ] CPU usage
- [ ] Memory usage
- [ ] Disk usage
- [ ] Network status
- [ ] PCIe temperature
- [ ] Uptime
- [ ] WireGuard status


# Installation
## Docker
1. Clone the repository:
   ```bash
   git clone https://github.com/pisteuralpin/ha-rpi-monitoring.git
    cd ha-rpi-monitoring
    ```
2. Create a `.env` file based on the provided `.env.example` and configure it with your MQTT broker details and desired settings.
3. Build and run the Docker container:
    ```bash
    docker-compose up -d --build
    ```

# Configuration
| Parameter                  | Description                                      | Default Value        |
|----------------------------|--------------------------------------------------|----------------------|
| `MQTT_BROKER`              | Address of the MQTT broker                       | `localhost`          |
| `MQTT_PORT`                | Port of the MQTT broker                          | `1883`               |
| `MQTT_USER`                | Username for MQTT authentication                 | `your_username`      |
| `MQTT_PASSWORD`            | Password for MQTT authentication                 | `your_password`      |
| `MQTT_PREFIX`              | Prefix for MQTT topics                           | `homeassistant/`     |
| `CPU_TEMPERATURE_ENABLED`  | Enable CPU temperature monitoring                | `true`               |
| `CPU_ENTITY_NAME`          | Unique entity name for CPU temperature sensor    | `cpu`                |
| `CPU_TEMPERATURE_INTERVAL` | Interval for CPU temperature updates             | `10s`                |
