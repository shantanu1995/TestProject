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
        console.log(json);
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
        <form>
          <h3>Add exercises</h3>
          <p>
            <code>POST /api/exercise/add</code>
          </p>
          <input id="uid" type="text" name="userId" placeholder="userId*" />
          <input
            id="desc"
            type="text"
            name="description"
            placeholder="description*"
          />
          <input
            id="dur"
            type="text"
            name="duration"
            placeholder="duration* (mins.)"
          />
          <input
            id="dur"
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
