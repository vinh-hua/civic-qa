import ReactDOM from 'react-dom';
import { HashRouter } from "react-router-dom";
import { createHashHistory } from 'history';
import { Provider } from 'react-redux'
import store from './Redux/Store/store'
import App from './App';
import './index.css';

const rootElement = document.getElementById("root");

ReactDOM.render(
    <HashRouter basename="">
        <Provider store={store}>
            <App />   
        </Provider>
    </HashRouter>, 
    rootElement
);
