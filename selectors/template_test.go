package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectTemplateList(t *testing.T) {
	templateList := SelectTemplateList(templateStateMock)
	assert.Equal(t, 2, len(templateList))
	for _, templateItem := range templateList {
		switch templateItem.TemplateKey {
		case "t1":
		case "t2":
		default:
			assert.Fail(t, templateItem.TemplateKey)
		}
	}
}

func TestSelectTemplateItem(t *testing.T) {
	templateItem := SelectTemplateItem(templateStateMock, "t2")
	assert.NotNil(t, templateItem)
	if templateItem != nil {
		assert.Equal(t, "t2", templateItem.TemplateKey)
	}
}
