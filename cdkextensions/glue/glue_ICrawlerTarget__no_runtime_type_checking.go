//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICrawlerTarget) validateBindParameters(crawler Crawler) error {
	return nil
}

