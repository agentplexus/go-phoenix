# Release Notes v0.1.0

**Release Date:** January 4, 2026

## Overview

This is the initial release of go-phoenix, the Go SDK for [Arize Phoenix](https://phoenix.arize.com/) - an open-source observability platform for LLM applications.

go-phoenix provides three main components that mirror the Python SDK:

- **phoenix-otel**: OpenTelemetry integration for sending traces
- **phoenix-client**: REST API client for managing resources
- **phoenix-evals**: Evaluation framework for LLM outputs

## Installation

```bash
go get github.com/agentplexus/go-phoenix@v0.1.0
```

## Features

### OpenTelemetry Integration (`otel` package)

Send traces to Phoenix using the standard OpenTelemetry SDK:

```go
import phoenixotel "github.com/agentplexus/go-phoenix/otel"

tp, err := phoenixotel.Register(
    phoenixotel.WithProjectName("my-app"),
    phoenixotel.WithBatch(true),
)
defer tp.Shutdown(ctx)

tracer := tp.Tracer("my-service")
ctx, span := tracer.Start(ctx, "llm-call")
span.End()
```

Key features:

- OTLP HTTP exporter with automatic endpoint configuration
- OpenInference semantic conventions for LLM observability
- Batch and sync span processing modes
- Environment variable configuration

### Phoenix Cloud Support

Connect to Phoenix Cloud with just two environment variables:

```bash
export PHOENIX_API_KEY=your-api-key
export PHOENIX_SPACE_ID=your-space-id  # From app.phoenix.arize.com/s/{space-id}
```

Or configure programmatically:

```go
tp, err := phoenixotel.Register(
    phoenixotel.WithSpaceID("your-space-id"),
    phoenixotel.WithAPIKey("your-api-key"),
    phoenixotel.WithProjectName("my-project"),
)
```

### REST API Client

Full REST API client for managing Phoenix resources:

```go
import phoenix "github.com/agentplexus/go-phoenix"

client, err := phoenix.NewClient(
    phoenix.WithSpaceID("your-space-id"),
    phoenix.WithAPIKey("your-api-key"),
)

// Projects
projects, _, err := client.ListProjects(ctx)
project, err := client.CreateProject(ctx, "my-project")

// Datasets
datasets, _, err := client.ListDatasets(ctx)
dataset, err := client.CreateDataset(ctx, "my-dataset", examples)

// Prompts
prompt, err := client.CreatePrompt(ctx, "my-prompt", template, model, provider)
prompt, err := client.GetPromptLatest(ctx, "my-prompt")

// Annotations
err := client.CreateSpanAnnotation(ctx, spanID, "quality", 0.95)
```

### Evaluation Framework (`evals` package)

Run evaluation metrics on LLM outputs:

```go
import (
    "github.com/agentplexus/go-phoenix/evals"
    "github.com/agentplexus/omniobserve/llmops/metrics"
)

evaluator := evals.NewEvaluator(phoenixClient)

result, err := evaluator.Evaluate(ctx, llmops.EvalInput{
    Input:   "What is the capital of France?",
    Output:  "Paris",
    SpanID:  "optional-span-id",  // Records results to Phoenix
}, metrics.NewExactMatchMetric())
```

### OpenInference Span Kinds

Full support for OpenInference semantic conventions:

| Span Kind | Constant | Description |
|-----------|----------|-------------|
| LLM | `SpanKindLLM` | LLM inference calls |
| Chain | `SpanKindChain` | Chain/workflow operations |
| Tool | `SpanKindTool` | Tool/function calls |
| Agent | `SpanKindAgent` | Agent operations |
| Retriever | `SpanKindRetriever` | Document retrieval |
| Embedding | `SpanKindEmbedding` | Embedding generation |
| Reranker | `SpanKindReranker` | Reranking operations |
| Guardrail | `SpanKindGuardrail` | Guardrail checks |

### omniobserve Integration

Use go-phoenix as a provider for the unified [omniobserve](https://github.com/agentplexus/omniobserve) interface:

```go
import (
    "github.com/agentplexus/omniobserve/llmops"
    _ "github.com/agentplexus/go-phoenix/llmops"  // Register provider
)

provider, err := llmops.Open("phoenix",
    llmops.WithAPIKey(os.Getenv("PHOENIX_API_KEY")),
    llmops.WithWorkspace(os.Getenv("PHOENIX_SPACE_ID")),
)
defer provider.Close()

ctx, trace, err := provider.StartTrace(ctx, "my-trace")
ctx, span, err := provider.StartSpan(ctx, "my-span",
    llmops.WithSpanType(llmops.SpanTypeLLM),
)
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `PHOENIX_API_KEY` | API key for authentication (required for Phoenix Cloud) |
| `PHOENIX_SPACE_ID` | Space identifier for Phoenix Cloud |
| `PHOENIX_PROJECT_NAME` | Default project name for traces |
| `PHOENIX_COLLECTOR_ENDPOINT` | Custom collector endpoint (optional) |
| `PHOENIX_CLIENT_HEADERS` | Additional headers (W3C Baggage format) |

## Dependencies

- `go.opentelemetry.io/otel` v1.39.0
- `github.com/agentplexus/omniobserve` v0.5.0
- `github.com/ogen-go/ogen` v1.18.0

## Known Limitations

- Auto-instrumentation is not available (Go limitation - no monkey patching)
- gRPC transport not yet implemented (HTTP only)
- Some Phoenix Cloud REST API responses have parsing issues (CreateDataset, CreatePrompt)

## What's Next

Planned for future releases:

- gRPC transport support
- Additional evaluation metrics
- Experiment run helpers
- Session management

## Links

- [GitHub Repository](https://github.com/agentplexus/go-phoenix)
- [Arize Phoenix](https://phoenix.arize.com/)
- [Phoenix Documentation](https://docs.arize.com/phoenix)
- [OpenInference Specification](https://github.com/Arize-ai/openinference)
