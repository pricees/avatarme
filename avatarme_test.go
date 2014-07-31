package avatar

import "testing"

func TestToHashMessage(t *testing.T) {
	var avatarTests = []struct {
		avatar Avatar
		out    string
	}{
		{
			Avatar{emailAddress: "john@example.net"},
			"john@example.net",
		},
		{
			Avatar{emailAddress: "john@example.net", ipAddress: "123"},
			"john@example.net123",
		},
		{
			Avatar{emailAddress: "john@example.net", ipAddress: "123", publicKey: "a"},
			"john@example.net123a",
		},
		{
			Avatar{ipAddress: "123", publicKey: "a"},
			"123a",
		},
	}

	for _, test := range avatarTests {
		avatar := test.avatar

		if test.out != avatar.hashMessage() {
			t.Errorf("%s should have equaled %s", test.out, avatar.hashMessage())
		}
	}
}

func TestIdenticalHash(t *testing.T) {
	var avatarTests = []struct {
		avatar1 Avatar
		avatar2 Avatar
	}{
		{
			Avatar{emailAddress: "john@example.net"},
			Avatar{emailAddress: "john@example.net"},
		},
		{
			Avatar{ipAddress: "187.0.1.123"},
			Avatar{ipAddress: "187.0.1.123"},
		},
	}

	for _, avatars := range avatarTests {
		avatar1 := avatars.avatar1
		avatar2 := avatars.avatar2

		if avatar1.Hash() != avatar2.Hash() {
			t.Errorf("%s should have equaled %s", avatar1.Hash(), avatar2.Hash())
		}
	}
}
