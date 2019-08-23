package route

import (
    "github.com/gorilla/mux"
    classificationbox "github.com/heaptracetechnology/machinebox-classificationbox/classificationbox"
    "log"
    "net/http"
)

//Route struct
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

//Routes array
type Routes []Route

var routes = Routes{
    Route{
        "CreateModel",
        "POST",
        "/createModel",
        classificationbox.CreateModel,
    },
    Route{
        "TeachModel",
        "POST",
        "/teachModel",
        classificationbox.TeachModel,
    },
    Route{
        "GetModel",
        "POST",
        "/getModel",
        classificationbox.GetModel,
    },
    Route{
        "MakePredictions",
        "POST",
        "/makePredictions",
        classificationbox.MakePredictions,
    },
    Route{
        "ListingModels",
        "POST",
        "/listModels",
        classificationbox.ListingModels,
    },
    Route{
        "DeleteModel",
        "POST",
        "/deleteModel",
        classificationbox.DeleteModel,
    },
}

//NewRouter route
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
