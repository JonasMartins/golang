package main

import (
	"fmt"
	"reflect"
)

type Chat struct {
	Id       int
	Messages []*Message
}

type Message struct {
	Id     int
	Body   string
	ChatId int
}

func MapGenericRelation[T, S any](x []*T, x_field string, xy_relation string, y []*S, y_field string) {

	var rx, ry, rxy reflect.Value
	for _, _x := range x {
		rx = reflect.ValueOf(_x).Elem()
		rxy = rx.FieldByName(xy_relation)
		for _, _y := range y {
			ry = reflect.ValueOf(_y).Elem()
			if rx.FieldByName(x_field).Interface() == ry.FieldByName(y_field).Interface() {
				rxy = reflect.Append(rxy, reflect.ValueOf(_y))
			}
		}
		rx.FieldByName(xy_relation).Set(rxy)
	}
}

func generics() {
	var c1 = Chat{1, []*Message{}}
	var c2 = Chat{2, []*Message{}}
	var c3 = Chat{3, []*Message{}}

	var m1 = Message{1, "test", 1}
	var m2 = Message{2, "2", 1}
	var m3 = Message{3, "3", 2}
	var m4 = Message{4, "4", 3}

	chats := []*Chat{&c1, &c2, &c3}
	messages := []*Message{&m1, &m2, &m3, &m4}
	MapGenericRelation(chats, "Id", "Messages", messages, "ChatId")
	fmt.Println(*chats[0].Messages[1])

}

func main() {
	Break()
}
