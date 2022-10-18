//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package glue

import (
	"fmt"
)

func (i *jsiiProxy_ICrawlerTarget) validateBindParameters(crawler Crawler) error {
	if crawler == nil {
		return fmt.Errorf("parameter crawler is required, but nil was provided")
	}

	return nil
}

