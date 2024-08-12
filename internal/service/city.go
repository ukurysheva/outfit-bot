package service

import "fmt"

var cityInfoTmpl = "Ваш город - %s."

func (s *Service) CityInfo(cityTitle string) string {
	return fmt.Sprintf(cityInfoTmpl, cityTitle)
}
