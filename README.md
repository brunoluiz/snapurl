# snapurl
⚡️ Easy website snapshots using Terminal, HTTP or GRPC

## Intro
It uses Chromium Headless to get website screenshots. It can be used as a module or as a tool. There is a GRPC Server (with HTTP Gateway) and a cli on `./cmd` folder.

## CLI

### Installation

Using `go get -u github.com/brunoluiz/snapurl/cmd/snapurl` to install the cli tool. It will be available as `snapurl` in the terminal.

### Usage

A simple `snapurl https://google.co.uk` will make it work. There are extra options which can be used to customise its behaviour:

Options:
```
--out-dir value      output directory (default: ".")
--out value          output path (folder + filename) -- overrides --out-dir
--wait-period value  wait period in seconds to render the page (default: 5)
```

The default filename is `screenshot-timestamp.png`
