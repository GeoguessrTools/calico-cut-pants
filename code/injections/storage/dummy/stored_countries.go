package dummy

import (
	"github.com/GeoguessrTools/calico-cut-pants/core"
)

var (
	argentina = core.Country{
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
		RightDrive: true,
		Meta: core.Meta{
			Car: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-meta-1.png",
					Description: "The google car in Argentina is black. You can see it by panning down.",
				},
			},
		},
		BollardsAndExtras: core.BollardsAndExtras{
			Notes: []string{
				"Argentina generally lacks bollards, it's not worth focusing on them",
			},
			Info: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-bus-sign.png",
					Description: "This sign is commonly seen at gas stations",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-street-name-sign.png",
					Description: "Street name signs are dark and perpendicular",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-gas-station-sign.png",
					Description: "Bus stop sign",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-highway-marker.png",
					Description: "Highway marker",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-bollards.png",
					Description: "Bollards",
				},
			},
		},
		SignageAndTraffic: core.SignageAndTraffic{
			YieldSigns: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-yield-1.png",
					Description: "",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-yield-2.png",
					Description: "",
				},
			},
			StopSigns: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-stop.png",
					Description: "",
				},
			},
			PedestrianSigns: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-pedestrian.png",
					Description: "",
				},
			},
			DirectionSigns: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-direction.png",
					Description: "",
				},
			},
			SpeedSigns: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-speed.png",
					Description: "",
				},
			},
			Chevrons: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-chevron-blue.png",
					Description: "",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-chevron-left-red.png",
					Description: "",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-chevron-right-red.png",
					Description: "",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-chevron-yellow-black.png",
					Description: "",
				},
			},
			LicensePlates: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-license-plate-new.png",
					Description: "New",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-license-plate-old.png",
					Description: "Old",
				},
			},
			HighwaySigns: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-highway.png",
					Description: "",
				},
			},
			StreetLines: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-street-line-1.png",
					Description: "",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-street-line-2.png",
					Description: "",
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
		UtilityPoles: core.UtilityPoles{
			ExtraNotes: []string{},
			Poles: []core.ImageWithDescription{
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-utility-pole-1.png",
					Description: "",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-utility-pole-2.png",
					Description: "",
				},
				{
					ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-utility-pole-blue.png",
					Description: "This pole can normally be found in Ushuaia, the southernmost point of Argentina.",
				},
			},
		},
		Misc: []core.ImageWithDescription{
			{
				ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-misc-1.png",
				Description: "Most parts of Argentina are covered in grassy field. Picture in the area south of Rosario/Buenos Ares.",
			},
			{
				ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-misc-2.png",
				Description: "The southern most point of Argentina near Ushuaia. Alpine feeling, a lot of bushy green and snow.",
			},
			{
				ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-misc-3.png",
				Description: "Dry soil and patches of grass and low bushes. Mostly in the southern part of Argentina, picture from Sierra Nevada.",
			},
			{
				ImageURL:    "https://calico-cut-geo-tools.s3.amazonaws.com/argentina-misc-4.png",
				Description: "The north of Argentina. A lot of green vegetation.",
			},
		},
	}

	storedCountries = map[string]core.Country{
		"argentina": argentina,
	}
)
