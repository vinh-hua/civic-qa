import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from "react-router-dom";
import { Provider } from 'react-redux'
import store from './Redux/Store/store'
import App from './App';
import './index.css';

const rootElement = document.getElementById("root");
ReactDOM.render(
    <Router>
        <Provider store={store}>
            <App />   
        </Provider>
    </Router>, 
    rootElement
);
