# Caching Proxy Server

A lightweight HTTP proxy server written in Go that caches GET requests from a specified origin server.

## Features

1. Caches GET requests for improved performance
1. Forwards all other HTTP methods directly to the origin
1. Configurable port and origin server
1. Thread-safe caching mechanism

## Prerequisites

Before running this application, make sure you have the following installed:

- Go (latest version recommended)

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/PranayBajracharya/caching-proxy.git
   cd caching-proxy
   ```

1. Run the application:

   ```
   go run . --origin=https://example.com --port=3000
   ```

   This will start the server on port 3000 and proxy requests to https://example.com

## Flags

<table align="center" style="justify-content: center;align-items: center;display: flex;">
  <tr>
    <th>Flag</th>
    <th>Example</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>--origin</td>
    <td>--origin=https://example.com</td>
    <td>The base URL of the origin server to proxy requests from</td>
  </tr>
  <tr>
    <td>--port</td>
    <td>--port=3000</td>
    <td>The port number to run the proxy server on (Default: 3000)</td>
  </tr>
</table>

## How It Works

1. When a GET request is received, the server first checks its cache
1. If the response is cached, it returns the cached version
1. If not cached, it forwards the request to the origin server and caches the response
1. All non-GET requests are forwarded directly to the origin server
