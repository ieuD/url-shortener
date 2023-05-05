package shortener

import (
	"errors"
	"github.com/ieud/url-shortener/internal/model"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepo,
	}
}

func (r *redirectService) Find(code string) (*model.Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *model.Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}
