package handlers

import (
	"github.com/Puddi1/GFS-Stack/data"
)

// HandleSaveCount takes the count and overrites the old user count in the databse
func HandleSaveCount(c data.Counter) (data.Counter, error) {

	return c, nil
}
