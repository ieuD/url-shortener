package shortener

import "github.com/ieud/url-shortener/internal/model"

type RedirectService interface {
	Find(code string) (*model.Redirect, error)
	Store(redirect *model.Redirect) error
}
