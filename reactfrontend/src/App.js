import React, { Component } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import Major from './components/major';
import Student from './components/student';
import Home from './components/home';

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <h2>React + Gin + MySQL</h2>
            <table>
              <tbody>
              <tr>
              <td><Link to={'/'} className="nav-link"> Home </Link></td>
              <td><Link to={'/major'} className="nav-link">Major</Link></td>
              <td><Link to={'/student'} className="nav-link">Student</Link></td>
              </tr>
              </tbody>
            </table>
          <hr />
          <Switch>
            <Route exact path='/' component={Home} />
            <Route path='/major' component={Major} />
            <Route path='/student' component={Student} />
          </Switch>
        </div>
      </Router>
    )
  }
}

export default App;