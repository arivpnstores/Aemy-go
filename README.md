# Go WhatsApp Bot

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://go.dev/)

A powerful and extensible WhatsApp bot written in Go, designed for easy interaction and customization.

## Features

- **WhatsApp Client**: Seamlessly connect and interact with the WhatsApp platform.
- **Message Handling**: Robust system for processing incoming and outgoing messages.
- **Customizable Configuration**: Easily configure the bot's behavior through simple settings.
- **Utility Suite**: Includes helpers for serialization, logging, and other common tasks.

## Prerequisites

Before you begin, ensure you have Go installed on your system.

### Windows (Recommended: Latest)

1.  **Download**: Get the latest Go installer from the [official Go website](https://go.dev/dl/).
2.  **Install**: Run the installer and follow the on-screen instructions.
3.  **Verify**: Open a new terminal (Command Prompt or PowerShell) and run:
    ```cmd
    go version
    ```

### Linux (Recommended: 1.22.0)

1.  **Download the Archive**:
    ```bash
    wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
    ```
2.  **Extract the Archive**:
    ```bash
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
    ```
3.  **Update PATH**: Add the following line to your `~/.profile` or `~/.bashrc`:
    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```
    Then, reload your shell profile: `source ~/.profile`.
4.  **Verify Installation**:
    ```bash
    go version
    ```

## Getting Started

1.  **Clone the Repository**:

    ```bash
    git clone https://github.com/arivpnstores/Aemy-go.git
    cd Aemy-go
    ```

2.  **Install Dependencies**:

    ```bash
    go get
    ```

3.  **Run the Application**:
    ```bash
    go run main.go
    ```

## Run as a Service (Linux - systemd)

You can run this bot automatically on system boot using systemd.

Step-by-Step

1. Create the service file:

  ```bash
sudo nano /etc/systemd/system/aemy.service
```

2. Paste the following configuration (adjust the path if needed):
 
 ```bash
[Unit]
Description=Aemy Go WhatsApp Bot
After=network.target

[Service]
Type=simple
WorkingDirectory=/root/Aemy-go
ExecStart=/usr/local/go/bin/go run main.go
Restart=on-failure
User=root

[Install]
WantedBy=multi-user.target
  ```

3. Reload and start the service:

  ```bash
sudo systemctl daemon-reexec
sudo systemctl daemon-reload
sudo systemctl enable aemy.service
sudo systemctl start aemy.service
```

4. Check the status:

  ```bash
systemctl status aemy.service
```

5. View real-time logs:

  ```bash
journalctl -u aemy.service -f
```


> ⚠️ Make sure the WorkingDirectory and ExecStart paths are correct for your system.
You can find the correct Go binary path using:


  ```bash
which go
```

## Troubleshooting

### Error: `go-sqlite3 requires cgo`

If you encounter the following error, it means the `go-sqlite3` package requires a C compiler, which may be disabled in your Go environment.

```
[Client ERROR] DB error: failed to upgrade database: failed to check if foreign keys are enabled: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
```

#### Linux Solution

Install the necessary build tools:

```bash
sudo apt update
sudo apt install build-essential gcc
```

#### Windows Solution

You need to install a GCC toolchain, such as [MinGW-w64](https://www.mingw-w64.org/). After installation, ensure the MinGW `bin` directory is added to your system's `PATH`.

## Author

- **[Seaavey]** - _Initial work_ - [seaavey](https://github.com/seaavey)

See also the list of [contributors](https://github.com/seaavey/Aemy-go/contributors) who participated in this project.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any feedback or find a bug.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
