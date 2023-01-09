//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package stacks

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsLoggingStack) validateAddDependencyParameters(target awscdk.Stack) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateAddTransformParameters(transform *string) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateAllocateLogicalIdParameters(cfnElement awscdk.CfnElement) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateExportStringListValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateExportValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateFormatArnParameters(components *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateGetLogicalIdParameters(element awscdk.CfnElement) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateRegionalFactParameters(factName *string) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateResolveParameters(obj interface{}) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateSplitArnParameters(arn *string, arnFormat awscdk.ArnFormat) error {
	return nil
}

func (a *jsiiProxy_AwsLoggingStack) validateToJsonStringParameters(obj interface{}) error {
	return nil
}

func validateAwsLoggingStack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAwsLoggingStack_IsStackParameters(x interface{}) error {
	return nil
}

func validateAwsLoggingStack_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAwsLoggingStackParameters(scope constructs.Construct, id *string, props *AwsLoggingStackProps) error {
	return nil
}

