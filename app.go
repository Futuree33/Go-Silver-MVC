package main

import (
	"Silver/Controllers"
	"Silver/Silver"
)

func main() {
	silver := Silver.App{}

	silver.AddRoute(Silver.Get, "/", Controllers.Hello)

	silver.Listen(8080)
}
