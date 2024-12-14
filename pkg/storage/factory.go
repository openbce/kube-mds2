package storage

type EngineBuilder func(string) Storage

var engines = make(map[string]EngineBuilder)

func New(engine, backend string) Storage {
	return engines[engine](backend)
}

func Register(name string, engine EngineBuilder) {
	engines[name] = engine
}
