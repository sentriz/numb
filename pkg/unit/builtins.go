package unit

import (
	"math/big"

	r "github.com/nkanaev/numb/pkg/ratutils"
	d "github.com/nkanaev/numb/pkg/dimension"
)

type unit struct {
	u         d.Measure
	names     string
	value     *big.Rat
	offset    *big.Rat
}

var units = []unit{
	{u: d.LENGTH, names: "LENGTH", value: r.ONE},
	{u: d.MASS, names: "MASS", value: r.ONE},
	{u: d.TIME, names: "TIME", value: r.ONE},

	{u: d.TEMPERATURE, names: "K, kelvin", value: r.ONE},
	{u: d.TEMPERATURE, names: "°C, degC, celsius", value: r.ONE, offset: r.Must("273.15")},
	{u: d.TEMPERATURE, names: "°F, degF, fahrenheit", value: r.Must("10/18"), offset: r.Must("459.67")},

	{u: d.ANGLE, names: "ANGLE", value: r.ONE},
	{u: d.DIGITAL, names: "DIGITAL", value: r.ONE},
	{u: d.VOLUME, names: "VOLUME", value: r.ONE},
	{u: d.ENERGY, names: "ENERGY", value: r.ONE},
	{u: d.FORCE, names: "FORCE", value: r.ONE},
	{u: d.PRESSURE, names: "PRESSURE", value: r.ONE},
	{u: d.POWER, names: "POWER", value: r.ONE},

	{u: d.RADIOACTIVITY, names: "RADIOACTIVITY", value: r.ONE},
	{u: d.AMOUNT_OF_SUBSTANCE, names: "AMOUNT_OF_SUBSTANCE", value: r.ONE},
	{u: d.CATALYCTIC_ACTIVITY, names: "kat, katal", value: r.ONE},
	{u: d.ELECTRIC_CAPACITANCE, names: "F, farad", value: r.ONE},
	{u: d.ELECTRIC_CHARGE, names: "C, coulomb", value: r.ONE},
	{u: d.ELECTRIC_CONDUCTANCE, names: "S, siemens", value: r.ONE},
	{u: d.ELECTRIC_CURRENT, names: "A, amp, ampere", value: r.ONE},
	{u: d.ELECTRIC_INDUCTANCE, names: "H, henry", value: r.ONE},
	{u: d.ELECTRIC_POTENTIAL, names: "V, volt", value: r.ONE},
	{u: d.ELECTRIC_RESISTANCE, names: "ohm", value: r.ONE},
	{u: d.FREQUENCY, names: "Hz, hertz", value: r.ONE},
	{u: d.ILLUMINANCE, names: "lx, lux", value: r.ONE},
	{u: d.IONIZING_RADIATION, names: "Sv, sievert", value: r.ONE},
	{u: d.LUMINOUS_FLUX, names: "lm, lumen", value: r.ONE},
	{u: d.LUMINOUS_INTENSITY, names: "cd, candela", value: r.ONE},
	{u: d.MAGNETIC_FLUX, names: "Wb, weber", value: r.ONE},
	{u: d.MAGNETIC_FLUX_DENSITY, names: "T, tesla", value: r.ONE},
	{u: d.RADIATION_DOSE, names: "Gy, gray", value: r.ONE},
	{u: d.SOLID_ANGLE, names: "sr, steradian", value: r.ONE},

	{u: d.CURRENCY, names: "CURRENCY", value: r.ONE},
}

type prefix struct {
	names string
	value *big.Rat
}

var prefixes = []prefix{
	// SI
	{"d, deci", r.Exp(10, -1)},
	{"c, centi", r.Exp(10, -2)},
	{"m, milli", r.Exp(10, -3)},
	{"u, micro", r.Exp(10, -6)},
	{"n, nano", r.Exp(10, -9)},
	{"p, pico", r.Exp(10, -12)},
	{"f, femto", r.Exp(10, -15)},
	{"a, atto", r.Exp(10, -18)},
	{"z, zepto", r.Exp(10, -21)},
	{"y, yocto", r.Exp(10, -24)},
	// SI
	{"d, deca", r.Exp(10, 1)},
	{"h, hecto", r.Exp(10, 2)},
	{"k, kilo", r.Exp(10, 3)},
	{"M, mega", r.Exp(10, 6)},
	{"G, giga", r.Exp(10, 9)},
	{"T, tera", r.Exp(10, 12)},
	{"P, peta", r.Exp(10, 15)},
	{"E, exa", r.Exp(10, 18)},
	{"Z, zetta", r.Exp(10, 21)},
	{"Y, yotta", r.Exp(10, 24)},
	// Binary
	{"Ki, kibi", r.Exp(1024, 1)},
	{"Mi, mebi", r.Exp(1024, 2)},
	{"Gi, gibi", r.Exp(1024, 3)},
	{"Ti, tebi", r.Exp(1024, 4)},
	{"Pi, pebi", r.Exp(1024, 5)},
	{"Ei, exi", r.Exp(1024, 6)},
	{"Zi, zebi", r.Exp(1024, 7)},
	{"Yi, yobi", r.Exp(1024, 8)},
}
