import requests
import unittest
import common

"""
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
"""

GATEWAY_URL = "http://localhost/v0"

class TestAccount(unittest.TestCase):

    def test_signup(self):
        print("Testing signup")
        common.make_user(GATEWAY_URL, common.generate_user())

    def test_login(self):
        print("Testing login")
        user = common.generate_user()
        common.make_user(GATEWAY_URL, user)

        creds = {"email": user["email"],"password": user["password"]}
        common.login(GATEWAY_URL, creds)

    def test_logout(self):
        print("Testing logout")
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        common.logout(GATEWAY_URL, auth)

    def test_getsession(self):
        print("Testing get session")
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        common.getsession(GATEWAY_URL, auth)

    
class TestForm(unittest.TestCase):

    def test_make_form(self):
        print("Testing make form")

        auth = common.make_user(GATEWAY_URL, common.generate_user())
        common.make_form(GATEWAY_URL, auth, common.generate_form())

    def test_get_forms(self):
        print("Testing get forms")

        auth = common.make_user(GATEWAY_URL, common.generate_user())
        common.make_form(GATEWAY_URL, auth, common.generate_form())
        common.make_form(GATEWAY_URL, auth, common.generate_form())
        common.get_forms(GATEWAY_URL, auth)

if __name__ == '__main__':
    unittest.main()