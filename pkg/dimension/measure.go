package dimension

var Measures = map[string]Dimensions{
	"dimensionless/angle":               Dimensions{},
	"length":                            Dimensions{Length: 1},
	"temperature":                       Dimensions{Temperature: 1},
	"area":                              Dimensions{Length: 2},
	"volume":                            Dimensions{Length: 3},
	"mass":                              Dimensions{Mass: 1},
	"time":                              Dimensions{Time: 1},
	"digital":                           Dimensions{Digital: 1},
	"currency":                          Dimensions{Currency: 1},
	"frequency/radioactivity":           Dimensions{Time: -1},
	"electric current":                  Dimensions{Current: 1},
	"amount of substance":               Dimensions{AmountOfSubstance: 1},
	"power":                             Dimensions{Mass: 1, Length: 2, Time: -3},
	"force":                             Dimensions{Mass: 1, Length: 1, Time: -2},
	"energy":                            Dimensions{Mass: 1, Length: 2, Time: -2},
	"electric charge":                   Dimensions{Time: 1, Current: 1},
	"electric potential":                Dimensions{Mass: 1, Length: 2, Time: -3, Current: -1},
	"electric capacitance":              Dimensions{Mass: -1, Length: -2, Time: 4, Current: 2},
	"electric conductance":              Dimensions{Mass: -1, Length: -2, Time: 3, Current: 2},
	"magnetic flux":                     Dimensions{Mass: 1, Length: 2, Time: -2, Current: -1},
	"magnetic flux density":             Dimensions{Mass: 1, Time: -2, Current: -1},
	"electric inductance":               Dimensions{Mass: 1, Length: 2, Time: -2, Current: -2},
	"electric resistance":               Dimensions{Mass: 1, Length: 2, Time: -3, Current: 2},
	"pressure":                          Dimensions{Mass: 1, Length: -1, Time: -2},
	"ionizing radiation/radiation dose": Dimensions{Length: 2, Time: -2},
	"catalyctic activity":               Dimensions{AmountOfSubstance: 1, Time: -1},
	"luminous flux/intensity":           Dimensions{LuminousIntensity: 1},
	"illuminance":                       Dimensions{LuminousIntensity: 1, Length: -1},
	"speed":                             Dimensions{Length: 1, Time: -1},
	"data rate":                         Dimensions{Digital: 1, Time: -1},
	"density":                           Dimensions{Mass: 1, Length: -3},
	"flow":                              Dimensions{Length: 3, Time: -1},
	"acceleration":                      Dimensions{Length: 1, Time: -2},
}
