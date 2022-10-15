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

	newZealand = core.Country{
		Name:         "New Zealand",
		FlagURL:      "https://cdn.britannica.com/17/3017-004-DCC13F9D/Flag-New-Zealand.jpg",
		MapURL:       "https://outthere.kiwi/wp-content/uploads/New-Zealand-largest-cities-map-v2-919x1024.jpg",
		Continent:    core.Oceania,
		Capital:      "Wellington",
		Languages:    []core.Language{core.English, core.Maori},
		PhoneNumbers: []string{"+64"},
		Notes: []string{
			"https://en.wikipedia.org/wiki/Road_signs_in_New_Zealand#Location_Referencing_Management_System_(LRMS)",
		},
		RightDrive:        false,
		TopLevelDomain:    ".nz",
		PhoneNumberFormat: "+64-A-BBBB-BBBB",
		Meta: &core.Meta{
			Car: []core.Hint{
				{
					ImageURL:    util.StrPtr("https://www.geohints.com/Sources/Cars/car_378.jpg"),
					Description: util.StrPtr("Car with red stripes on the front that says Highlands"),
				},
				{
					ImageURL:    util.StrPtr("https://www.geohints.com/Sources/Cars/car_528.jpg"),
					Description: util.StrPtr("Blue front of car displaying"),
				},
			},
		},
		LicensePlates: []core.Hint{
			{
				ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/7/71/NEW_ZEALAND_2006_-LICENSE_PLATE%2C_AAA-000_SERIES_SHORT_DIES_-_Flickr_-_woody1778a.jpg"),
				Description: util.StrPtr("License plates issued in 2006"),
			},
			{
				ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/7/7c/New_zealand_license_plate_gg.jpg"),
				Description: util.StrPtr("License plates issued between 1973 and 2006"),
			},
			{
				ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/9/9c/2006_Ferrari_F430_F1._FZ200_%2816819988608%29_%28cropped%29.jpg"),
				Description: util.StrPtr("Europlate style"),
			},
		},
		BollardsNSuch: []core.Hint{
			{
				ImageURL:          util.StrPtr("https://geohints.com/Sources/Bollards/bollard_218.jpg"),
				GoogleMapsExample: util.StrPtr("https://www.google.com/maps?q&layer=c&cbll=-39.3099342,175.7364567"),
			},
			{
				ImageURL:          util.StrPtr("https://geohints.com/Sources/Bollards/bollard_394.jpg"),
				GoogleMapsExample: util.StrPtr("https://goo.gl/maps/KZCLGzkbtNskXSkQ9"),
			},
			{
				ImageURL:          util.StrPtr("https://geohints.com/Sources/Bollards/bollard_301.jpg"),
				GoogleMapsExample: util.StrPtr("https://maps.google.com/maps?q=&layer=c&cbll=-45.1745633,170.3869734"),
				Tags:              []string{"red", "yellow", "bollard", "dashes"},
			},
			{
				ImageURL:          util.StrPtr("https://geohints.com/Sources/Bollards/bollard_241.jpg"),
				GoogleMapsExample: util.StrPtr("https://maps.google.com/maps?q=&layer=c&cbll=-45.9822522,168.7499536"),
				Tags:              []string{"green", "white", "bollard", "mile marker"},
			},
		},
		Signs: &core.Signs{
			Yield: []core.Hint{
				{
					ImageURL: util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/5/55/New_Zealand_road_sign_R2-2.svg"),
				},
				{
					ImageURL: util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/thumb/7/7a/New_Zealand_road_sign_R2-2F.svg/675px-New_Zealand_road_sign_R2-2F.svg.png"),
				},
			},
			Stop: []core.Hint{
				{
					ImageURL: util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/thumb/2/25/New_Zealand_road_sign_R2-1.svg/600px-New_Zealand_road_sign_R2-1.svg.png"),
				},
				{
					ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/thumb/5/56/New_Zealand_road_sign_R2-4.svg/600px-New_Zealand_road_sign_R2-4.svg.png"),
					Description: util.StrPtr("Stop sign for school areas"),
				},
				{
					ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/New_Zealand_road_sign_R2-4_%28perforated%29.svg/600px-New_Zealand_road_sign_R2-4_%28perforated%29.svg.png"),
					Description: util.StrPtr("Alternative perforated top sign for school areas"),
				},
			},
			Pedestrian: []core.Hint{
				{
					ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/thumb/8/86/New_Zealand_road_sign_W16-1-FYG.svg/1024px-New_Zealand_road_sign_W16-1-FYG.svg.png"),
					Description: util.StrPtr("Standard pedestrian crossing sign"),
				},
				{
					ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/thumb/c/cf/New_Zealand_road_sign_W16-2-FYG.svg/600px-New_Zealand_road_sign_W16-2-FYG.svg.png"),
					Description: util.StrPtr("Cross walk coming up sign"),
				},
			},
			Speed: []core.Hint{
				{
					ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/New_Zealand_road_sign_R1-1_%2860%29.svg/1024px-New_Zealand_road_sign_R1-1_%2860%29.svg.png"),
					Description: util.StrPtr("Standard speed sign"),
				},
				{
					ImageURL:    util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/thumb/7/7d/New_Zealand_road_sign_R1-2_%2860%29.svg/1024px-New_Zealand_road_sign_R1-2_%2860%29.svg.png"),
					Description: util.StrPtr("Variable speed sign. Can also have black background"),
				},
				{
					ImageURL:           util.StrPtr("https://upload.wikimedia.org/wikipedia/commons/9/99/Products-ez-szsk144-rs62-w16-5-1-kura-school-speed-30km-sign-thumb.jpg"),
					Description:        util.StrPtr("Maori Te Kura school zone sign"),
					Significance:       util.StrPtr("There is only one Te Kura school in the South of NZ. The rest are in the North."),
					RegionSpecificity:  util.Float32Prt{0.4},
					CountrySpecificity: util.Float32Prt{1},
				},
			},
		},
		AreaCodes: []core.AreaCode{
			{
				Code:     "9",
				Location: "Auckland/North Island - North/Pokeno/Whangarei",
			},
			{
				Code:     "3",
				Location: "Christchurch/Dunedin/Invercargill/Nelson/South Island/Timaru",
			},
			{
				Code:     "7",
				Location: "Hamilton/North Island - Central and East/Rotorua/Tauranga",
			},
			{
				Code:     "6",
				Location: "Hastings/Napier/New Plymouth/North Island - West/Palmerston North/Wanganui",
			},
			{
				Code:     "24",
				Location: "New Zealand External Territories",
			},
			{
				Code:     "4",
				Location: "Wellington Region",
			},
		},
		UtilityPoles: &core.UtilityPoles{
			Poles: []core.Hint{
				{
					ImageURL:          util.StrPtr("https://geohints.com/Sources/Poles/pole_37.jpg"),
					Description:       util.StrPtr("Has a depression and white bottom"),
					Tags:              []string{"concrete", "grey"},
					GoogleMapsExample: util.StrPtr("https://maps.google.com/maps?q=&layer=c&cbll=-40.3220863,175.8779131"),
				},
				{
					ImageURL:          util.StrPtr("https://geohints.com/Sources/Poles/pole_57.jpg"),
					Description:       util.StrPtr("Dark concrete with a posseum guard"),
					Tags:              []string{"posseum guard", "concrete"},
					GoogleMapsExample: util.StrPtr("https://maps.google.com/maps?q=&layer=c&cbll=-37.8128412,175.4398179"),
				},
				{
					ImageURL:          util.StrPtr("https://geohints.com/Sources/Poles/pole_147.jpg"),
					Description:       util.StrPtr("Double pole"),
					Tags:              []string{"double pole"},
					GoogleMapsExample: util.StrPtr("https://goo.gl/maps/3BkpvdYEMVFCBmzH6"),
				},
			},
		},
	}

	storedCountries = map[string]core.Country{
		"argentina":   argentina,
		"new-zealand": newZealand,
	}
)
