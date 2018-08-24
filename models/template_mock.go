package models

var templateStateMock = &TemplateQueryState{
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

var templateStateJSONMock = `{
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
