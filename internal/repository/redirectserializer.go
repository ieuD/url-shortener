package repository

import "github.com/ieud/url-shortener/internal/model"

type RedirectSerializer interface {
	Decode(input []byte) (*model.Redirect, error)
	Encode(input *model.Redirect) ([]byte, error)
}
