package selectors

import "github.com/Gameye/gameye-sdk-go/src/models"

type TemplateItem = models.TemplateQueryArgItem

/**
 * Select a list of templates.
 * @param templateState template state
 */
func SelectTemplateList(
	templateState *models.TemplateQueryState,
) (
	templateList []*TemplateItem,
) {
	// return Object.values(templateState.template).
	//     filter(Boolean).
	//     map(i => i as TemplateItem);
	return
}

/**
 * Get details about a single template from a template-state as returned by
 * the gameye api.
 * @param templateState template state
 * @param templateKey identifier of the template
 */
func SelectTemplateItem(
	templateState *models.TemplateQueryState,
	templateKey string,
) (
	templateItem *TemplateItem,
) {
	// const templateItem = templateState.template[templateKey];
	// if (!templateItem) return null;
	// return templateItem;
	return
}
