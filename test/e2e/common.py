import random
import requests
import string

def randstr(n):
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
        user["email"] = f"{randstr(12)}@example.com"
        password = randstr(12)
        user["password"], user["passwordConfirm"] = password, password
        user["firstName"] = "test_fname"
        user["lastName"] = "test_lname"
        return user

def generate_form():
    form = dict()
    form["name"] = f"Form: {randstr(12)}"
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