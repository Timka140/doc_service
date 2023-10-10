package post

import (
	"fmt"
	"log"

	"projects/doc/doc_service/internal/web_server/post/methods_post"

	"github.com/gin-gonic/gin"
)

type IPost interface {
}
type TPost struct {
	router *gin.Engine
}
type TInPost struct {
	Router *gin.Engine
}

func NewPost(in *TInPost) (IPost, error) {
	t := &TPost{
		router: in.Router,
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("NewPost(): критическая ошибка, err=%v", r)
		}
	}()

	posts, err := methods_post.NewMethodsPost()
	if err != nil {
		return nil, fmt.Errorf("NewPost(): инициализация методов post, err=%w", err)
	}

	err = posts.Range(&methods_post.TInPostPage{}, func(page methods_post.IPostPage) {
		t.router.POST(fmt.Sprintf("/api/%v", page.GetPath()), page.GetContext)
	})
	if err != nil {
		return nil, fmt.Errorf("NewPost(): добавление routers post, err=%w", err)
	}

	return t, nil
}
