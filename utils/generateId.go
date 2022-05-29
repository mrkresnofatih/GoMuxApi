package utils

import (
	xid "github.com/rs/xid"
)

func GeneratePrefixedId(prefix string) string {
	return prefix + "-" + xid.New().String()
}
