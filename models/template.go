package models

import "github.com/mitchellh/mapstructure"

/*
CreateTemplateQueryState will create a new TemplateQueryState from
a string / interface map
*/
func CreateTemplateQueryState(anyState *map[string]interface{}) (
	state *TemplateQueryState,
	err error,
) {
	state = &TemplateQueryState{}

	mapstructure.Decode(anyState, state)
	for _, templateItem := range state.Template {
		for argIndex, argItem := range templateItem.Arg {
			anyArgItem := argItem.(map[string]interface{})
			switch anyArgItem["type"] {
			case "string":
				typedArgItem := StringArgConfigItem{}
				mapstructure.Decode(&anyArgItem, &typedArgItem)
				templateItem.Arg[argIndex] = typedArgItem
			case "number":
				typedArgItem := NumberArgConfigItem{}
				mapstructure.Decode(&anyArgItem, &typedArgItem)
				templateItem.Arg[argIndex] = typedArgItem
			default:
				err = ErrUnknownType
				return
			}
		}
	}
	return
}

type TemplateQueryState struct {
	Template TemplateQueryArgIndex
}

type TemplateQueryArgIndex = map[string]*TemplateQueryArgItem

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
	Name               string   `mapstructure:"name"`
	DefaultValue       string   `mapstructure:"defaultValue"`
	ValidatePattern    string   `mapstructure:"validatePattern"`
	ValidateIgnoreCase bool     `mapstructure:"validateIgnoreCase"`
	Option             []string `mapstructure:"option"`
}

var TemplateStateMock = TemplateQueryState{
	Template: TemplateQueryArgIndex{
		"t1": &TemplateQueryArgItem{
			TemplateKey: "t1",
			Arg: ArgConfig{
				NumberArgConfigItem{
					Name:         "tickRate",
					DefaultValue: 64,
					Option:       []int{64, 128},
				},
			},
		},
		"t2": &TemplateQueryArgItem{
			TemplateKey: "t2",
			Arg: ArgConfig{
				StringArgConfigItem{
					Name:         "steamToken",
					DefaultValue: "",
				},
				StringArgConfigItem{
					Name:         "hostname",
					DefaultValue: "gameye.com Match Server",
				},
			},
		},
	},
}

var TemplateStateJSONMock = `{
    "template": {
        "t1": {
            "templateKey": "t1",
            "arg": [{
                "name": "tickRate",
                "type": "number",
                "defaultValue": 64,
                "option": [64, 128]
            }]
        },
        "t2": {
            "templateKey": "t2",
            "arg": [{
                "name": "steamToken",
                "type": "string",
                "defaultValue": ""
            }, {
                "name": "hostname",
                "type": "string",
                "defaultValue": "gameye.com Match Server"
            }]
        }
    }
}`
