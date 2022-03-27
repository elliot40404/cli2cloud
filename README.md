# CLI2CLOUD - pipe stdout to web

![GitHub release (latest by date)](https://img.shields.io/github/v/release/elliot40404/cli2cloud?style=flat-square)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/elliot40404/cli2cloud/node_ci?style=flat-square)
![GitHub Repo stars](https://img.shields.io/github/stars/elliot40404/cli2cloud?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/elliot40404/cli2cloud?style=flat-square)

```sh
      _ _ ____      _                 _
  ___| (_)___ \ ___| | ___  _   _  __| |
 / __| | | __) / __| |/ _ \| | | |/ _` |
| (__| | |/ __/ (__| | (_) | |_| | (_| |
 \___|_|_|_____\___|_|\___/ \__,_|\__,_|

cli2cloud - A command line utility to pipe stdout to a web interface
Version: 0.1.0
```

- [Table of Contents](#fr---find-and-replace)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Examples](#examples)

## Web interface

[https://cli2cloud.herokuapp.com/](https://cli2cloud.herokuapp.com/)

## Installation

```sh
git clone https://github.com/elliot40404/cli2cloud
cd cli2cloud/cli
go build
```

## Usage

#### 1. Download/Build the binary and add it to your PATH

#### 2. When using for the first time run `c2c`

This should generate a 8 character long key.

```sh
$ c2c
No key found
Creating new key...
Created key: 4MQrxK9p
```

#### 3. Now if you run c2c again it should show the key along with the instructions

```sh
$ c2c
Cli-2-Cloud
A command line interface for piping output to web
Key: 4MQrxK9p
Web-UI: https://cli2cloud.herokuapp.com/4MQrxK9p
Version: 0.1.0
Usage:
    $ command | c2c - pipe stdout to web
    $ c2c           - print this help menu
                      also generates a new key if none exists
```

#### 4. Now you can use the key to pipe the output to the web

Open up a browser and go to `https://cli2cloud.herokuapp.com/#/cli?id=<key>`

### 5. Now you can pipe the output to the web

```sh
echo "Hello World" | c2c
```

## License

MIT
