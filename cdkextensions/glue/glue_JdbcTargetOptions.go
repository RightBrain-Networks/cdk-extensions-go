package glue


// Configuration for Crawler JDBC target.
type JdbcTargetOptions struct {
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

