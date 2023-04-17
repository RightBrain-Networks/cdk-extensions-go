//go:build no_runtime_type_checking

package athena

// Building without runtime type checking enabled, so all the below just return nil

func validateAnalyticsEngine_ApacheSparkParameters(options ApacheSparkEngineOptions) error {
	return nil
}

func validateAnalyticsEngine_AthenaSqlParameters(options *AthenaSqlEngineOptions) error {
	return nil
}

