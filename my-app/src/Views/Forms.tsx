import { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import * as Endpoints from '../Constants/Endpoints';

export type FormData = {
    id: string;
    name: string;
}

export function Forms() {
    const [formData, setFormData] = useState<FormData[]>([]);

    const getForms = async() => {
        const authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Base + "/forms", {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            alert("There was an error trying to get your forms.");
            return;
        }
        const formResponses = await response.json();
        formResponses.forEach(function(form: any) {

        });
    }
        
    useEffect(() => {
        getForms();
    }, []);

    return(
        <div className="dashboard subdashboard">
            <Header title="Your Forms"></Header>
            <div>

            </div>
        </div>
    );
}