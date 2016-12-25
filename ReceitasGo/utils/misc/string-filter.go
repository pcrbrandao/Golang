package misc

import "regexp"

// Usado para validar strings
// alfanumerico.MatchString() bool faz a análise
var ALFANUMERICO = regexp.MustCompile(`[0-9A-Za-z]$`)
