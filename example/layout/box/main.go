package main

import (
	"ChromaTerm/pkg/components/element"
)

func main() {
	box1 := element.NewBox(20, 4)
	box2 := element.NewBox(20, 4)
	box3 := element.NewBox(20, 4)
	engine := element.NewEngine()

	box1.Position(0, 0)
	box2.Position(22, 0)
	box3.Position(44, 0)
	box1.SetTitle("Left", element.BoxTitleLayoutTopLeft)
	box2.SetTitle("Center", element.BoxTitleLayoutTopCenter)
	box3.SetTitle("Right", element.BoxTitleLayoutTopRight)
	engine.AddElement(box1)
	engine.AddElement(box2)
	engine.AddElement(box3)

	engine.Refresh()
}
