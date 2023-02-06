package ssd

import (
	"github.com/deckarep/golang-set"
)

type modelKey struct {
	brand        mapset.Set
	model        mapset.Set
	productPages mapset.Set
	others       mapset.Set
}
