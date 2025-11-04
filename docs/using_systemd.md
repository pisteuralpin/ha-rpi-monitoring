# Using systemd to run the Raspberry Pi Monitoring Service
To run the Raspberry Pi Monitoring application as a background service on your Raspberry Pi, you can use `systemd`. This allows the application to start automatically on boot and restart if it crashes.

## 1. Create a systemd Service File
1. Create a new service file for the Raspberry Pi Monitoring application. You can do this by creating a file named `ha-rpi-monitoring.service` in the `~/.config/systemd/user` directory:

   ```bash
   sudo nano /etc/systemd/system/ha-rpi-monitoring.service
   ```
2. Add the following content to the `ha-rpi-monitoring.service` file. Make sure to adjust the `WorkingDirectory` and `ExecStart` paths to match your setup:
    ```ini
    [Unit]
    Description=Raspberry Pi Monitoring via MQTT Service
    After=network.target

    [Service]
    Type=simple
    WorkingDirectory=/home/<user>/projects/ha-rpi-monitoring
    ExecStart=/usr/local/bin/start_ha_rpi_monitoring.sh
    Restart=always
    RestartSec=5
    Environment="PATH=/usr/local/bin:/usr/bin:/usr/local/go/bin"

    [Install]
    WantedBy=default.target
    ```

## 2. Create a Start Script
1. Create a shell script named `start_ha_rpi_monitoring.sh` in the `/usr/local/bin/` directory to set the environment variables and start the application:
    ```bash
    #!/bin/bash
    cd /home/<user>/projects/ha-rpi-monitoring

    mkdir -p build

    if [ ! -f build/main ] || [ "$(find . -name '*.go' -newer build/main -print -quit)" ]; then
        echo "Rebuilding Go binary..."
        /usr/local/go/bin/go build -o build/main .
    fi

    exec /home/<user>/projects/ha-rpi-monitoring/build/main
    ```
2. Make the script executable:
    ```bash
    sudo chmod +x /usr/local/bin/start_ha_rpi_monitoring.sh
    ```

## 3. Enable and Start the Service
1. Reload the systemd manager configuration to recognize the new service:
   ```bash
   systemctl --user daemon-reload
   ```
2. Enable the service to start on boot:
   ```bash
   systemctl --user enable ha-rpi-monitoring.service
   ```
3. Start the service immediately:
   ```bash
   systemctl --user start ha-rpi-monitoring.service
   ```
4. Check the status of the service to ensure it is running:
   ```bash
   systemctl --user status ha-rpi-monitoring.service
   ```