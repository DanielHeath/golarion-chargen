package main

import (
	"bytes"
	"gen/types"
	"html/template"
	"io/ioutil"
	"math/rand"
	"time"

	dom "github.com/dominikh/go-js-dom"
	"github.com/gopherjs/gopherjs/js"
)

var character = types.Character{}
var tmpl = template.Must(template.New("main").Parse(htmlTemplate))

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	character = character.FillInTheBlanks()
	js.Global.Set("repickNationality", repickNationality)
	js.Global.Set("repickMum", repickMum)
	js.Global.Set("repickDad", repickDad)
	js.Global.Set("repickRace", repickRace)
	js.Global.Set("repickInfancy", repickInfancy)
	js.Global.Set("repickChildhood", repickChildhood)
	js.Global.Set("repickDeity", repickDeity)

	go rerender()
}

func rerender() {
	body := dom.GetWindow().Document().QuerySelector("body")
	body.SetInnerHTML("")
	str, err := runTemplate(tmpl, character)
	if err != nil {
		log(err.Error())
	} else {
		body.SetInnerHTML(str)
	}
}

func repickNationality() {
	oldNationality := character.Nationality
	character.Nationality = ""
	character.Father.Nationality = ""
	character.Mother.Nationality = ""
	character = character.FillInTheBlanks()
	if character.Nationality == oldNationality {
		repickNationality()
	} else {
		character.SpentFatePoints++
		go rerender()
	}
}

func repickDad() {
	character.Father = types.Parent{}
	character.SpentFatePoints++
	character = character.FillInTheBlanks()
	go rerender()
}

func repickMum() {
	character.Mother = types.Parent{}
	character.SpentFatePoints++
	character = character.FillInTheBlanks()
	go rerender()
}

func repickDeity() {
	old := character.Deity
	character.Deity = ""
	character = character.FillInTheBlanks()
	if character.Deity == old {
		repickDeity()
	} else {
		character.SpentFatePoints++
		go rerender()
	}
}

func repickChildhood() {
	old := character.Childhood
	character.Childhood = types.Childhood{}
	character = character.FillInTheBlanks()
	if character.Childhood == old {
		repickChildhood()
	} else {
		character.SpentFatePoints++
		go rerender()
	}
}

func repickInfancy() {
	oldInfancy := character.Infancy
	character.Infancy = types.Infancy{}
	character = character.FillInTheBlanks()
	if character.Infancy == oldInfancy {
		repickInfancy()
	} else {
		character.SpentFatePoints++
		go rerender()
	}
}

func repickRace() {
	oldRace := character.Race.Name
	character.Race = types.Race{}
	character.Name = ""
	character.Surname = ""
	character.Mother.Race = types.Race{}
	character.Mother.Name = ""
	character.Mother.Surname = ""
	character.Father.Race = types.Race{}
	character.Father.Name = ""
	character.Father.Surname = ""
	character = character.FillInTheBlanks()
	if character.Race.Name == oldRace {
		repickRace()
	} else {
		character.SpentFatePoints++
		go rerender()
	}
}

const htmlTemplate = `
<p>
	{{.Name}} {{.Surname}}, 
	a {{.Sex}} {{.Race.Name}} 
	(<a href="#" onclick="repickRace()">Repick Race</a>) 
	from {{.Nationality}} 
	(<a href="#" onclick="repickNationality()">Repick Nationality</a>)
</p>
<p>
	<a href="#" onclick="repickDad()">Repick Dad</a>
	Father: {{.Father}}
</p>
<p>
	<a href="#" onclick="repickMum()">Repick Mum</a>
	Mother: {{.Mother}}</p>
<p>
<a href="#" onclick="repickInfancy()">Repick Infancy</a> 
After your birth, you were cared for by {{.Infancy}}.
</p>
<p>
<a href="#" onclick="repickChildhood()">Repick Childhood</a> 
As you grew older you spent lots of time {{ .Childhood.Activity }} {{.Childhood.Location}}.
</p>
<p>
<a href="#" onclick="repickDeity()">Repick Deity</a>
You have come to the service of {{ .Deity }}. 
</p>
<p>
Stat rolls:
<ul>
  <li>Strength: {{.Stats.Strength}}</li>
  <li>Dexterity: {{.Stats.Dexterity}}</li>
  <li>Constitution: {{.Stats.Constitution}}</li>
  <li>Wisdom: {{.Stats.Wisdom}}</li>
  <li>Intelligence: {{.Stats.Intelligence}}</li>
  <li>Charisma: {{.Stats.Charisma}}</li>
</ul>
</p>
<p>
You have {{.Stats.BaseFatePoints}} fate points to spend.<br/>
You have spent {{ .SpentFatePoints }} of them.
</p>
`

func runTemplate(t *template.Template, data interface{}) (string, error) {
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(&buf)
	return string(result), err
}

func log(args ...interface{}) {
	dom.GetWindow().Console().Call("log", args...)
}
