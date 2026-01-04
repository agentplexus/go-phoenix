package otel

import "go.opentelemetry.io/otel/attribute"

// OpenInference semantic conventions for LLM observability.
// These attributes are compatible with Phoenix and the OpenInference specification.
// See: https://github.com/Arize-ai/openinference

// Span kind attributes for LLM operations.
const (
	// SpanKindLLM indicates an LLM call span.
	SpanKindLLM = "LLM"
	// SpanKindChain indicates a chain/workflow span.
	SpanKindChain = "CHAIN"
	// SpanKindTool indicates a tool call span.
	SpanKindTool = "TOOL"
	// SpanKindAgent indicates an agent span.
	SpanKindAgent = "AGENT"
	// SpanKindRetriever indicates a retrieval span.
	SpanKindRetriever = "RETRIEVER"
	// SpanKindEmbedding indicates an embedding span.
	SpanKindEmbedding = "EMBEDDING"
	// SpanKindReranker indicates a reranker span.
	SpanKindReranker = "RERANKER"
	// SpanKindGuardrail indicates a guardrail span.
	SpanKindGuardrail = "GUARDRAIL"
)

// OpenInference attribute keys.
const (
	// OpenInferenceSpanKind is the span kind attribute key.
	OpenInferenceSpanKind = "openinference.span.kind"

	// Input/Output attributes
	InputValue  = "input.value"
	InputMime   = "input.mime_type"
	OutputValue = "output.value"
	OutputMime  = "output.mime_type"

	// LLM attributes
	LLMModelName            = "llm.model_name"
	LLMProvider             = "llm.provider"
	LLMInvocationParams     = "llm.invocation_parameters"
	LLMTokenCountPrompt     = "llm.token_count.prompt"     //nolint:gosec // Not a credential
	LLMTokenCountCompletion = "llm.token_count.completion" //nolint:gosec // Not a credential
	LLMTokenCountTotal      = "llm.token_count.total"      //nolint:gosec // Not a credential

	// Message attributes
	LLMInputMessages  = "llm.input_messages"
	LLMOutputMessages = "llm.output_messages"

	// Tool attributes
	ToolName        = "tool.name"
	ToolDescription = "tool.description"
	ToolParameters  = "tool.parameters"

	// Retrieval attributes
	RetrievalDocuments = "retrieval.documents"

	// Embedding attributes
	EmbeddingModelName  = "embedding.model_name"
	EmbeddingEmbeddings = "embedding.embeddings"
	EmbeddingText       = "embedding.text"

	// Metadata attributes
	MetadataKey = "metadata"
	TagsKey     = "tag.tags"

	// Session/Thread attributes
	SessionID = "session.id"
	UserID    = "user.id"
)

// Attribute helper functions for common LLM attributes.

// WithSpanKind sets the OpenInference span kind.
func WithSpanKind(kind string) attribute.KeyValue {
	return attribute.String(OpenInferenceSpanKind, kind)
}

// WithInput sets the input value attribute.
func WithInput(input string) attribute.KeyValue {
	return attribute.String(InputValue, input)
}

// WithOutput sets the output value attribute.
func WithOutput(output string) attribute.KeyValue {
	return attribute.String(OutputValue, output)
}

// WithModelName sets the LLM model name.
func WithModelName(model string) attribute.KeyValue {
	return attribute.String(LLMModelName, model)
}

// WithLLMProvider sets the LLM provider name.
func WithLLMProvider(provider string) attribute.KeyValue {
	return attribute.String(LLMProvider, provider)
}

// WithTokenCounts sets the token count attributes.
func WithTokenCounts(prompt, completion, total int) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.Int(LLMTokenCountPrompt, prompt),
		attribute.Int(LLMTokenCountCompletion, completion),
		attribute.Int(LLMTokenCountTotal, total),
	}
}

// WithToolName sets the tool name attribute.
func WithToolName(name string) attribute.KeyValue {
	return attribute.String(ToolName, name)
}

// WithSessionID sets the session ID attribute.
func WithSessionID(id string) attribute.KeyValue {
	return attribute.String(SessionID, id)
}

// WithUserID sets the user ID attribute.
func WithUserID(id string) attribute.KeyValue {
	return attribute.String(UserID, id)
}

// WithMetadata sets a metadata attribute as JSON string.
func WithMetadata(metadata string) attribute.KeyValue {
	return attribute.String(MetadataKey, metadata)
}

// LLMSpanAttributes returns common attributes for an LLM span.
func LLMSpanAttributes(model, provider string, promptTokens, completionTokens int) []attribute.KeyValue {
	return []attribute.KeyValue{
		WithSpanKind(SpanKindLLM),
		WithModelName(model),
		WithLLMProvider(provider),
		attribute.Int(LLMTokenCountPrompt, promptTokens),
		attribute.Int(LLMTokenCountCompletion, completionTokens),
		attribute.Int(LLMTokenCountTotal, promptTokens+completionTokens),
	}
}

// ToolSpanAttributes returns common attributes for a tool span.
func ToolSpanAttributes(name, description string) []attribute.KeyValue {
	attrs := []attribute.KeyValue{
		WithSpanKind(SpanKindTool),
		WithToolName(name),
	}
	if description != "" {
		attrs = append(attrs, attribute.String(ToolDescription, description))
	}
	return attrs
}

// RetrieverSpanAttributes returns common attributes for a retriever span.
func RetrieverSpanAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		WithSpanKind(SpanKindRetriever),
	}
}

// ChainSpanAttributes returns common attributes for a chain span.
func ChainSpanAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		WithSpanKind(SpanKindChain),
	}
}

// AgentSpanAttributes returns common attributes for an agent span.
func AgentSpanAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		WithSpanKind(SpanKindAgent),
	}
}
