package model

var (
	CityNameList = []string{"msc", "nnov", "spb", "nsk", "ekb", "kzn"}

	CityTitleByName = map[string]string{
		"msc":  "Москва",
		"nnov": "Нижний Новгород",
		"spb":  "Санкт-Петербург",
		"nsk":  "Новосибирск",
		"ekb":  "Екатеринбург",
		"kzn":  "Казань",
	}

	CityCordsByName = map[string][]float64{
		"msc":  {55.7540471, 37.620405},
		"nnov": {56.3240627, 44.0053913},
		"spb":  {59.9391313, 30.3159004},
		"nsk":  {55.028191, 82.9211489},
		"ekb":  {56.8385216, 60.6054911},
		"kzn":  {55.7943584, 49.1114975},
	}
)

func IsCityName(str string) bool {
	if _, ex := CityTitleByName[str]; ex {
		return true
	}

	return false
}

func IsChangeCityButton(str string) bool {
	return str == ChangeCityButtonName
}
