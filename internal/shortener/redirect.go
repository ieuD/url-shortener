package shortener

import "github.com/ieud/url-shortener/internal/model"

type RedirectRepository interface {
	Find(code string) (*model.Redirect, error)
	Store(redirect *model.Redirect) error
}
