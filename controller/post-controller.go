package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/viniciusdsouza/go-rest-api/entity"
	"github.com/viniciusdsouza/go-rest-api/errors"
	"github.com/viniciusdsouza/go-rest-api/service"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	GetByIdPost(w http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	result, err := postService.Create(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (*controller) GetByIdPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	post, err := postService.FindOnebyId(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the post"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
