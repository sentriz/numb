package unit

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/unit/dimension"
)

type unitEntry struct {
	Unit Unit
	Exp  int
}

type UnitList []unitEntry

func (u1 UnitList) Conforms(u2 UnitList) bool {
	return u1.Dimension().Equals(u2.Dimension())
}

func (u1 UnitList) IsSingle() bool {
	return len(u1) == 1 && u1[0].Exp == 1
}

func (u UnitList) Dimension() dimension.Dimensions {
	var d dimension.Dimensions
	for _, x := range u {
		d = d.Add(x.Unit.dimension.Dims.Exp(x.Exp))
	}
	return d
}

func (units UnitList) normalize(n *big.Rat) *big.Rat {
	if len(units) == 0 {
		return n
	}
	num := new(big.Rat).Set(n)
	if units.IsSingle() {
		// (n + u.offset) * u.value
		u := units[0].Unit
		if u.offset != nil {
			num.Add(num, u.offset)
		}
		num = num.Mul(num, u.value)
	} else {
		for _, u := range units {
			num.Mul(num, ratexp(u.Unit.value, u.Exp))
		}
	}
	return num
}

func (units UnitList) denormalize(n *big.Rat) *big.Rat {
	if len(units) == 0 {
		return n
	}
	num := new(big.Rat).Set(n)
	if units.IsSingle() {
		// (n / u.value) - u.offset
		u := units[0].Unit
		num.Quo(num, u.value)
		if u.offset != nil {
			num.Sub(num, u.offset)
		}
	} else {
		for _, u := range units {
			num.Quo(num, ratexp(u.Unit.value, u.Exp))
		}
	}
	return num
}

func (u UnitList) String() string {
	if len(u) == 2 && u[0].Exp == 1 && u[1].Exp == -1 {
		return u[0].Unit.String() + "/" + u[1].Unit.String()
	}

	b := make([]string, 0, len(u))
	for _, entry := range u {
		if entry.Exp == 0 {
			continue
		} else if entry.Exp == 1 {
			b = append(b, entry.Unit.String())
		} else {
			b = append(b, fmt.Sprintf("%s^%d", entry.Unit.String(), entry.Exp))
		}
	}
	return strings.Join(b, " ")
}

func (u1 UnitList) Exp(x int) UnitList {
	u2 := UnitList{}
	for _, u := range u1 {
		u2 = append(u2, unitEntry{
			Unit: u.Unit,
			Exp: u.Exp * x,
		})
	}
	return u2
}

func (this UnitList) Mul(other UnitList) UnitList {
	c := UnitList{}
	for _, u := range this {
		c = append(c, u)
	}
	for _, u := range other {
		c = append(c, u)
	}
	return c
}

func (this UnitList) Quo(other UnitList) UnitList {
	c := UnitList{}
	for _, u := range this {
		c = append(c, u)
	}
	for _, u := range other {
		c = append(c, unitEntry{
			Unit: u.Unit,
			Exp:  -u.Exp,
		})
	}
	return c
}
