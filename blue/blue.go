package blue

import (
	"net/http"
	"play-go/blue/context"
	"play-go/blue/router"
	"play-go/blue/types"
)

type RouterGroup struct {
	Enginee     *Enginee
	Prefix      string
	Parent      *RouterGroup
	Middlewares []types.HandlerFunc
}

type Enginee struct {
	*RouterGroup
	router *router.Router
	Groups []*RouterGroup
}

func New() *Enginee {
	enginee := &Enginee{
		router: router.NewRouter(),
	}

	// 作为最顶层的 Group
	enginee.RouterGroup = &RouterGroup{Enginee: enginee}
	enginee.Groups = []*RouterGroup{enginee.RouterGroup}

	return enginee
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	enginee := g.Enginee
	newGroup := &RouterGroup{
		Prefix:  g.Prefix + prefix,
		Enginee: enginee,
		Parent:  g,
	}
	enginee.Groups = append(enginee.Groups, newGroup)
	return newGroup
}

func (g *RouterGroup) GET(router string, handler types.HandlerFunc) {
	g.addRoute("GET", router, handler)
}

func (g *RouterGroup) POST(router string, handler types.HandlerFunc) {
	g.addRoute("POST", router, handler)
}

func (g *RouterGroup) addRoute(method string, comp string, handler types.HandlerFunc) {
	pattern := g.Prefix + comp
	g.Enginee.router.AddRoute(method, pattern, handler)
}

func (e *Enginee) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Enginee) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := context.NewContext(w, r)
	e.router.Handler(context)
}
