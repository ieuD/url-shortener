package json

import (
	"encoding/json"
	"github.com/ieud/url-shortener/internal/model"
	"github.com/pkg/errors"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*model.Redirect, error) {
	redirect := &model.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *model.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
