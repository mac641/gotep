package lib

import "sync"

type context struct {
	configEnvironment string
	configFilePath    string
	filePath          string
	pathPrefix        string
	verbose           bool
}

var (
	once = sync.Once{}
	ctx  *context
)

// Creates one thread-safe instance of context using default values
func GetContext() *context {
	once.Do(func() {
		ctx = &context{
			configEnvironment: "default",
			configFilePath:    "http-client.env.json",
			filePath:          "default.http",
			pathPrefix:        ".",
			verbose:           false,
		}
	})
	return ctx
}

func (ctx *context) GetConfigEnvironment() string {
	return ctx.configEnvironment
}

func (ctx *context) GetConfigFilePath() string {
	return ctx.configFilePath
}

func (ctx *context) GetFilePath() string {
	return ctx.filePath
}

func (ctx *context) GetPathPrefix() string {
	return ctx.pathPrefix
}

func (ctx *context) GetVerbose() bool {
	return ctx.verbose
}

func (ctx *context) SetConfigEnvironment(configEnvironment string) {
	ctx.configEnvironment = configEnvironment
}

func (ctx *context) SetConfigFilePath(configFilePath string) {
	ctx.configFilePath = configFilePath
}

func (ctx *context) SetFilePath(filePath string) {
	ctx.filePath = filePath
}

func (ctx *context) SetPathPrefix(pathPrefix string) {
	ctx.pathPrefix = pathPrefix
}

func (ctx *context) SetVerbose(verbose bool) {
	ctx.verbose = verbose
}
