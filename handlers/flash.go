package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

const (
	AlertSuccess = "AlertSuccess"
	AlertError   = "AlertError"
	AlertWarning = "AlertWarning"
)

// RedirectFlash is a general struct type to handle RedirectWithFlash requests
type RedirectFlash struct {
	Ctx  *fiber.Ctx
	Data fiber.Map
	URL  string
}

// Creates and return a new RedirectFlash struct
func NewRedirectFlash(ctx *fiber.Ctx, data fiber.Map, url string) RedirectFlash {
	return RedirectFlash{
		Ctx:  ctx,
		Data: data,
		URL:  url,
	}
}

// LEARN HOW TO PASS COMPLEX STRUCTURES IN FIBER MAP TO BE ANDLED BY HTML TEMPLATES
// // NotifyAlert is the struct type to pass in RedirectFlash.Data map if you wish
// // to trigger a notify.
// type NotifyAlert struct {
// 	AlertType AlertType `json:"AlertType"`
// }
// // AlertType    AlertType    `json:"AlertType"`
// // AlertTitle   AlertTitle   `json:"AlertTitle"`
// // AlertMessage AlertMessage `json:"AlertMessage"`
// type AlertType struct {
// 	Type string `json:"Type"`
// }
// type AlertTitle struct {
// 	Title string
// }
// type AlertMessage struct {
// 	Message string
// }
// Creates and return a new RedirectFlash struct
// func NewNotifyAlert(aType AlertType, aTitle AlertTitle, aMessage AlertMessage) NotifyAlert {
// 	return NotifyAlert{
// 		AlertType:    aType,
// 		AlertTitle:   aTitle,
// 		AlertMessage: aMessage,
// 	}
// }

// AddRedirectData is a global optional type to write on fiber.Map
type AddRedirectData func(*fiber.Map)

// RedirectWithFlash will take the context, the data and the url to redirect the user
// optionally it will take any AddRedirectData functions to add values to RedirectFlash.Data
func HandleRedirectWithFlash(rf RedirectFlash, ard ...AddRedirectData) error {
	for _, fn := range ard {
		fn(&rf.Data)
	}
	return flash.WithData(rf.Ctx, rf.Data).Redirect(rf.URL)
}

// WithNotifyAlert adds a notify alert data to the data within a RedirectWithFlash.
func WithNotifyAlert(AlertType string, AlertTitle string, AlertMessage string) AddRedirectData {
	return func(d *fiber.Map) {
		(*d)["AlertType"] = AlertType
		(*d)["AlertTitle"] = AlertTitle
		(*d)["AlertMessage"] = AlertMessage
	}
}
