package services

import (
	"errors"
	"time"

	"github.com/ieud/url-shortener/internal/model"
	"github.com/ieud/url-shortener/internal/repository"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepo repository.RedirectRepository
}

// Find implements repository.RedirectService
func (r *redirectService) Find(code string) (*model.Redirect, error) {
	return r.redirectRepo.Find(code)
}

// Store implements repository.RedirectService
func (r *redirectService) Store(redirect *model.Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "Services.Shortener.Store")
	}

	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}

func NewRedirectService(redirectRepo repository.RedirectRepository) repository.RedirectService {
	return &redirectService{
		redirectRepo,
	}
}
