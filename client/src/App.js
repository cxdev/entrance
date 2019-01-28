import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import 'semantic-ui-css/semantic.min.css'

import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import CommandAddPage from './containers/CommandAddPage'
import CommandsPage from './containers/CommandsPage'
import CommandDetailPage from './containers/CommandDetailPage'
import JobsPage from './containers/JobsPage'
import JobDetailPage from './containers/JobDetailPage'

const Index = () => <h2>Home</h2>;

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <nav>
            <ul>
              <li>
                <Link to="/">Home</Link>
              </li>
              <li>
                <Link to="/admin/add_command">Add Command</Link>
              </li>
              <li>
                <Link to="/commands/">Commands</Link>
              </li>
              <li>
                <Link to="/jobs/">Jobs</Link>
              </li>
            </ul>
          </nav>

          <Route path="/" exact component={Index} />
          <Route path="/admin/add_command" component={CommandAddPage} />
          <Route path="/commands/" component={CommandsPage} />
          <Route path="/command/:commandId" component={CommandDetailPage} />
          <Route path="/jobs/" component={JobsPage} />
          <Route path="/job/:jobId" component={JobDetailPage} />
        </div>
      </Router>
    );
  }
}

export default App;
