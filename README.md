# trace-demo

`trace-demo` is a simple demo to generate traces between different operations, using OpenTelemetry to trace the requests, and Jaeger to export and visualize the traces. 

This is a side project to generate data for developing advanced topology among different modules in [flashcat](https://flashcat.com).

## Usage

1. Start Jaeger, and the http server.

```bash
docker-compose up -d
go run .
```

2. Open Jaeger UI and check the generated traces.

```
http://localhost:16686
```