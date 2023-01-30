//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package aps

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertManagerConfiguration) validateAddInhibitRuleParameters(id *string, options AlertManagerInhibitRuleProps) error {
	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateAddRecieverParameters(id *string, options *AlertManagerReceiverProps) error {
	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateAddTemplateParameters(template AlertManagerTemplate) error {
	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateAddTimeIntervalParameters(id *string, options *TimeIntervalProps) error {
	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateGenerateTemplateNameParameters(idx *float64) error {
	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateRenderAlertManagerConfigParameters(scope constructs.IConstruct) error {
	return nil
}

func (a *jsiiProxy_AlertManagerConfiguration) validateRenderTemplatesParameters(_scope constructs.IConstruct) error {
	return nil
}

func validateAlertManagerConfiguration_FromFullConfigurationContentParameters(scope constructs.IConstruct, id *string, content *string) error {
	return nil
}

func validateAlertManagerConfiguration_FromFullConfigurationFileParameters(scope constructs.IConstruct, id *string, path *string) error {
	return nil
}

func validateAlertManagerConfiguration_FromSplitConfigurationContentParameters(scope constructs.IConstruct, id *string, content *string) error {
	return nil
}

func validateAlertManagerConfiguration_FromSplitConfigurationFilesParameters(scope constructs.IConstruct, id *string, path *string) error {
	return nil
}

func validateAlertManagerConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAlertManagerConfigurationParameters(scope constructs.IConstruct, id *string, props *AlertManagerConfigurationProps) error {
	return nil
}

