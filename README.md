# Home Assistant Raspberry Pi Monitoring
A simple Go application to monitor Raspberry Pi devices and report their status to Home Assistant via MQTT.

## Features
- [x] CPU Temperature
- [x] Power supply
- [ ] Fan status
- [ ] PCIe temperature
- [ ] Uptime
- [ ] WireGuard status


## Configuration
Create a .env file in the root directory with the following variables:
| Variable      | Description                                        | Example                                |
| ------------- | -------------------------------------------------- | -------------------------------------- |
| MQTT_BROKER   | The address of your MQTT broker                    | localhost                              |
| MQTT_PORT     | The port of your MQTT broker                       | 1883                                   |
| MQTT_USERNAME | The username for your MQTT broker                  | user                                   |
| MQTT_PASSWORD | The password for your MQTT broker                  | password                               |
| LOG_LEVEL     | The logging level (e.g., debug, info, warn, error) | info                                   |
| CONFIG_PATH   | Path to the configuration file                     | ~/.config/ha-rpi-monitoring/config.yml |

When you first launch the application, a `.config` folder will be created in your home directory (`~/.config/ha-rpi-monitoring/`), containing a default `config.yaml` file. You can edit this file to customize the application's behavior.


## Building and Running
Make sure you have Go installed on your system. Then, you can build and run the application using the following commands:
```bash
go build -o build/main .
./build/main
```


### Using systemd to Manage the Service
You can use systemd to manage the Raspberry Pi Monitoring application as a service. More details can be found in the [Using systemd to run the Raspberry Pi Monitoring Service](docs/using_systemd.md) guide.

## Thanks
Thanks to jfikar to his amazing work on [estimating the power consumption of a Raspberry Pi 5](https://github.com/jfikar/RPi5-power)