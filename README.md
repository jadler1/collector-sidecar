# Graylog Sidecar

[![Build Status](https://travis-ci.org/Graylog2/collector-sidecar.svg?branch=master)](https://travis-ci.org/Graylog2/collector-sidecar) [![Go Report Card](https://goreportcard.com/badge/github.com/graylog2/collector-sidecar)](https://goreportcard.com/report/github.com/graylog2/collector-sidecar)

# WARNING - ALPHA VERSION

The master branch is tracking the upcoming 1.0 version of the sidecar. Please see the [0.x branch](https://github.com/Graylog2/collector-sidecar/tree/0.x) for the current released 0.x version.

**Required Graylog version:** 3.0 and later.

The Graylog Sidecar is a supervisor process for 3rd party log collectors like NXLog and filebeat. The Sidecar program is able to fetch and validate configuration files from a Graylog server for various log collectors. You can think of it like a centralized configuration and process management system for your log collectors.

## Documentation

Please check our official [documentation](http://docs.graylog.org/en/latest/pages/collector_sidecar.html) for more information.
Especially the [Step-by-Step](http://docs.graylog.org/en/2.4/pages/collector_sidecar.html#step-by-step-guide) guide to get the first setup running.

## Installation

| Sidecar version  | Graylog server version |
| ---------------- | ---------------------- |
| 0.0.9            | 2.1.x                  |
| 0.1.x            | 2.2.x, 2.3.x, 2.4.x    |
| 1.x.x            | 3.0.x                  |

[Download a package](https://github.com/Graylog2/collector-sidecar/releases) and install it on the target system.


### Beats backend
**Ubuntu**

The Beats binaries (Filebeat and Winlogeventbeat) are included in the Sidecar package. So installation is just one command.

```
  $ sudo dpkg -i graylog-sidecar_1.0.0-1_amd64.deb
```

Edit `/etc/graylog/sidecar/sidecar.yml`, you should set at least the correct URL to your Graylog server and proper tags.
The tags are used to define which configurations the host should receive.

Create a system service and start it

```
  $ sudo graylog-sidecar -service install

  [Ubuntu 14.04 with Upstart]
  $ sudo start graylog-sidecar

  [Ubuntu 16.04 with Systemd]
  $ sudo systemctl start graylog-sidecar
```

**CentOS**

```
  $ sudo rpm -i graylog-sidecar-1.0.0-1.x86_64.rpm
```

Activate the Sidecar as a system service

```
  $ sudo graylog-sidecar -service install
  $ sudo systemctl start graylog-sidecar
```

**Windows**

_**The Windows installation path changed to `C:\Program Files` with version 0.0.9, please stop and uninstall former installations before doing the update**_

```
  $ graylog_sidecar_installer.exe
```

It's also possible to run the installer in silent mode with

```
  $ graylog_sidecar_installer.exe /S -SERVERURL=http://10.0.2.2:9000/api -TAGS="windows,iis"
```

Edit `C:\Program Files\graylog\sidecar\sidecar.yml`.

```
  $ C:\Program Files\Graylog\sidecar\graylog-sidecar.exe -service install
  $ C:\Program Files\Graylog\sidecar\graylog-sidecar.exe -service start
```

All installer options:

| Parameter             | Description                           | Default                   |
|-----------------------|---------------------------------------|---------------------------|
| `-SERVERURL`          | URL to the Graylog API                | http://127.0.0.1:9000/api |
| `-NODENAME`           | Name of the instance                  | graylog-sidecar           |
| `-APITOKEN`           | The server API token                  |                           |
| `-UPDATE_INTERVAL`    | Seconds between configuration updates | 10                        |
| `-TLS_SKIP_VERIFY`    | Ignore self-signed API certificates   | false                     |
| `-SEND_STATUS`        | Send host metrics back to Graylog     | true                      |

### NXLog backend

**Ubuntu**

Install the NXLog package from the offical download [page](https://nxlog.org/products/nxlog-community-edition/download)

```
  $ sudo /etc/init.d/nxlog stop
  $ sudo update-rc.d -f nxlog remove
  $ sudo gpasswd -a nxlog adm
 
  $ sudo dpkg -i graylog-sidecar_1.0.0-1_amd64.deb
  $ sudo chown -R nxlog.nxlog /var/spool/graylog-sidecar/nxlog
```

Edit `/etc/graylog/sidecar/sidecar.yml`accordingly.

```
  $ sudo graylog-sidecar -service install

  [Ubuntu 14.04 with Upstart]
  $ sudo start graylog-sidecar

  [Ubuntu 16.04 with Systemd]
  $ sudo systemctl start graylog-sidecar
```

**CentOS**

```
  $ sudo service nxlog stop
  $ sudo chkconfig --del nxlog
  $ sudo gpasswd -a nxlog root
  $ sudo chown -R nxlog.nxlog /var/spool/graylog-sidecar/nxlog

  $ sudo rpm -i graylog-sidecar-1.0.0-1.x86_64.rpm
```

Activate the Sidecar as a system service

```
  $ sudo graylog-sidecar -service install
  $ sudo systemctl start graylog-sidecar
```

**Windows**

_**The Windows installation path changed to `C:\Program Files` with version 0.0.9, please stop and uninstall former installations before doing the update**_

Also notice that the NXLog file input is currently not able to do a SavePos for file tailing, this will be fixed in a future version.

Install the NXLog package from the offical download [page](https://nxlog.org/products/nxlog-community-edition/download) and deactivate the
system service. We just need the binaries installed on that host.

```
  $ C:\Program Files (x86)\nxlog\nxlog -u

  $ graylog_sidecar_installer.exe
```

Edit `C:\Program Files\graylog\sidecar\sidecar.yml`, you should set at least the correct URL to your Graylog server and proper tags.

```
  $ C:\Program Files\graylog\sidecar\graylog-sidecar.exe -service install
  $ C:\Program Files\graylog\sidecar\graylog-sidecar.exe -service start
```

## Uninstall on Windows

```
  $ C:\Program Files\graylog\sidecar\graylog-sidecar.exe -service stop
  $ C:\Program Files\graylog\sidecar\graylog-sidecar.exe -service uninstall
```

## Debugging

Run the Sidecar in foreground mode for debugging purposes. Simply call it like this and look out for error messages:

```
  $ graylog-sidecar -debug -c /etc/graylog/sidecar/sidecar.yml
```

## Configuration

There are a couple of configuration settings for the Sidecar:

| Parameter           | Description                                                                                                                           |
|---------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| `server_url`        | URL to the Graylog API, e.g. `http://127.0.0.1:9000/api/`                                                                             |
| `server_api_token`  | The API token to use for authentication against the Graylog API                                                                       |
| `update_interval`   | The interval in seconds the sidecar will fetch new configurations from the Graylog server                                             |
| `tls_skip_verify`   | Ignore errors when the REST API was started with a self-signed certificate                                                            |
| `send_status`       | Send the status of each backend back to Graylog and display it on the status page for the host                                        |
| `list_log_files`    | Send a directory listing to Graylog and display it on the host status page. This can also be a list of directories                    |
| `node_name`         | Name of the Sidecar instance, will also show up in the web interface                                                                  |
| `node_id`           | Unique ID (UUID) of the instance. This can be an ID string or a path to an ID file                                                    |
| `log_path`          | A path to a directory where the Sidecar can store the output of each running collector backend                                        |
| `log_rotation_time` | Rotate the stdout and stderr logs of each collector after X seconds                                                                   |
| `log_max_age`       | Delete rotated log files older than Y seconds                                                                                         |

## Compile

  * Clone the repository into your `$GOPATH` under `src/github.com/Graylog2/collector-sidecar`
  * run `make` to install the dependencies and build the binary for the local platform
  * run `make help` to see more targets

