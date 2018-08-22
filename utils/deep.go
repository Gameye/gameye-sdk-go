package utils

/*
SetIn deeply sets a value in source via path
*/
func SetIn(
	source map[string]interface{},
	path []string,
	value interface{},
) (
	target map[string]interface{},
) {
	if len(path) == 0 {
		target = value.(map[string]interface{})
		return
	}

	target = make(map[string]interface{}, len(source))
	for key, value := range source {
		target[key] = value
	}

	sourceParent := source
	targetParent := target

	keyIndex := 0
	parentPathLength := len(path) - 1
	for keyIndex < parentPathLength {
		key := path[keyIndex]

		targetChild := make(map[string]interface{})
		if targetParent[key] != nil &&
			sourceParent != nil &&
			sourceParent[key] == targetParent[key] {

			for key, value := range targetParent[key].(map[string]interface{}) {
				targetChild[key] = value
			}
		}

		targetParent[key] = targetChild
		targetParent = targetChild

		keyIndex++
	}

	{
		key := path[keyIndex]
		if value == nil {
			delete(targetParent, key)
		} else {
			targetParent[key] = value
		}
	}

	return
}
