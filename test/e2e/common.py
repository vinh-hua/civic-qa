import random
import requests
import string
from urllib.parse import urlencode


def randstr(low, high=None):
        # default high == low
        if high is None:
            high = low
        
        # if high != low, pick random [low, high]
        n = low
        if high != low:
            n = random.randint(low, high)
            
        return "".join(random.choice(string.ascii_lowercase) for _ in range(n))



def make_user(URL, user_dict):
        resp = requests.post(URL+"/signup", json=user_dict, headers={"content-type": "application/json"})
        if resp.status_code != 201:
            raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")
        
        return resp.headers["Authorization"]

def login(URL, credentials_dict):
    resp = requests.post(URL+"/login", json=credentials_dict, headers={"content-type": "application/json"})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")
    
    return resp.headers["Authorization"]
    
def logout(URL, auth_header):
    resp = requests.post(URL+"/logout", headers={"content-type": "application/json", "Authorization": auth_header})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

def getsession(URL, auth_header):
    resp = requests.get(URL+"/getsession", headers={"content-type": "application/json", "Authorization": auth_header})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.json()

def generate_user():
        user = dict()
        user["email"] = f"{randstr(8,12)}@example.com"
        password = randstr(12)
        user["password"], user["passwordConfirm"] = password, password
        user["firstName"] = randstr(4,8)
        user["lastName"] = randstr(4,10)
        return user

def generate_form():
    form = dict()
    form["name"] = f"Form: {randstr(8,12)}"
    return form

def make_form(URL, auth_header, form_dict):
    resp = requests.post(URL+"/forms",
        headers={"content-type": "application/json",
        "Authorization": auth_header},
        json=form_dict)

    if resp.status_code != 201:
            raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.json()

def get_forms(URL, auth_header):
    resp = requests.get(URL+"/forms", headers={"content-type": "application/json", "Authorization": auth_header})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.json()

def get_form(URL, auth_header, form_id):
    resp = requests.get(URL+"/forms/"+str(form_id), headers={"content-type": "application/json", "Authorization": auth_header})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.json()

def get_form_client(URL, form_id):
    resp = requests.get(URL+"/form/"+str(form_id), headers={"content-type": "application/json"})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.text

def generate_response():
    return { 
        "name": f"{randstr(4,8)} {randstr(5,12)}", 
        "email": f"{randstr(6,12)}@example.com",
        "inquiryType": random.choice(["general", "casework"]),
        "subject": randstr(8, 20),
        "body": randstr(20, 45)
    }

def post_form_user(URL, form_id, form_dict):
    resp = requests.post(URL+"/form/"+str(form_id), data=form_dict)
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.text


def get_response(URL, auth_header, resp_id):
    resp = requests.get(URL+"/responses/"+str(resp_id), headers={"Authorization": auth_header})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.json()

def get_responses(URL, auth_header):
    resp = requests.get(URL+"/responses", headers={"Authorization": auth_header})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.json()

def get_responses_params(URL, auth_header, queryParams: 'dict'):
    resp = requests.get(URL+f"/responses?{urlencode(queryParams)}", headers={"Authorization": auth_header})
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.json()

def patch_response(URL, auth_header, resp_id, new_state):
    body = {"active": new_state}
    resp = requests.patch(URL+"/responses/"+str(resp_id), headers={"Authorization": auth_header}, json=body)
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.status_code

def post_mailto(URL, body):
    resp = requests.post(URL+"/mailto", headers={"content-type": "application/json"}, json=body)
    if resp.status_code != 200:
        raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    return resp.text


