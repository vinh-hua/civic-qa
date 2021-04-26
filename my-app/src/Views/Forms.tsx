import { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { FormCard } from '../Components/FormCard';
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
        var forms: FormData[] = [];
        formResponses.forEach(function(form: any) {
            var id = form.id;
            var name = form.name;
            forms.push({id, name});
        });
        setFormData(forms);
    }

    function getFormLink(id: string) {
        var iframeString: string =  "<iframe src=\"" + Endpoints.Base + "/form/" + id + "\"></iframe>";
        navigator.clipboard.writeText(iframeString);
    }
        
    useEffect(() => {
        getForms();
    }, []);

    var formCards: any[] = [];
    formData.forEach(form => formCards.push(<FormCard id={form.id} name={form.name} getForm={getFormLink}></FormCard>));

    return(
        <div className="dashboard subdashboard">
            <Header title="Your Forms"></Header>
            <div>
                {formCards}
            </div>
        </div>
    );
}