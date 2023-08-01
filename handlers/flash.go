package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

const (
	AlertSuccess = "AlertSuccess"
	AlertError   = "AlerError"
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

// NotifyAlert is the struct type to pass in RedirectFlash.Data map if you wish
// to trigger a notify.
type NotifyAlert struct {
	AlertType    string
	AlertTitle   string
	AlertMessage string
}

// Creates and return a new RedirectFlash struct
func NewNotifyAlert(aType string, aTitle string, aMessage string) NotifyAlert {
	return NotifyAlert{
		AlertType:    aType,
		AlertTitle:   aTitle,
		AlertMessage: aMessage,
	}
}

// AddRedirectData is a global optional type to write on fiber.Map
type AddRedirectData func(*fiber.Map)

// RedirectWithFlash will take the context, the data and the url to redirect the user
// optionally it will take any AddRedirectData functions to add values to RedirectFlash.Data
func HandleRedirectWithFlash(rf RedirectFlash, ard ...AddRedirectData) error {
	for _, fn := range ard {
		fn(&rf.Data)
	}

	log.Println("Redirect flash internal: ", rf)

	return flash.WithData(rf.Ctx, rf.Data).Redirect(rf.URL)
}

// WithNotifyAlert adds a notify alert data to the data within a RedirectWithFlash.
func WithNotifyAlert(na NotifyAlert) AddRedirectData {
	return func(d *fiber.Map) {
		(*d)["NotifyAlert"] = na
	}
}
