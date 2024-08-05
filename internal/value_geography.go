package internal

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// TODO: implement other geography types
// multipoint
// linestring
// ...
type geographyType interface {
	ToString() (string, error)
}

type GeographyValue struct {
	g geographyType
}

func (g *GeographyValue) Add(_ Value) (Value, error) {
	return nil, fmt.Errorf("unsupported add operator for geography value")
}

func (g *GeographyValue) Sub(_ Value) (Value, error) {
	return nil, fmt.Errorf("unsupported sub operator for geography value")
}

func (g *GeographyValue) Mul(_ Value) (Value, error) {
	return nil, fmt.Errorf("unsupported mul operator for geography value")
}

func (g *GeographyValue) Div(_ Value) (Value, error) {
	return nil, fmt.Errorf("unsupported div operator for geography value")
}

func (g *GeographyValue) EQ(_ Value) (bool, error) {
	return false, fmt.Errorf("unsupported eq operator for geography value")
}

func (g *GeographyValue) GT(_ Value) (bool, error) {
	return false, fmt.Errorf("unsupported gt operator for geography value")
}

func (g *GeographyValue) GTE(_ Value) (bool, error) {
	return false, fmt.Errorf("unsupported gte operator for geography value")
}

func (g *GeographyValue) LT(_ Value) (bool, error) {
	return false, fmt.Errorf("unsupported lt operator for geography value")
}

func (g *GeographyValue) LTE(_ Value) (bool, error) {
	return false, fmt.Errorf("unsupported lte operator for geography value")
}

func (g *GeographyValue) ToInt64() (int64, error) {
	return 0, fmt.Errorf("unsupported ToInt64 operator for geography value")
}

func (g *GeographyValue) ToString() (string, error) {
	return g.g.ToString()
}

func (g *GeographyValue) ToBytes() ([]byte, error) {
	v, err := g.ToString()
	if err != nil {
		return nil, err
	}

	return []byte(v), nil
}

func (g *GeographyValue) ToFloat64() (float64, error) {
	return 0, fmt.Errorf("unsupported ToFloat64 operator for geography value")
}

func (g *GeographyValue) ToBool() (bool, error) {
	return false, fmt.Errorf("unsupported ToBool operator for geography value")
}

func (g *GeographyValue) ToArray() (*ArrayValue, error) {
	return nil, fmt.Errorf("unsupported ToArray operator for geography value")
}

func (g *GeographyValue) ToStruct() (*StructValue, error) {
	return nil, fmt.Errorf("unsupported ToStruct operator for geography value")
}

func (g *GeographyValue) ToJSON() (string, error) {
	s, err := g.ToString()
	if err != nil {
		return "", err
	}

	return strconv.Quote(s), nil
}

func (g *GeographyValue) ToTime() (time.Time, error) {
	return time.Time{}, fmt.Errorf("unsupported ToTime operator for geography value")
}

func (g *GeographyValue) ToRat() (*big.Rat, error) {
	return nil, fmt.Errorf("unsupported ToRat operator for geography value")
}

func (g *GeographyValue) Format(verb rune) string {
	str, err := g.ToString()
	if err != nil {
		return "error"
	}

	switch verb {
	case 't':
		return str
	case 'T':
		return fmt.Sprintf(`GEOGRAPHY %q`, str)
	}

	return str
}

func (g *GeographyValue) Interface() interface{} {
	s, err := g.ToString()
	if err != nil {
		return nil
	}

	return s
}

type geographyPoint struct {
	longitude float64
	latitude  float64
}

func (g *geographyPoint) ToString() (string, error) {
	return fmt.Sprintf("POINT(%f %f)", g.longitude, g.latitude), nil
}
