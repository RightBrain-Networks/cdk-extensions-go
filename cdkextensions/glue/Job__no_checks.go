//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Job) validateAddArgumentParameters(key *string, value *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateAddConnectionParameters(connection Connection) error {
	return nil
}

func (j *jsiiProxy_Job) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateJob_FromJobArnParameters(scope constructs.IConstruct, id *string, jobArn *string) error {
	return nil
}

func validateJob_FromJobNameParameters(scope constructs.IConstruct, id *string, jobName *string) error {
	return nil
}

func validateJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJob_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateJob_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewJobParameters(scope constructs.Construct, id *string, props *JobProps) error {
	return nil
}

