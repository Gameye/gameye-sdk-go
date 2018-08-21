package models

type TemplateQueryState struct {
	Template TemplateQueryArgIndex
}

type TemplateQueryArgIndex = map[string]*TemplateQueryArgItem

type TemplateQueryArgItem struct {
	TemplateKey string
	Arg         ArgConfig
}

type ArgConfig = []ArgConfigItem

type ArgConfigItem = interface{}

type NumberArgConfigItem struct {
	// type "number"
	Name             string
	DefaultValue     int
	MinimumValue     int
	MinimumExclusive bool
	MaximumValue     int
	MaximumExclusive bool
	Option           []int
}
type StringArgConfigItem struct {
	// type "string"
	Name               string
	DefaultValue       string
	ValidatePattern    string
	ValidateIgnoreCase bool
	Option             []string
}
