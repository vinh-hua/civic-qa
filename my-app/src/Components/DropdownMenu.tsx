import { useState, Dispatch, SetStateAction } from 'react';
import * as Constants from '../Constants/constants';
import './DropdownMenu.css';

export type DropdownMenuProps = {
    chartView: string;
    setChartView: Dispatch<SetStateAction<string>>;
};

export function DropdownMenu(props: DropdownMenuProps) {
    // dropdown menu state
    const [showMenu, toggleMenu] = useState(false);

    // set chart view and auto close menu
    function SetChartViewAndToggleMenu(view: string) {
        props.setChartView(view);
        toggleMenu(showMenu => !showMenu);
    }

    return (
        <div>
            <button className="dropdown-menu-btn" onClick={() => toggleMenu(showMenu => !showMenu)}>
                <p className="dropdown-menu-btn-text">{props.chartView}</p>
                {showMenu ? <img className="dropdown-menu-arrow" src="./assets/icons/up-arrow.png" /> : <img className="dropdown-menu-arrow" src="./assets/icons/down-arrow.png" />}
            </button>
            {showMenu ? 
                <div className="menu">
                    <button onClick={() => SetChartViewAndToggleMenu(Constants.Responses)} className="dropdown-menu-btn"><p className="dropdown-menu-btn-text">{Constants.Responses}</p></button>
                    <button onClick={() => SetChartViewAndToggleMenu(Constants.Topics)} className="dropdown-menu-btn"><p className="dropdown-menu-btn-text">{Constants.Topics}</p></button>
                </div> : <div />}
        </div>
    );
}