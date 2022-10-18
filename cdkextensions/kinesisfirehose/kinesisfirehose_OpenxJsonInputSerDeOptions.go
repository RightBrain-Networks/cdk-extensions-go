package kinesisfirehose


type OpenxJsonInputSerDeOptions struct {
	CaseInsensitive *bool `field:"optional" json:"caseInsensitive" yaml:"caseInsensitive"`
	ColumnKeyMappings *map[string]*string `field:"optional" json:"columnKeyMappings" yaml:"columnKeyMappings"`
	ConvertDotsToUnderscores *bool `field:"optional" json:"convertDotsToUnderscores" yaml:"convertDotsToUnderscores"`
}

