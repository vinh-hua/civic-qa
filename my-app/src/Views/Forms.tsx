import { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { FormCard } from '../Components/FormCard';
import './Forms.css';
import * as Endpoints from '../Constants/Endpoints';

export type FormData = {
    id: string;
    name: string;
}

export function Forms() {
    const [formData, setFormData] = useState<FormData[]>([]);
    const [newFormName, setNewFormName] = useState("");

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

    const createForm = async(e: any) => {
        e.preventDefault();
        const authToken = localStorage.getItem("Authorization") || "";
        var newForm = {name: newFormName};
        var jsonNewForm = JSON.stringify(newForm);
        const response = await fetch(Endpoints.Base + "/forms", {
            method: "POST",
            body: jsonNewForm,
            headers: new Headers({
                "Content-Type": "application/json",
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            alert("There was an error trying to get your forms.");
            return;
        }
        setNewFormName("");
        getForms();
    }
        
    useEffect(() => {
        getForms();
    }, []);

    var formCards: any[] = [];
    formData.forEach(form => formCards.push(<FormCard id={form.id} name={form.name} getForm={getFormLink}></FormCard>));

    return(
        <div className="dashboard subdashboard">
            <Header title="Your Forms"></Header>
            <div className="form-cards">
                {formCards}
            </div>
            <hr className="forms-create-divider"/>
            <div className="new-form-container">
                <h1>Create a form</h1>
                <form className="new-form" autoComplete="off" onSubmit={createForm}>
                    <input className="new-form-input" name="name" type="text" value={newFormName} placeholder="Form Name" onChange={e => setNewFormName(e.target.value)} required/>
                    <input className="new-form-create-btn" type="submit" value="Create" />
                </form>
            </div>
        </div>
    );
}