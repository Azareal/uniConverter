package common

import (
	"strconv"
	"strings"
	"unicode"
)

type Version struct {
	Major  int
	Minor  int
	Patch  int
	Tag    string
	TagNum int
}

func SplitVersion(version string) (*Version, error) {
	tagSplitter := strings.Split(version, "-")
	mainSplitter := strings.Split(tagSplitter[0], ".")

	var extractVersionBit = func(index int) (value int, err error) {
		if len(mainSplitter) > (index + 1) {
			value, err = strconv.Atoi(mainSplitter[index])
		}
		return value, err
	}

	// If a bit is omitted, initialise it to zero
	major, _ := extractVersionBit(0)
	minor, _ := extractVersionBit(1)
	patch, _ := extractVersionBit(2)

	var outVer = &Version{Major: major, Minor: minor, Patch: patch}

	// Some software use a dot instead of a dash for tags
	if len(mainSplitter) > 3 {
		var tagText string
		var tagNumber string
		var textBit bool
		for _, char := range mainSplitter[2] {
			if textBit {
				if unicode.IsDigit(char) {
					textBit = true
				}
				tagText += string(char)
			} else {
				tagNumber += string(char)
			}
		}

		outVer.Tag = tagText
		convNum, err := strconv.Atoi(tagNumber)
		if err != nil {
			return nil, err
		}
		outVer.TagNum = convNum
	}

	return outVer, nil
}
