import requests
import unittest
import string
import random

"""
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
"""

GATEWAY_URL = "http://localhost/v0"

class TestAccount(unittest.TestCase):

    def randstr(self, n):
        return "".join(random.choice(string.ascii_lowercase) for _ in range(n))

    def generate_user(self):
        user = dict()
        user["email"] = f"{self.randstr(12)}@example.com"
        password = self.randstr(12)
        user["password"], user["passwordConfirm"] = password, password
        user["firstName"] = "test_fname"
        user["lastName"] = "test_lname"
        return user

    def make_user(self, user_dict):
        resp = requests.post(GATEWAY_URL+"/signup", json=user_dict, headers={"content-type": "application/json"})
        if resp.status_code != 201:
            raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")
        
        return resp.headers["Authorization"]

    def login(self, credentials_dict):
        resp = requests.post(GATEWAY_URL+"/login", json=credentials_dict, headers={"content-type": "application/json"})
        if resp.status_code != 200:
            raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")
        
        return resp.headers["Authorization"]
        
    def logout(self, auth_header):
        resp = requests.post(GATEWAY_URL+"/logout", headers={"content-type": "application/json", "Authorization": auth_header})
        if resp.status_code != 200:
            raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

    def getsession(self, auth_header):
        resp = requests.get(GATEWAY_URL+"/getsession", headers={"content-type": "application/json", "Authorization": auth_header})
        if resp.status_code != 200:
            raise ValueError(f"Status code: {resp.status_code} | error: {resp.text}")

        return resp.json()


    def test_signup(self):
        print("Testing signup")
        self.make_user(self.generate_user())

    def test_login(self):
        print("Testing login")
        user = self.generate_user()
        self.make_user(user)

        creds = {"email": user["email"],"password": user["password"]}
        self.login(creds)

    def test_logout(self):
        print("Testing logout")
        auth = self.make_user(self.generate_user())
        self.logout(auth)

    def test_getsession(self):
        print("Testing get session")
        auth = self.make_user(self.generate_user())
        self.getsession(auth)

if __name__ == '__main__':
    unittest.main()