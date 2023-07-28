# postx
Imagine `cURL`, but on steroids. postx is just that. A blisteringly-fast, concurrent and easy to use CLI tool that greatly expedites your development by allowing for speedy data transfers and robust endpoint testing capabilities.

## Features
### Perform HTTP requests
Just like cURL, postx allows you to perform the basic HTTP requests (GET, POST, HEAD, PUT & DELETE), in addition to requests involving form data.

### Make requests concurrently
Postx helps you test your API endpoints by helping you perform N requests simultaneously. You can also perform said concurrent requests in loop with a set timeout between each iteraton to further stress-test your endpoints.

## Usage

```
postx <Command> [-h|--help] [-p|--parallel <integer>] [-l|--loop]
             [-i|--include] [-o|--output "<value>"]

             A CLI tool for transferring data over HTTP.

Commands:

  get     Perform a GET request
  head    Perform a HEAD request
  post    Perform a POST request
  put     Perform a PUT request
  delete  Perform a DELETE request
  form    Submit a HTML form

Arguments:

  -h  --help      Print help information
  -p  --parallel  number; Perform n requests in parallel
  -l  --loop      Loop request forever (with a 1s timeout)
  -i  --include   Include response headers in the output
  -o  --output    Specify output file
```

## Installation
### Through Go toolchain
If you have the Go toolchain installed on your device, postx can be installed by running the following command in your terminal:
```
go install github.com/abh1sheke/postx
```

> Make sure the Go binary is in your PATH before you can run postx

<br />

### Building from source
Follow these instructions to build postx from source:
  * Install the Go toolchain from [here](https://go.dev/doc/install)
  * Clone this repository
  ```
  git clone https://github.com/abh1sheke/postx && cd postx
  ```
  * Build the executable by running:
  ```
  go build -o postx
  ```
  * Add postx to PATH by running:
  ```
  # On macOS and Linux:
  export PATH="$PATH:/path/to/postx"

  # On Windows
  set PATH=%PATH%;C:\path\to\postx\
  ```
