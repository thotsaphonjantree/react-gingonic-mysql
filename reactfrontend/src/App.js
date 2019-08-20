import React, { Component } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import Major from './components/Major';
import Student from './components/Student';
import Home from './components/Home';
import PostStudent from './components/PostStudent';
import StudentInfo from './components/StudentInfo';
import Studentbymajor from './components/Studentbymajor';
class App extends Component {
  render() {
    return (
      <Router>
        <div>
        <center>
          <div style={{backgroundColor: "#ffffff"}}>
          <h2>React  Gin  MySQL</h2>
          <table>
            <tbody>
              <tr>
                <td><Link to={'/'} className="nav-link"> Home </Link></td>
                <td><Link to={'/major'} className="nav-link">Major</Link></td>
                <td><Link to={'/students'} className="nav-link">Student</Link></td>
                <td><Link to={'/poststudent'} className="nav-link">Add Student</Link></td>
              </tr>
            </tbody>
          </table>
          <hr />
          </div>
          <Switch>
            <Route exact path='/' component={Home} />
            <Route path='/major' component={Major} />
            <Route path='/students' component={Student} />
            <Route path='/poststudent' component={PostStudent} />
            <Route path="/student/:studentId" component={StudentInfo} />
            <Route path="/sbymajor/:majorId" component={Studentbymajor} />
          </Switch>
          </center>
        </div>
      </Router>
    )
  }
}

export default App;