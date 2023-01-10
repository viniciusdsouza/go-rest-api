package main

import (
	"fmt"
	"net/http"

	"github.com/viniciusdsouza/go-rest-api/controller"
	router "github.com/viniciusdsouza/go-rest-api/http"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
