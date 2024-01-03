# Zing 
Zing is a cross-platform, fast and easy-to-use HTTP client for the command-line.

## Features
### Perform HTTP requests
Zing allows you to make HTTP requests right from the command-line. All HTTP methods are supported including requests involving forms (`x-www-form-urlencoded`,  `multipart/form-data`).

### Proxied requests
Zing also allows you to route your requests through a proxy server.

## Usage

```
A cross-platform, fast and easy-to-use HTTP client for the command-line.

Usage:
  zing <url> [flags]

Flags:
  -d, --data stringArray        set http POST data
      --data-form stringArray   set http POST data as multipart/form-data (name=value)
  -j, --data-json string        set http POST data as JSON
      --data-text string        set http POST data as text/plain
  -F, --file stringArray        set MIME multipart MIME file (name=file)
  -H, --headers stringArray     set request headers (name=value)
  -h, --help                    help for zing
  -i, --include                 include request headers in output
  -m, --method string           http request method (default "get")
  -M, --multipart               send request data as multipart/form-data
  -o, --output string           specify output file
  -p, --proxy string            proxy url
  -t, --timeout int             set request timeout (in seconds) (default 10)
  -A, --user-agent string       specify User-Agent to send (default "zing/0.1.1")
  -v, --version                 version for zing
```

## Roadmap
- [ ] Read and perform requests (concurrently) from a user-created file.
- [ ] Perform FTP file uploads.
- [ ] Perform SMTP operations
- [ ] Allow users to benchmark backend server performance.
