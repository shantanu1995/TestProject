import React, { Component } from "react";
class Exercise extends Component {
  constructor(props) {
    super(props);
    this.state = { username: "" };
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleSubmit(event) {
    fetch("http://localhost:3000/api/exercise/new-user", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        username: this.refs.username.value
      })
    })
      .then(res => res.json())
      .then(json => {
        alert(JSON.stringify(json));
      });
    event.preventDefault();
  }
  handleSubmit1(event) {
    fetch("http://localhost:3000/api/exercise/add", {
      method: "POST",

      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        id: this.refs.uid.value,
        description: this.refs.desc.value,
        duration: this.refs.dur.value,
        date: this.refs.dat.value
      })
    })
      .then(res => res.json())
      .then(json => {
        alert(JSON.stringify(json));
      });
    event.preventDefault();
  }

  handleEmailChange = e => {
    this.setState({ username: e.target.value });
  };
  state = {};
  render() {
    return (
      <div className="container">
        <h1>Exercise tracker</h1>
        <form
          name="usercreate"
          className="form-create-user"
          onSubmit={this.handleSubmit}
        >
          <h3>Create a New User</h3>
          <p>
            <code>POST /api/exercise/new-user</code>
          </p>
          <input
            type="text"
            ref="username"
            name="username"
            placeholder="username"
            onChange={this.handleChange}
          />
          <input type="submit" value="Submit" />
        </form>
        <form
          name="usercreate1"
          className="form-create-user1"
          onSubmit={this.handleSubmit1}
        >
          <h3>Add exercises</h3>
          <p>
            <code>POST /api/exercise/add</code>
          </p>
          <input ref="uid" type="text" name="userId" placeholder="userId*" />
          <input
            ref="desc"
            type="text"
            name="description"
            placeholder="description*"
          />
          <input
            ref="dur"
            type="text"
            name="duration"
            placeholder="duration* (mins.)"
          />
          <input
            ref="dat"
            type="text"
            name="date"
            placeholder="date (yyyy-mm-dd)"
          />
          <input type="submit" value="Submit" />
        </form>
        <p>
          <strong>GET users's exercise log: </strong>
          <code>
            GET /api/exercise/log?(id)[&amp;from][&amp;to][&amp;limit]
          </code>
        </p>
        <p>
          <strong>()</strong> = required, <strong>[ ]</strong> = optional
        </p>
        <p>
          <strong>from, to</strong> = dates (yyyy-mm-dd); <strong>limit</strong>{" "}
          = number
        </p>
      </div>
    );
  }
}

export default Exercise;