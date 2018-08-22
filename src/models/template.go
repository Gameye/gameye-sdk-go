package models

type TemplateQueryState struct {
	Template TemplateQueryArgIndex
}

type TemplateQueryArgIndex = map[string]TemplateQueryArgItem

type TemplateQueryArgItem struct {
	TemplateKey string    `mapstructure:"templateKey"`
	Arg         ArgConfig `mapstructure:"arg"`
}

type ArgConfig = []ArgConfigItem

type ArgConfigItem = interface{}

type NumberArgConfigItem struct {
	// type "number"
	Name             string `mapstructure:"name"`
	DefaultValue     int    `mapstructure:"defaultValue"`
	MinimumValue     int    `mapstructure:"minimumValue"`
	MinimumExclusive bool   `mapstructure:"minimumExclusive"`
	MaximumValue     int    `mapstructure:"maximumValue"`
	MaximumExclusive bool   `mapstructure:"maximumExclusive"`
	Option           []int  `mapstructure:"option"`
}
type StringArgConfigItem struct {
	// type "string"
	Name               string   `mapstructure:"Name"`
	DefaultValue       string   `mapstructure:"DefaultValue"`
	ValidatePattern    string   `mapstructure:"ValidatePattern"`
	ValidateIgnoreCase bool     `mapstructure:"ValidateIgnoreCase"`
	Option             []string `mapstructure:"Option"`
}
