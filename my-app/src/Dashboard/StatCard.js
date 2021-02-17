import React from 'react';
import './StatCard.css';

export function StatCard(props) {
    return (
        <div class="stat-card">
            <h1 class="stat-card-title">{props.title}</h1>
            <p class="stat-card-data">{props.data}</p>
        </div>
    );
}