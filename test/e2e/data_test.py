from gateway_test import GATEWAY_URL
import unittest
import common

GATEWAY_URL = "http://localhost/v0"

class TestData(unittest.TestCase):


    def test_data(self):
        users = 10
        forms_per_user = 10
        responses_per_form = 100

        for u in range(users):
            user = common.generate_user()
            auth = common.make_user(GATEWAY_URL, user)

            for f in range(forms_per_user):
                
                form = common.make_form(GATEWAY_URL, auth, common.generate_form())
                
                print(f"User: {u}, form: {f}")

                for r in range(responses_per_form):
                    resp = common.generate_response()
                    common.post_form_user(GATEWAY_URL, form["id"], resp)



if __name__ == "__main__":
    unittest.main()