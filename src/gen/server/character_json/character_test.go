package character_json

import (
	"bytes"
	"encoding/json"
	"gen/types"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	rand.Seed(42)

	rawjson, err := ioutil.ReadFile("testdata/partial_character.json")
	die(err)
	partialChar := types.Character{}

	err = json.Unmarshal(rawjson, &partialChar)
	die(err)

	ts := httptest.NewServer(http.HandlerFunc(CreateCharacter))
	defer ts.Close()
	res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(rawjson))
	die(err)
	defer res.Body.Close()

	responseChar := types.Character{}

	err = json.NewDecoder(res.Body).Decode(&responseChar)
	die(err)

	assertEq(t, "Kathleen", responseChar.Name)
	assertEq(t, "Blair", responseChar.Surname)
	assertEq(t, "Half-orc", responseChar.Race.Name)
	assertEq(t, "Female", responseChar.Sex)
	assertEq(t, "Osirion", responseChar.Nationality)

	assertEq(t, "Beulah", responseChar.Mother.Name)
	assertEq(t, "Murphy", responseChar.Mother.Surname)
	assertEq(t, "Human", responseChar.Mother.Race.Name)

	assertNeq(t, "", responseChar.Father.Name)
	assertNeq(t, "", responseChar.Father.Surname)
	if !(responseChar.Father.Race.Name == "Orc" ||
		responseChar.Father.Race.Name == "Half-orc") {
		t.Errorf(
			"Expected fathers race to be orc or half-orc, got '%s'",
			responseChar.Father.Race.Name,
		)
	}

	assertEq(t, int(9), responseChar.Stats.Strength)
	assertNeq(t, int(0), responseChar.Stats.Dexterity)
	assertEq(t, int(13), responseChar.Stats.Constitution)
	assertEq(t, int(9), responseChar.Stats.Intelligence)
	assertNeq(t, int(0), responseChar.Stats.Wisdom)
	assertEq(t, int(14), responseChar.Stats.Charisma)
}

func assertEq(t *testing.T, v1 interface{}, v2 interface{}) {
	if v1 != v2 {
		t.Errorf("Expected %+v (%T) to equal %+v (%T)", v1, v1, v2, v2)
	}
}
func assertNeq(t *testing.T, v1 interface{}, v2 interface{}) {
	if v1 == v2 {
		t.Errorf("Expected %+v (%T) not to equal %+v (%T)", v1, v1, v2, v2)
	}
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
