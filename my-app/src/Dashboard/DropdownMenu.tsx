import { useState } from 'react';
import * as Constants from '../Constants/constants';
import './DropdownMenu.css';

export function DropdownMenu() {
    // dropdown menu state
    const [showMenu, toggleMenu] = useState(false);

    return (
        <div>
            <button className="dropdown-menu-btn" onClick={() => toggleMenu(showMenu => !showMenu)}>
                <p className="dropdown-menu-btn-text">All Emails</p>
                {showMenu ? <img className="dropdown-menu-arrow" src="./assets/icons/up-arrow.png" /> : <img className="dropdown-menu-arrow" src="./assets/icons/down-arrow.png" />}
            </button>
            {showMenu ? 
                <div className="menu">
                    <button className="dropdown-menu-btn"><p className="dropdown-menu-btn-text">{Constants.AllEmails}</p></button>
                    <button className="dropdown-menu-btn"><p className="dropdown-menu-btn-text">{Constants.UnreadEmails}</p></button>
                    <button className="dropdown-menu-btn"><p className="dropdown-menu-btn-text">{Constants.ResponseTime}</p></button>
                    <button className="dropdown-menu-btn"><p className="dropdown-menu-btn-text">{Constants.Topics}</p></button>
                </div> : <div />}
        </div>
    );
}