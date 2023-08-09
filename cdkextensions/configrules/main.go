package configrules

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.config_rules.IamPasswordPolicy",
		reflect.TypeOf((*IamPasswordPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "configRuleArn", GoGetter: "ConfigRuleArn"},
			_jsii_.MemberProperty{JsiiProperty: "configRuleComplianceType", GoGetter: "ConfigRuleComplianceType"},
			_jsii_.MemberProperty{JsiiProperty: "configRuleId", GoGetter: "ConfigRuleId"},
			_jsii_.MemberProperty{JsiiProperty: "configRuleName", GoGetter: "ConfigRuleName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isCustomWithChanges", GoGetter: "IsCustomWithChanges"},
			_jsii_.MemberProperty{JsiiProperty: "isManaged", GoGetter: "IsManaged"},
			_jsii_.MemberProperty{JsiiProperty: "maxPasswordAge", GoGetter: "MaxPasswordAge"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordLength", GoGetter: "MinimumPasswordLength"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onComplianceChange", GoMethod: "OnComplianceChange"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onReEvaluationStatus", GoMethod: "OnReEvaluationStatus"},
			_jsii_.MemberProperty{JsiiProperty: "passwordReusePrevention", GoGetter: "PasswordReusePrevention"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "remediationConfiguration", GoGetter: "RemediationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "remediationPolicy", GoGetter: "RemediationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "remediationRole", GoGetter: "RemediationRole"},
			_jsii_.MemberProperty{JsiiProperty: "requireLowercaseCharacters", GoGetter: "RequireLowercaseCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "requireNumbers", GoGetter: "RequireNumbers"},
			_jsii_.MemberProperty{JsiiProperty: "requireSymbols", GoGetter: "RequireSymbols"},
			_jsii_.MemberProperty{JsiiProperty: "requireUppercaseCharacters", GoGetter: "RequireUppercaseCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "ruleScope", GoGetter: "RuleScope"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IamPasswordPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awsconfigManagedRule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.config_rules.IamPasswordPolicyProps",
		reflect.TypeOf((*IamPasswordPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.config_rules.VpcDefaultSecurityGroupClosed",
		reflect.TypeOf((*VpcDefaultSecurityGroupClosed)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "configRuleArn", GoGetter: "ConfigRuleArn"},
			_jsii_.MemberProperty{JsiiProperty: "configRuleComplianceType", GoGetter: "ConfigRuleComplianceType"},
			_jsii_.MemberProperty{JsiiProperty: "configRuleId", GoGetter: "ConfigRuleId"},
			_jsii_.MemberProperty{JsiiProperty: "configRuleName", GoGetter: "ConfigRuleName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isCustomWithChanges", GoGetter: "IsCustomWithChanges"},
			_jsii_.MemberProperty{JsiiProperty: "isManaged", GoGetter: "IsManaged"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onComplianceChange", GoMethod: "OnComplianceChange"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onReEvaluationStatus", GoMethod: "OnReEvaluationStatus"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "remediationConfiguration", GoGetter: "RemediationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "remediationPolicy", GoGetter: "RemediationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "remediationRole", GoGetter: "RemediationRole"},
			_jsii_.MemberProperty{JsiiProperty: "ruleScope", GoGetter: "RuleScope"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VpcDefaultSecurityGroupClosed{}
			_jsii_.InitJsiiProxy(&j.Type__awsconfigManagedRule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.config_rules.VpcDefaultSecurityGroupClosedProps",
		reflect.TypeOf((*VpcDefaultSecurityGroupClosedProps)(nil)).Elem(),
	)
}
