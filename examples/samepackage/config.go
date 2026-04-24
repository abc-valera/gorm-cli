package samepackage

import "gorm.io/cli/gorm/genconfig"

var _ = genconfig.Config{
	IsSamePackage: true,
}
