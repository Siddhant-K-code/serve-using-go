# 🚀 Go-FileServer: A Lightweight HTTP File Server

Welcome to Go-FileServer, a nifty and efficient HTTP file server written in Go! It is tailored to be both powerful and user-friendly. Whether you're a developer, sysadmin, or just someone who loves toying around with servers, you'll find Go-FileServer both practical and fun to use! 🌟

## Features

- **🌐 Serve Files Over HTTP**: Quickly serve files from any directory.
- **🔌 Optional Public Access**: Choose between local or public access.
- **🔐 CORS Support**: Easily configure CORS settings for cross-origin resource sharing.
- **🌍 Auto Open in Browser**: Launch the server and view in a browser with a single command.
- **🔍 Flexible Port Configuration**: Choose your desired port.
- **👨‍💻 Cross-Platform Support**: Works seamlessly on Linux, Windows, and macOS.

## Getting Started

Just start a Gitpod workspace by clicking the button below:

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/Siddhant-K-code/serve-using-go)

### Usage

Run the server with default settings:

```bash
go run main.go
```

Advanced Usage:

```bash
go run main.go -port=8080 -public=true -cors-allow="https://<your-domain-url>.com" -open
```

Flags:

- `-port`: Set the port number (default: 8000).
- `-public`: Listen on all interfaces (default: listens only on localhost).
- `-cors-allow`: Set CORS allowed origins.
- `-open`: Automatically open in the default web browser.

## License

Distributed under the MIT License. See `LICENSE` for more information.
