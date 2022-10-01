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
	// guaranteed fields
	FlagURL           string     `json:"flag_url"`
	MapURL            string     `json:"map_url"`
	Continent         Continent  `json:"continent"`
	Capital           string     `json:"capital"`
	Languages         []Language `json:"languages"`
	PhoneNumbers      []string   `json:"phone_numbers"`
	RightDrive        bool       `json:"right_drive"`
	TopLevelDomain    string     `json:"top_level_domain"`
	PhoneNumberFormat string     `json:"phone_number_format"`
	// non guaranteed
	Meta          *Meta         `json:"meta,omitempty"`
	Signs         *Signs        `json:"signs,omitempty"`
	AreaCodes     []AreaCode    `json:"area_codes,omitempty"`
	UtilityPoles  *UtilityPoles `json:"utility_poles,omitempty"`
	Misc          []Hint        `json:"misc,omitempty"`
	Vehicles      []Hint        `json:"vehicles,omitempty"`
	BollardsNSuch []Hint        `json:"bollards_barricades_snowpoles,omitempty"`
	LicensePlates []Hint        `json:"license_plates,omitempty"`
	Road          *Road         `json:"road,omitempty"`
	Notes         []string      `json:"notes,omitempty"`
	// LanguageIdentifiers maps each language to specific traits, words etc.
	// This is a placeholder and will likely require construction of a language specific struct.
	LanguageIdentifiers map[Language]Hint `json:"language_identifiers,omitempty"`
}

type Meta struct {
	Car    []Hint `json:"car,omitempty"`
	Sky    []Hint `json:"sky,omitempty"`
	Camera []Hint `json:"camera,omitempty"`
}

// Hint is the default representation of identifiable country features.
type Hint struct {
	// ImageURL points to an image of the hint.
	ImageURL *string `json:"image_url,omitempty"`
	// AltImageURL points to a different image. This could be a reverse side, a blurred version, etc.
	AltImageURL *string `json:"alt_image_url,omitempty"`
	// GoogleMapsExample is a link to where this hint can be found in the wild.
	GoogleMapsExample *string `json:"google_maps_example,omitempty"`
	// Description a visual description of the hint (what features of the image you're looking for).
	Description *string `json:"description,omitempty"`
	// Significance is what makes this hint useful (to what level does it tell you where you are).
	// Make sure this field has as much specifying information as possible.
	// This should be left blank for any hint that could be applied to multiple countries.
	Significance *string `json:"significance,omitempty"`
	// RegionSpecificity is a decimal from 0-1 representing how specific this hint is to a region.
	RegionSpecificity *float32 `json:"region_specificity,omitempty"`
	// CountrySpecificity is a decimal from 0-1 representing how specific this hint is to the country.
	CountrySpecificity *float32 `json:"country_specificity,omitempty"`
	// Frequency is a decimal from 0-1 representing how often this hint shows up.
	Frequency *float32 `json:"frequency,omitempty"`
	// Identifier specified what kind of thing it is. Ex: for a map, this might be COUNTRY_MAP, chevron might be BLUE_WHITE_CHEVRON.
	Identifier *string `json:"identifier,omitempty"`
	// Tags are extra bits of information for searching, filtering, etc
	Tags []string `json:"tags,omitempty,omitempty"`
}

type Signs struct {
	Yield      []Hint `json:"yield,omitempty"`
	Stop       []Hint `json:"stop,omitempty"`
	Pedestrian []Hint `json:"pedestrian,omitempty"`
	Direction  []Hint `json:"direction,omitempty"`
	Speed      []Hint `json:"speed,omitempty"`
	Highway    []Hint `json:"highway,omitempty"`
	Chevrons   []Hint `json:"chevrons,omitempty"`
	Street     []Hint `json:"street,omitempty"`
	Transit    []Hint `json:"transit,omitempty"`
}

type Road struct {
	StreetLines []Hint `json:"street_lines,omitempty"`
	Material    []Hint `json:"material,omitempty"`
	// Quality is a value 1-10 signifying how good or bad the roads are. 10 is good, 1 is bad.
	Quality *float32 `json:"quality,omitempty"`
}

type AreaCode struct {
	Code     string `json:"code"`
	Location string `json:"location"`
}

type UtilityPoles struct {
	Poles      []Hint   `json:"poles,omitempty"`
	ExtraNotes []string `json:"extra_notes,omitempty"`
}
