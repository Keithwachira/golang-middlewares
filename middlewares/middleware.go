
package middleware
import (
"log"
"time"
	"net/http"
)


//we define a function thats takes   in a http.Handler  and return a http.Handler
type Middleware func( http.Handler) http.Handler


func LogTime() Middleware {
	return func(h http.Handler) http.Handler {


		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Print("This request was sent at : ")
			log.Println(time.Now())
			h.ServeHTTP(w, r)
		})
	}
}





func SecondMiddleware() Middleware {
	return func(h http.Handler) http.Handler {


		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Print("I like chocolate")

			h.ServeHTTP(w, r)
		})
	}
}



func Adapt(h http.Handler, adapters ...Middleware) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}