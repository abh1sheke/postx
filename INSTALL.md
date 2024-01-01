# Installation

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
