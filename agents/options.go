package agents

import (
	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
)

type CreationOptions struct {
	prompt                  prompts.PromptTemplate
	memory                  schema.Memory
	callbacksHandler        callbacks.Handler
	errorHandler            *ParserErrorHandler
	maxIterations           int
	returnIntermediateSteps bool
	outputKey               string
	promptPrefix            string
	formatInstructions      string
	promptSuffix            string

	// openai
	systemMessage string
	extraMessages []prompts.MessageFormatter
}

// CreationOption is a function type that can be used to modify the creation of the agents
// and executors.
type CreationOption func(*CreationOptions)

func executorDefaultOptions() CreationOptions {
	return CreationOptions{
		maxIterations: _defaultMaxIterations,
		outputKey:     _defaultOutputKey,
		memory:        memory.NewSimple(),
	}
}

func mrklDefaultOptions() CreationOptions {
	return CreationOptions{
		promptPrefix:       _defaultMrklPrefix,
		formatInstructions: _defaultMrklFormatInstructions,
		promptSuffix:       _defaultMrklSuffix,
		outputKey:          _defaultOutputKey,
	}
}

func conversationalDefaultOptions() CreationOptions {
	return CreationOptions{
		promptPrefix:       _defaultConversationalPrefix,
		formatInstructions: _defaultConversationalFormatInstructions,
		promptSuffix:       _defaultConversationalSuffix,
		outputKey:          _defaultOutputKey,
	}
}

func openAIFunctionsDefaultOptions() CreationOptions {
	return CreationOptions{
		systemMessage: "You are a helpful AI assistant.",
		outputKey:     _defaultOutputKey,
	}
}

func (co CreationOptions) getMrklPrompt(tools []tools.Tool) prompts.PromptTemplate {
	if co.prompt.Template != "" {
		return co.prompt
	}

	return createMRKLPrompt(
		tools,
		co.promptPrefix,
		co.formatInstructions,
		co.promptSuffix,
	)
}

func (co CreationOptions) getConversationalPrompt(tools []tools.Tool) prompts.PromptTemplate {
	if co.prompt.Template != "" {
		return co.prompt
	}

	return createConversationalPrompt(
		tools,
		co.promptPrefix,
		co.formatInstructions,
		co.promptSuffix,
	)
}

// WithMaxIterations is an option for setting the max number of iterations the executor
// will complete.
func WithMaxIterations(iterations int) CreationOption {
	return func(co *CreationOptions) {
		co.maxIterations = iterations
	}
}

// WithOutputKey is an option for setting the output key of the agent.
func WithOutputKey(outputKey string) CreationOption {
	return func(co *CreationOptions) {
		co.outputKey = outputKey
	}
}

// WithPromptPrefix is an option for setting the prefix of the prompt used by the agent.
func WithPromptPrefix(prefix string) CreationOption {
	return func(co *CreationOptions) {
		co.promptPrefix = prefix
	}
}

// WithPromptFormatInstructions is an option for setting the format instructions of the prompt
// used by the agent.
func WithPromptFormatInstructions(instructions string) CreationOption {
	return func(co *CreationOptions) {
		co.formatInstructions = instructions
	}
}

// WithPromptSuffix is an option for setting the suffix of the prompt used by the agent.
func WithPromptSuffix(suffix string) CreationOption {
	return func(co *CreationOptions) {
		co.promptSuffix = suffix
	}
}

// WithPrompt is an option for setting the prompt the agent will use.
func WithPrompt(prompt prompts.PromptTemplate) CreationOption {
	return func(co *CreationOptions) {
		co.prompt = prompt
	}
}

// WithReturnIntermediateSteps is an option for making the executor return the intermediate steps
// taken.
func WithReturnIntermediateSteps() CreationOption {
	return func(co *CreationOptions) {
		co.returnIntermediateSteps = true
	}
}

// WithMemory is an option for setting the memory of the executor.
func WithMemory(m schema.Memory) CreationOption {
	return func(co *CreationOptions) {
		co.memory = m
	}
}

// WithCallbacksHandler is an option for setting a callback handler to an executor.
func WithCallbacksHandler(handler callbacks.Handler) CreationOption {
	return func(co *CreationOptions) {
		co.callbacksHandler = handler
	}
}

// WithParserErrorHandler is an option for setting a parser error handler to an executor.
func WithParserErrorHandler(errorHandler *ParserErrorHandler) CreationOption {
	return func(co *CreationOptions) {
		co.errorHandler = errorHandler
	}
}

type OpenAIOption struct{}

func NewOpenAIOption() OpenAIOption {
	return OpenAIOption{}
}

func (o OpenAIOption) WithSystemMessage(msg string) CreationOption {
	return func(co *CreationOptions) {
		co.systemMessage = msg
	}
}

func (o OpenAIOption) WithExtraMessages(extraMessages []prompts.MessageFormatter) CreationOption {
	return func(co *CreationOptions) {
		co.extraMessages = extraMessages
	}
}
