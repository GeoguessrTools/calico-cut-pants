package dummy

import (
	"github.com/GeoguessrTools/calico-cut-pants/core"
	util "github.com/GeoguessrTools/calico-cut-pants/utilities"
)

var (
	argentina = core.Country{
		Name:         "Argentina",
		FlagURL:      "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-flag.png",
		MapURL:       "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-map.png",
		Continent:    core.SouthAmerica,
		Capital:      "Buenos Aires",
		Languages:    []core.Language{core.Spanish},
		PhoneNumbers: []string{"+54"},
		Notes: []string{
			"Argentina is a big country and a big part of it is covered by Street View",
			"Argentina is a very flat country except for the western border with Chile and the south around Ushuaia",
			"The southern part lacks any real vegetation, exception is the southern part near Ushuaia",
			"The middle line often consists of a dashed white line next to a yellow, continuous line",
			"Some Argentinian license plates have a distinctive black point in the middle of the plate",
		},
		RightDrive:     true,
		TopLevelDomain: ".ar",
		Meta: &core.Meta{
			Car: []core.Hint{
				{
					ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-meta-1.png"),
					Description: util.StrPtr("The google car in Argentina is black. You can see it by panning down."),
				},
			},
		},
		LicensePlates: []core.Hint{
			{
				ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-license-plate-new.png"),
			},
			{
				ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-license-plate-old.png"),
			},
		},
		BollardsNSuch: []core.Hint{

			{
				ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-gas-station-sign.png"),
				Description: util.StrPtr("This is the most common gas station in the country"),
			},
			{
				ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-highway-marker.png"),
				Description: util.StrPtr("Highway marker"),
			},
			{
				ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-bollards.png"),
				Description: util.StrPtr("Bollards"),
			},
		},
		Signs: &core.Signs{
			Yield: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-yield-1.png"),
				},
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-yield-2.png"),
				},
			},
			Stop: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-stop.png"),
				},
			},
			Pedestrian: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-pedestrian.png"),
				},
			},
			Direction: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-direction.png"),
				},
			},
			Speed: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-speed.png"),
				},
			},
			Chevrons: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-chevron-blue.png"),
				},
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-chevron-left-red.png"),
				},
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-chevron-right-red.png"),
				},
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-chevron-yellow-black.png"),
				},
			},
			Highway: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-highway.png"),
				},
			},
			Street: []core.Hint{
				{
					ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-street-name-sign.png"),
					Description: util.StrPtr("Street name signs are dark and perpendicular"),
				},
			},
			Transit: []core.Hint{
				{
					ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-bus-sign.png"),
					Description: util.StrPtr("Bus stop sign"),
				},
			},
		},
		Road: &core.Road{
			StreetLines: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-street-line-1.png"),
				},
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-street-line-2.png"),
				},
			},
		},
		AreaCodes: []core.AreaCode{
			{
				Code:     "11",
				Location: "Buenos Aires",
			},
			{
				Code:     "264",
				Location: "San Juan",
			},
			{
				Code:     "341",
				Location: "Rosario",
			},
			{
				Code:     "342",
				Location: "Santa Fe",
			},
			{
				Code:     "351",
				Location: "Cordoba",
			},
		},
		UtilityPoles: &core.UtilityPoles{
			ExtraNotes: []string{},
			Poles: []core.Hint{
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-utility-pole-1.png"),
				},
				{
					ImageURL: util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-utility-pole-2.png"),
				},
				{
					ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-utility-pole-blue.png"),
					Description: util.StrPtr("This pole can normally be found in Ushuaia, the southernmost point of Argentina."),
				},
			},
		},
		Misc: []core.Hint{
			{
				ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-misc-1.png"),
				Description: util.StrPtr("Most parts of Argentina are covered in grassy field. Picture in the area south of Rosario/Buenos Ares."),
			},
			{
				ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-misc-2.png"),
				Description: util.StrPtr("The southern most point of Argentina near Ushuaia. Alpine feeling, a lot of bushy green and snow."),
			},
			{
				ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-misc-3.png"),
				Description: util.StrPtr("Dry soil and patches of grass and low bushes. Mostly in the southern part of Argentina, picture from Sierra Nevada."),
			},
			{
				ImageURL:    util.StrPtr("https://calico-cut-geo-tools.s3.amazonaws.com/argentina-misc-4.png"),
				Description: util.StrPtr("The north of Argentina. A lot of green vegetation."),
			},
		},
	}

	storedCountries = map[string]core.Country{
		"argentina": argentina,
	}
)
