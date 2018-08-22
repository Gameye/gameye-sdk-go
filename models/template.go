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

/*
TemplateQueryState is data coming from the api
*/
type TemplateQueryState struct {
	Template TemplateQueryArgIndex
}

/*
TemplateQueryArgIndex is data coming from the api
*/
type TemplateQueryArgIndex = map[string]*TemplateQueryArgItem

/*
TemplateQueryArgItem is data coming from the api
*/
type TemplateQueryArgItem struct {
	TemplateKey string    `mapstructure:"templateKey"`
	Arg         ArgConfig `mapstructure:"arg"`
}

/*
ArgConfig is data coming from the api
*/
type ArgConfig = []ArgConfigItem

/*
ArgConfigItem is data coming from the api
*/
type ArgConfigItem = interface{}

/*
NumberArgConfigItem is data coming from the api
*/
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

/*
StringArgConfigItem is data coming from the api
*/
type StringArgConfigItem struct {
	// type "string"
	Name               string   `mapstructure:"name"`
	DefaultValue       string   `mapstructure:"defaultValue"`
	ValidatePattern    string   `mapstructure:"validatePattern"`
	ValidateIgnoreCase bool     `mapstructure:"validateIgnoreCase"`
	Option             []string `mapstructure:"option"`
}

/*
TemplateStateMock is mock data
*/
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

/*
TemplateStateJSONMock is json mock data
*/
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
