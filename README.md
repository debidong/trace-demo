# trace-demo

`trace-demo` is a simple demo to generate traces between different operations, using OpenTelemetry to trace the requests, and Jaeger to export and visualize the traces. 

This is a side project to generate test data for developing advanced topology among different modules in [flashcat](https://flashcat.com).

## Usage

1. Start Jaeger, and the http server.

```bash
docker-compose up -d

go build
./trace-demo -n china -c config.yaml
./trace-demo -n america -c config.yaml
./trace-demo -n england -c config.yaml
```

2. Open Jaeger UI and check the generated traces.

```
http://localhost:16686
```

3. (Optional) Customize trace services & operations

Edit `config.yaml` to add/remove services & operations, and restart the http server.

A server & operation is defined as below:

```yaml
server:
  <server_name>:
    addr: <server_address>
    uri:
    - <operation_uri_1>
    - <operation_uri_2>
    - ...
```

## Reference

- [OpenTelemetry - Go Jaeger Exporter(TODO: deprecated, use OTLP collector instead)](https://pkg.go.dev/go.opentelemetry.io/otel/exporters/jaeger)
- [OpenTelemetry - instrumenting a Go application](https://opentelemetry.io/docs/languages/go/instrumentation/)
- [Jaeger - apis](https://www.jaegertracing.io/docs/1.66/apis/)