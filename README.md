# Zing 
Zing is a cross-platform, fast and easy-to-use HTTP client for the command-line.

## Features
### Perform HTTP requests
Zing allows you to make HTTP requests right from the command-line. All HTTP methods are supported including requests involving forms (`x-www-form-urlencoded`,  `multipart/form-data`).

### Proxied requests
Zing also allows you to route your requests through a proxy server.

## Roadmap
- [ ] Read and perform requests (concurrently) from a user-created file.
- [ ] Perform FTP file uploads.
- [ ] Perform SMTP operations
- [ ] Allow users to benchmark backend server performance.

## Usage

```
A fast and easy-to-use HTTP client for the command-line.

Usage:
  zing [flags]

Flags:
  -d, --data stringArray      set http POST data (name=value)
  -F, --file stringArray      set MIME multipart MIME file (name=data)
  -H, --headers stringArray   set request headers (name=value)
  -h, --help                  help for zing
  -i, --include               include request headers in output
  -m, --method string         http request method (default "get")
  -M, --multipart             send request data as multipart/form-data
  -o, --output string         specify output file
  -p, --proxy string          proxy url
  -t, --timeout int           set request timeout (in seconds) (default 10)
  -u, --url string            endpoint url
  -v, --version               version for zing
```
