package helpers

import (
	"reflect"

	uuid "github.com/satori/go.uuid"
)

// recieves two reletational objetcs x and y
// the joins columns must be passed too, and the x join that
// will store the y entities like:
// x.xy_relation = append(x.xy_relation, y.Yfield)
// at the end the relatio is mapped
// the method is not prefectely generic because, the ID join field
// is of type UUID, so is needed to specify this, if the tables
// joins uses integers this cast will not be necessary
func MapGenericRelation[T, S any](x []*T, x_field string, xy_relation string, y []*S, y_field string) {
	var rx, ry, rxy reflect.Value
	for _, _x := range x {
		rx = reflect.ValueOf(_x).Elem()
		rxy = rx.FieldByName(xy_relation)
		for _, _y := range y {
			ry = reflect.ValueOf(_y).Elem()
			if rx.FieldByName(x_field).Interface().(*uuid.UUID).String() == ry.FieldByName(y_field).Interface().(*uuid.UUID).String() {
				rxy = reflect.Append(rxy, reflect.ValueOf(_y))
			}
		}
		rx.FieldByName(xy_relation).Set(rxy)
	}
}

func (h *Helpers) ContainsString(x []*string, y string) bool {
	for _, _x := range x {
		if *_x == y {
			return true
		}
	}
	return false
}
