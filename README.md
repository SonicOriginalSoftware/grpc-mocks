# grpc-testing

Testing utilities for gRPC services. Provides mock implementations and test
helpers for common gRPC infrastructure components.

## Overview

This library contains testing utilities that facilitate unit and integration
testing of gRPC services. Mock implementations allow controlled testing of
observability and infrastructure components without external dependencies.

## Installation

```bash
go get git.sonicoriginal.software/grpc-testing
```

## Packages

### mocks

Mock implementations of OpenTelemetry metrics components. Provides inspectable
mock meters with controllable error injection for testing code that uses
OpenTelemetry metrics.

Mock meters track recorded values and allow tests to verify that metrics are
being recorded correctly. Supports error injection to test error handling paths
in metric recording code.

## Usage

Import the mocks package in your test files. Use the mock implementations in
place of real OpenTelemetry components to verify metric recording behavior and
test error handling.

## Requirements

- Go 1.21 or later
- OpenTelemetry libraries

## License

Apache License 2.0
