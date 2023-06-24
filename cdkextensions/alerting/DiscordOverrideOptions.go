package alerting


type DiscordOverrideOptions struct {
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	Mentions *[]*string `field:"optional" json:"mentions" yaml:"mentions"`
}

