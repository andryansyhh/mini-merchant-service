package formatter

import (
	"mini-merchant-service/entity"
	"time"
)

type OutletFormat struct {
	ID         string `json:"id"`
	OutletName string `json:"outlet_name"`
	Picture    string `json:"picture"`
	UserID     string `json:"user_id"`
}

type OutletDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatOutlet(outlet entity.Outlets) OutletFormat {
	var formatOutlet = OutletFormat{
		ID:         outlet.OutletID,
		OutletName: outlet.OutletName,
		Picture:    outlet.Picture,
		UserID:     outlet.UserID,
	}

	return formatOutlet
}

func FormatDeleteOutlet(msg string) OutletDeleteFormat {
	var deleteFormat = OutletDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
