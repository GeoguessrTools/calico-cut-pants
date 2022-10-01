package core

type Continent string

const (
	Asia         Continent = "ASIA"
	NorthAmerica Continent = "NORTH_AMERICA"
	SouthAmerica Continent = "SOUTH_AMERICA"
	Africa       Continent = "AFRICA"
	Europe       Continent = "EUROPE"
	Australia    Continent = "AUSTRALIA"
)

type Language string

const (
	Spanish Language = "SPANISH"
)

type Country struct {
	FlagURL           string                 `json:"flag_url"`
	MapURL            string                 `json:"map_url"`
	Continent         Continent              `json:"continent"`
	Capital           string                 `json:"capital"`
	Languages         []Language             `json:"languages"`
	PhoneNumbers      []string               `json:"phone_numbers"`
	Notes             []string               `json:"notes"`
	RightDrive        bool                   `json:"right_drive"`
	Meta              Meta                   `json:"meta"`
	BollardsAndExtras BollardsAndExtras      `json:"bollards_and_extras"`
	SignageAndTraffic SignageAndTraffic      `json:"signage_and_traffic"`
	AreaCodes         []AreaCode             `json:"area_codes"`
	UtilityPoles      UtilityPoles           `json:"utility_poles"`
	Misc              []ImageWithDescription `json:"misc"`
}

type Meta struct {
	Car    []ImageWithDescription `json:"car"`
	Sky    []ImageWithDescription `json:"sky"`
	Camera []ImageWithDescription `json:"camera"`
}

type ImageWithDescription struct {
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type BollardsAndExtras struct {
	Notes []string               `json:"notes"`
	Info  []ImageWithDescription `json:"info"`
}

type SignageAndTraffic struct {
	YieldSigns      []ImageWithDescription `json:"yield_signs"`
	StopSigns       []ImageWithDescription `json:"stop_signs"`
	PedestrianSigns []ImageWithDescription `json:"pedestrian_signs"`
	DirectionSigns  []ImageWithDescription `json:"direction_signs"`
	SpeedSigns      []ImageWithDescription `json:"speed_signs"`
	Chevrons        []ImageWithDescription `json:"chevrons"`
	LicensePlates   []ImageWithDescription `json:"license_plates"`
	HighwaySigns    []ImageWithDescription `json:"highway_signs"`
	StreetLines     []ImageWithDescription `json:"street_lines"`
}

type AreaCode struct {
	Code     string `json:"code"`
	Location string `json:"location"`
}

type UtilityPoles struct {
	Poles      []ImageWithDescription `json:"poles"`
	ExtraNotes []string               `json:"extra_notes"`
}
