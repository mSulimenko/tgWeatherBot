package router

type Router struct {
	commands map[string]handler
}

type handler func(update Update)
