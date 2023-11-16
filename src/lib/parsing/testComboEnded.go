package parsing

import "github.com/benbjohnson/immutable"

// RETURNS TRUE IF A NEW COMBO IS POSSIBLE
func TestComboEnded(
	currCombo string,
	comboSetItr immutable.SetIterator[string],
) bool {
	if comboSetItr.Done() == false {
		currComboCmp, _ := comboSetItr.Next()
		if len(currCombo) < len(currComboCmp) &&
			currCombo == currComboCmp[0:len(currCombo)] {
			return (true)
		} else {
			return TestComboEnded(currCombo, comboSetItr)
		}
	}
	return (false)
}
