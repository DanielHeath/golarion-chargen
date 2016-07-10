class App extends React.Component {

  constructor() {
    super()
    this.state = {
      count: 0
    }
  };

  componentWillMount() {
    fetch("/character.json")
      .then((resp) => resp.json())
      .then((character) => this.setState({character: character}))
  }


      // var pre = document.getElementById("json")
      // var form = document.getElementById("form")

      // var onFetch = (promise) => {
      //   promise.then(
      //     (resp) => resp.json()
      //   ).then((json) =>
      //     pre.innerText = JSON.stringify(json, null, 4)
      //   )
      // }
      // onFetch(fetch("/character.json"))

      // form.onsubmit = () => {
      //   var char = JSON.parse(pre.innerText)
      //   console.log(delete char.Father)

      //   onFetch(fetch(
      //     "/character.json",
      //     {method: "POST", body: JSON.stringify(char)}
      //   ));
      //   return false;
      // }

  render() {
    if (!this.state.character) {
      return (
        <div>
          Loading...
        </div>
      )
    }

    var Deletable = (props) => (
      <div>
        <button
          onClick={props.onDelete}
          className="delete-button"
        >
          <img
            src="/vendor/izzy-cross.svg"
            alt="delete"
          />
        </button>
        {props.propName}: {props.children}
      </div>
    );

    return (
      <div>
        <button onClick={this.handleClick.bind(this)}>
          Click me! Number of clicks: {this.state.count}
        </button>
        <pre>{JSON.stringify(this.state.character)}</pre>
        <Deletable
          propName="Name"
          onDelete={() => {
            this.deleteCharacterProp('Name')
            this.deleteCharacterProp('Surname')
          }}
        >
          <span className="first-name">
            {this.state.character.Name}
          </span>
          {' '}
          <span className="surname">
            {this.state.character.Surname}
          </span>
        </Deletable>

        <dl>

            <dt>Name</dt>
            <dd>
            </dd>

            <dt>Race</dt>
            <dd>
              <span className="race">
                {this.state.character.Race}
              </span>
            </dd>

            <dt>Sex</dt>
            <dd>
              <span className="sex">
                {this.state.character.Sex}
              </span>
            </dd>
        </dl>
      </div>
    );
  }

  handleClick() {
    this.setState({
      count: this.state.count + 1,
    });
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('app')
);

var sample = {
  "Race": "Half-orc",
  "Sex": "Female",
  "Nationality": "Ustalav",
  "Mother": {
    "Name": "Julia",
    "Surname": "Tate",
    "Race": "Half-orc",
    "Occupation": "Knight",
    "Nationality": "Ustalav"
  },
  "Father": {
    "Name": "Edward",
    "Surname": "Hoffman",
    "Race": "Human",
    "Occupation": "Artist",
    "Nationality": "Osirion"
  },
  "Stats": {
    "Strength": 14,
    "Dexterity": 12,
    "Constitution": 12,
    "Wisdom": 10,
    "Intelligence": 11,
    "Charisma": 12
  }
}
