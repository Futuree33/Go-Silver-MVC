package Controllers

import (
	"Silver/Silver/Utils"
)

func Hello(request Utils.Request, response Utils.Response) {
	response.SendView("view.html")
}
