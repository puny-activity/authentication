package werr

import "fmt"

func WrapEE(highError error, lowError error) error {
	return fmt.Errorf("%w: %w", highError, lowError)
}

func WrapES(highError error, lowString error) error {
	return fmt.Errorf("%w: %s", highError, lowString)
}

func WrapSE(highString error, lowError error) error {
	return fmt.Errorf("%s: %w", highString, lowError)
}
