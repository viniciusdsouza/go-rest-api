package main

import (
	"fmt"
	"net/http"

	"github.com/viniciusdsouza/go-rest-api/controller"
	router "github.com/viniciusdsouza/go-rest-api/http"
	"github.com/viniciusdsouza/go-rest-api/repository"
	"github.com/viniciusdsouza/go-rest-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.GETBYID("/posts/", postController.GetByIdPost)

	httpRouter.SERVE(port)
}
