import './App.css';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import DashboardPage from './Pages/DashboardPage';

function App() {
  return (
    <Router>
      <nav>
        <ul>
          <li><Link to="/">Dashboard</Link></li>
        </ul>
      </nav>
      <Route path="/" component={DashboardPage}/>
      <div className="App">
      </div>
    </Router>
  );
}

export default App;
