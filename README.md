# postx
A blazingly-fast, multi-threaded, cURL-like tool for transferring data over HTTP protocol and for testing/debugging API endpoints.

## Usage

```
usage: postx <Command> [-h|--help] [-p|--parallel <integer>] [-l|--loop]
             [-o|--output "<value>"]

             A CLI tool to help you test RESTful endpoints

Commands:

  get     Perform a GET request
  head    Perform a HEAD request
  post    Perform a POST request
  put     Perform a PUT request
  delete  Perform a DELETE request

Arguments:

  -h  --help      Print help information
  -p  --parallel  number; Perform n requests in parallel
  -l  --loop      Loop request forever (with a 1s timeout)
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
