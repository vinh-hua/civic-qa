import React from 'react';
import './ProfileHeader.css';

export function ProfileHeader() {
    return (
        <div className="profile-heading">
            <p className="profile-name">Profile Name</p>
            <svg height="100" width="100">
                <circle cx="25" cy="25" r="24" stroke="#DFE0EB" stroke-width="1" fill="white" />
            </svg>
        </div>
    );
}