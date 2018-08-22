package mocks

import "github.com/Gameye/gameye-sdk-go/src/models"

var TemplateStateMock = models.TemplateQueryState{
	Template: models.TemplateQueryArgIndex{
		"t1": models.TemplateQueryArgItem{
			TemplateKey: "t1",
			Arg: models.ArgConfig{
				models.NumberArgConfigItem{
					Name:         "tickRate",
					DefaultValue: 64,
					Option:       []int{64, 128},
				},
			},
		},
		"t2": models.TemplateQueryArgItem{
			TemplateKey: "t2",
			Arg: models.ArgConfig{
				models.StringArgConfigItem{
					Name:         "steamToken",
					DefaultValue: "",
				},
				models.StringArgConfigItem{
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
