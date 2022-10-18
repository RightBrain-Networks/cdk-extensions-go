//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JdbcTarget) validateAddExclusionParameters(exclusion *string) error {
	return nil
}

func (j *jsiiProxy_JdbcTarget) validateAddPathParameters(path *string) error {
	return nil
}

func (j *jsiiProxy_JdbcTarget) validateBindParameters(_crawler Crawler) error {
	return nil
}

func validateNewJdbcTargetParameters(connection Connection, options *JdbcTargetOptions) error {
	return nil
}

