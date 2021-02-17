import React from 'react';
import './StatCard.css';

const StatCard = (props) => {
    return (
        <div class="stat-card">
            <h1 class="stat-card-title">{props.title}</h1>
            <p class="stat-card-data">{props.data}</p>
        </div>
    );
}

export default StatCard;