//go:build no_runtime_type_checking

package multi-az-observability

// Building without runtime type checking enabled, so all the below just return nil

func validateNewCanaryTestMetricsOverrideParameters(props *CanaryTestMetricsOverrideProps) error {
	return nil
}

