package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/viniciusdsouza/go-rest-api/entity"
)

type PostRepositoryMock struct {
	mock.Mock
}

func (repositoryMock *PostRepositoryMock) Save(post *entity.Post) (*entity.Post, error) {
	args := repositoryMock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (repositoryMock *PostRepositoryMock) FindAll() ([]entity.Post, error) {
	args := repositoryMock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (repositoryMock *PostRepositoryMock) FindOneById(id int64) (*entity.Post, error) {
	args := repositoryMock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	repositoryMock := new(PostRepositoryMock)

	var id int64 = 1

	// Setup expectations
	post := entity.Post{
		ID:    id,
		Title: "Teste",
		Text:  "Testando",
	}
	repositoryMock.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(repositoryMock)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	repositoryMock.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, post.ID, result[0].ID)
	assert.Equal(t, post.Title, result[0].Title)
	assert.Equal(t, post.Text, result[0].Text)
}
func TestCreate(t *testing.T) {
	repositoryMock := new(PostRepositoryMock)

	// Setup expectations
	post := entity.Post{
		Title: "Teste",
		Text:  "Testando",
	}
	repositoryMock.On("Save").Return(&post, nil)

	testService := NewPostService(repositoryMock)

	result, err := testService.Create(&post)

	// Mock Assertion: Behavioral
	repositoryMock.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, post.ID, result.ID)
	assert.Equal(t, post.Title, result.Title)
	assert.Equal(t, post.Text, result.Text)
	assert.Nil(t, err)
}
func TestFindOnebyId(t *testing.T) {
	repositoryMock := new(PostRepositoryMock)

	var id int64 = 1

	// Setup expectations
	post := entity.Post{
		ID:    id,
		Title: "Teste",
		Text:  "Testando",
	}
	repositoryMock.On("FindOneById").Return(&post, nil)

	testService := NewPostService(repositoryMock)

	result, _ := testService.FindOnebyId(id)

	// Mock Assertion: Behavioral
	repositoryMock.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, post.ID, result.ID)
	assert.Equal(t, post.Title, result.Title)
	assert.Equal(t, post.Text, result.Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "Testando"}
	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}
