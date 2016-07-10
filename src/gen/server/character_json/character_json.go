package character_json

import (
	"encoding/json"
	"gen/types"
	"io/ioutil"
	"net/http"
)

func CreateCharacter(rw http.ResponseWriter, r *http.Request) {
	c := types.Character{}
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &c)
	c = c.FillInTheBlanks()
	err = json.NewEncoder(rw).Encode(c)
	if err != nil {
		panic(err)
	}
}
