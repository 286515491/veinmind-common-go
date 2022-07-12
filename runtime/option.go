package runtime

import (
	"github.com/chaitin/veinmind-common-go/pkg/auth"
	"github.com/pkg/errors"
)

type Option func(c Client) (Client, error)

// WithAuth parse auth path to config entity
func WithAuth(path string) Option {
	return func(c Client) (Client, error) {
		if path == "" {
			return nil, errors.New("auth config path can't be empty")
		}

		authConfig, err := auth.ParseAuthConfig(path)
		if err != nil {
			return nil, err
		}

		err = c.Auth(*authConfig)
		if err != nil {
			return nil, err
		}

		return c, nil
	}
}
