# postx
A blazingly-fast CURL-like tool to help you test RESTful API endpoints from command line.

## Usage

```
usage: postx [-h|--help] -m|--method "<value>" -u|--url "<value>" [-d|--data
             "<value>"] [-H|--headers "<value>" [-H|--headers "<value>" ...]]
             [-r|--repeat <integer>] [-l|--loop "<value>"]

             A CLI tool to help you test RESTful endpoints

Arguments:

  -h  --help     Print help information
  -m  --method   GET | POST; HTTP method
  -u  --url      URL of endpoint
  -d  --data     JSON; POST data
  -H  --headers  key:value; Set request headers
  -r  --repeat   number; Repeat the request n number of times
  -l  --loop     true | false; Perform n repitions forever (with a 1s timeout)
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
