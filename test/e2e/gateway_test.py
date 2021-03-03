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

    def test_get_form(self):
        print("Testing get form")

        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())

        common.get_form(GATEWAY_URL, auth, form["id"])

    def test_get_form_user(self):
        print("Testing get form: user")

        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())

        common.get_form_user(GATEWAY_URL, form["id"])

    def test_post_form_user(self):
        print("Testing post form: user")
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())

        common.post_form_user(GATEWAY_URL, form["id"], common.generate_response())

    def test_get_responses(self):
        print("Testing get responses")
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())

        common.post_form_user(GATEWAY_URL, form["id"], common.generate_response())
        common.post_form_user(GATEWAY_URL, form["id"], common.generate_response())
        common.post_form_user(GATEWAY_URL, form["id"], common.generate_response())

        common.get_responses(GATEWAY_URL, form["id"], auth)

    def test_get_response(self):
        print("Testing get response")
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())

        common.post_form_user(GATEWAY_URL, form["id"], common.generate_response())
        resps = common.get_responses(GATEWAY_URL, form["id"], auth)

        resp_id = resps[0]["id"]
        common.get_response(GATEWAY_URL, resp_id, auth)


if __name__ == '__main__':
    unittest.main()