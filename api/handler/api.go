package handler

import (
	"fmt"
	"log"
	"regexp"

	"github.com/exuan/waka-api/api/message"
	"github.com/exuan/waka-api/validator"
	"github.com/labstack/echo/v4"
)

func ApiHeartbeat(c echo.Context) error {
	vhs := make([]*validator.Heartbeat, 0)
	if err := c.Bind(&vhs); err != nil {
		return err
	}
	machine := c.Request().Header.Get("X-Machine-Name")
	tz := c.Request().Header.Get("Timezone")
	ip := c.RealIP()
	log.Println(ParseUserAgent(c.Request().Header.Get("User-Agent")))
	for _, v := range vhs {
		v.Timezone = tz
		v.Machine = machine
		v.Ip = ip
		log.Println(v)
	}

	return c.JSON(message.OK, nil)
}
func ParseUserAgent(ua string) (string, string, error) {
	re := regexp.MustCompile(`(?iU)^wakatime\/[\d+.]+\s\((\w+)-.*\)\s.+\s([^\/\s]+)-wakatime\/.+$`)
	groups := re.FindAllStringSubmatch(ua, -1)

	log.Println(groups)
	if len(groups) == 0 || len(groups[0]) != 3 {
		return "", "", fmt.Errorf("failed to parse user agent string")
	}
	log.Println(groups[0][0])
	return groups[0][1], groups[0][2], nil
}
