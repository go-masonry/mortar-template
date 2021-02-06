package validations

import (
	"context"
	"fmt"
)

func (impl *helloworldValidationsImpl) CheckAuth(ctx context.Context) error {
	// we only log for error here
	token, err := impl.deps.TokenExtractor.FromContext(ctx)
	if err == nil {
		_, err = token.Map()
		if err != nil { // now that's an error
			return fmt.Errorf("failed to produce a map from authorization header, %w", err)
		}
	} else {
		impl.deps.Logger.WithError(err).Warn(ctx, "token extraction from Auth header failed")
	}
	return nil
}
