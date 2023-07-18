package data

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model

	Free       uint
	Basic      uint
	Enterprise uint
}
