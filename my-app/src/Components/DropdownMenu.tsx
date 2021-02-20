import React, { useState } from 'react';

export function DropdownMenu() {
    // dropdown menu state
    const [showMenu, toggleMenu] = useState(false);

    return (
        <div>
            <button onClick={() => toggleMenu(showMenu => !showMenu)}>
                {showMenu ? <img src="./assets/icons/up-arrow.png"></img> : <img src="./assets/icons/down-arrow.png"></img>}
            </button>
        </div>
    );
}