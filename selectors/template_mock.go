package selectors

import "github.com/Gameye/gameye-sdk-go/models"

var templateStateMock = &models.TemplateQueryState{
	Template: models.TemplateQueryArgIndex{
		"t1": &models.TemplateQueryArgItem{
			TemplateKey: "t1",
			Arg: models.ArgConfig{
				models.NumberArgConfigItem{
					Name:         "tickRate",
					DefaultValue: 64,
					Option:       []int{64, 128},
				},
			},
		},
		"t2": &models.TemplateQueryArgItem{
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
