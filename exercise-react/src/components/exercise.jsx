import React, { Component } from "react";
import styles from "./styling.css";
class Exercise extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      id: "",
      description: "",
      duration: "",
      date: ""
    };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleSubmit1 = this.handleSubmit1.bind(this);
  }
  styles = {
    fontSize: 20,
    fontWeight: "bold"
  };
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
        var winPrint = window.open(
          "",
          "",
          "left=0,top=0,width=800,height=600,toolbar=0,scrollbars=0,status=0"
        );
        winPrint.document.write(JSON.stringify(json));
        winPrint.document.close();
        winPrint.focus();
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
        var winPrint = window.open(
          "",
          "",
          "left=0,top=0,width=800,height=600,toolbar=0,scrollbars=0,status=0"
        );
        winPrint.document.write(JSON.stringify(json));
        winPrint.document.close();
        winPrint.focus();
      });
    event.preventDefault();
  }

  handleEmailChange = e => {
    this.setState({ username: e.target.value });
  };
  state = {};
  render() {
    return (
      <div className="container  badge-secondary">
        <h1>Exercise tracker</h1>

        <form
          name="usercreate"
          className="badge badge-primary m-2"
          style={this.styles}
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
            className="m-3"
            onChange={this.handleChange}
            required
          />
          <input
            type="submit"
            value="Submit"
            className="btn btn-secondary btn-sm"
          />
        </form>

        <div>
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        </div>
        <div>
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        </div>

        <form
          name="usercreate1"
          className="badge badge-warning m-2"
          onSubmit={this.handleSubmit1}
        >
          <h3>Add exercises</h3>
          <p>
            <code style={{ fontSize: 20 }}>POST /api/exercise/add</code>
          </p>
          <td>
            <input
              ref="uid"
              type="text"
              name="userId"
              placeholder="userId*"
              className="m-3"
              required
            />
          </td>
          <td>
            <input
              ref="desc"
              type="text"
              name="description"
              placeholder="description*"
              className="m-3"
              required
            />
          </td>
          <td>
            <input
              ref="dur"
              type="text"
              name="duration"
              pattern="[0-9]*"
              placeholder="duration* (mins.)"
              className="m-3"
              required
            />
          </td>
          <td>
            <input
              ref="dat"
              type="text"
              pattern="(?:19|20)[0-9]{2}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1[0-9]|2[0-9])|(?:(?!02)(?:0[1-9]|1[0-2])-(?:30))|(?:(?:0[13578]|1[02])-31))"
              name="date"
              className="m-3"
              placeholder="date (yyyy-mm-dd)"
            />
          </td>
          <td>
            <input
              type="submit"
              value="Submit"
              className="btn btn-secondary btn-sm"
            />
          </td>
        </form>
        <div>
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        </div>
        <div>
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        </div>
        <p className="badge badge-success m-2" style={{ fontSize: 20 }}>
          <strong>GET users's exercise log: </strong>
          <code>
            GET /api/exercise/log?(id)[&amp;from][&amp;to][&amp;limit]
          </code>
        </p>
        <div>
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        </div>
        <p className="badge badge-info m-2" style={{ fontSize: 20 }}>
          <strong>()</strong> = required, <strong>[ ]</strong> = optional
        </p>
        <p className="badge badge-info m-2" style={{ fontSize: 20 }}>
          <strong>from, to</strong> = dates (yyyy-mm-dd); <strong>limit</strong>{" "}
          = number
        </p>
      </div>
    );
  }
}

export default Exercise;
