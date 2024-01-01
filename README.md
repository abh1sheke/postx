# Zing 
Imagine `cURL`, but on steroids. zing is just that. A blisteringly-fast, concurrent and easy to use CLI tool that greatly expedites your development by allowing for speedy data transfers and robust endpoint testing capabilities.

## Features
### Perform HTTP requests
Just like cURL, zing allows you to perform the basic HTTP requests (GET, POST, HEAD, PUT & DELETE), in addition to requests involving form data.

### Make requests concurrently
zing helps you test your API endpoints by helping you perform N requests simultaneously. You can also perform said concurrent requests in loop with a set timeout between each iteraton to further stress-test your endpoints.

## Usage

```
A fast and feature-rich alternative to cURL.

Usage:
  zing [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Perform a DELETE request
  get         Perform a GET request
  head        Perform a HEAD request
  help        Help about any command
  post        Perform a POST request
  put         Perform a PUT request

Flags:
  -h, --help            help for zing
  -i, --include         include request headers in output
  -o, --output string   specify output file
  -t, --time            show time taken to make requests
```

## Installation
### Through Go toolchain
If you have the Go toolchain installed on your device, zing can be installed by running the following command in your terminal:
```
go install github.com/abh1sheke/zing
```

> Make sure the Go binary is in your PATH before you can run zing

<br />

### Building from source
Follow these instructions to build zing from source:
  * Install the Go toolchain from [here](https://go.dev/doc/install)
  * Clone this repository
  ```
  git clone https://github.com/abh1sheke/zing && cd zing
  ```
  * Build the executable by running:
  ```
  go build -o zing
  ```
  * Add zing to PATH by running:
  ```
  # On macOS and Linux:
  export PATH="$PATH:/path/to/zing"

  # On Windows
  set PATH=%PATH%;C:\path\to\zing\
  ```
