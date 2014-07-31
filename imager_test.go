package avatar

import "testing"

func TestIdenticonCreate(t *testing.T) {
	var identicon = Identicon{
		avatar:          Avatar{emailAddress: "john@example.net"},
		filename:        "foo.png",
		borderSize:      15,
		squareSize:      25,
		gridSize:        7,
		backgroundColor: 0,
	}

	if ok := identicon.Create(); !ok {
		t.Error("Some shit just went down!")
	}
}
