package mystenbcs

// https://github.com/fardream/go-bcs/blob/main/bcs/tag.go
import (
	"fmt"
	"strings"
)

const tagName = "bcs"

type tagValue int64

const (
	tagValue_Optional tagValue = 1 << iota // optional
	tagValue_Ignore                        // -
)

func parseTagValue(tag string) (tagValue, error) {
	var r tagValue
	tagSegs := strings.Split(tag, ",")
	for _, seg := range tagSegs {
		seg := strings.TrimSpace(seg)
		if seg == "" {
			continue
		}
		switch seg {
		case "optional":
			r |= tagValue_Optional
		case "-":
			return tagValue_Ignore, nil
		default:
			return 0, fmt.Errorf("unknown tag: %s in %s", seg, tag)
		}
	}

	return r, nil
}

func (t tagValue) isOptional() bool {
	return t&tagValue_Optional != 0
}

func (t tagValue) isIgnored() bool {
	return t&tagValue_Ignore != 0
}
