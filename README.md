# xjftpd

a FTP server


## Usage
`xjftpd [port] [username] [password] [directory]`

Examples:
- `xjftpd 21 user pass /home/user/ftp`
- `xjftpd.exe 21 user pass C:\ftp`


## Features

- Only 1 user
- Read/write/delete files, directories
- No config files
- No environment variables
- Perfect for running as a service
- Suitable for scripting with sigterm


## Why
- Does not rely on local system users
- No complex permissions issues
- Does not require config files to override 1 home directory (vsftpd requires *multiple*)
- Literally no simple FTP servers for linux or windows in 2024 (??)
## Running as a Service

To run xjftpd as a service, create a systemd service file:

```sh
sudo nano /etc/systemd/system/xjftpd.service
```

```ini
[Unit]
Description=xjftpd FTP Server
After=network.target

[Service]
ExecStart=/usr/local/bin/xjftpd 21 user pass /home/user/ftpdir
Restart=always
Type=simple
User=www-data
Group=www-data
Restart=always
RestartSec=3
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=xjftpd

[Install]
WantedBy=multi-user.target
```

Enable and start the service:

```sh
sudo systemctl enable xjftpd
sudo systemctl start xjftpd
```


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.