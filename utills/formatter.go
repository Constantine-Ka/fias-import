package utills

import "log"

func IndexOf(arr []string, need string) int {
	if len(arr) == 0 {
		return -1
	}
	for ind, val := range arr {
		if val == need {
			return ind
		}
	}
	log.Println(000)
	return -1
}
func ComparisonInvalid(slice1, slice2 []string) []string {
	invalidElem := []string{}
	for _, s1 := range slice1 {
		if IndexOf(slice2, s1) == -1 {
			invalidElem = append(invalidElem, s1)
		}
	}
	return invalidElem

}
