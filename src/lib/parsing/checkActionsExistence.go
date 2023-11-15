package parsing

import "ft_ality/src/lib"

func checkActionsExistence(
	inputActions []string,
	availableActions []string,
) bool {
	if len(inputActions) > 0 {
		if lib.Contains(inputActions[0], availableActions) == false {
			return (false)
		} else {
			return (checkActionsExistence(inputActions[1:], availableActions))
		}
	}
	return (true)
}
