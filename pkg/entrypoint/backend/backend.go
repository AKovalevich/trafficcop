package backend

import (
	"net/http"
	"fmt"

	"github.com/lytics/multibayes"
	"github.com/AKovalevich/trafficcop/pkg/route"
)

type Entrypoint struct {
	Name string
	Instance *multibayes.Classifier
	Routes []route.Route
}

// Create new entrypoint
func New() *Entrypoint {
	entrypoint := &Entrypoint{}
	entrypoint.Instance = multibayes.NewClassifier()
	entrypoint.Instance.MinClassSize = 0

	return entrypoint
}

func (txe *Entrypoint) RoutesList() []route.Route {
	return txe.Routes
}

// Start entrypoint
func (txe *Entrypoint) Start() {}

// Stop enptrypoint
func (txe *Entrypoint) Stop() {}

// Initialize entrypoint
func (txe *Entrypoint) Init() {
	txe.Routes = []route.Route{
		{
			Path: "/hello",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Welcome!\n")
			},
		},
	}
}

func (txe *Entrypoint) String() string {
	return txe.Name
}
