import common
import unittest
import time

GATEWAY_URL = "http://localhost/v0"


def timeit(name, itr):
    """Return a function decorator that prints the time
       taken for execution, with a given name

    Args:
        name (str): test/method name
    """
    def decor(f):
        def inner(*args, **kwargs):
            t_start = time.perf_counter()
            f(*args, **kwargs)
            elapsed = time.perf_counter() - t_start
            print(f"{name} executed in {elapsed} seconds | N={itr} | {elapsed / itr} per iteration")
        return inner
    return decor

class TestLoad(unittest.TestCase):
    
    def test_signup(self):
        N = 100
        users = [common.generate_user() for _ in range(N)]

        @timeit(f"test_signup", N)
        def run():
            for u in users:
                common.make_user(GATEWAY_URL, u)
        run()

    def test_login(self):
        N = 100
        user = common.generate_user()
        common.make_user(GATEWAY_URL, user)

        creds = {"email": user["email"], "password": user["password"]}

        @timeit(f"test_login", N)
        def run():
            for _ in range(N):
                common.login(GATEWAY_URL, creds)
        run()

    def test_get_form_client(self):
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())

        N = 100

        @timeit(f"test_get_form_client", N)
        def run():
            for _ in range(N):
                common.get_form_client(GATEWAY_URL, form["id"])
        run()

    def test_get_responses(self):
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())
        form2 = common.make_form(GATEWAY_URL, auth, common.generate_form())
        form3 = common.make_form(GATEWAY_URL, auth, common.generate_form())

        common.post_form_user(GATEWAY_URL, form["id"], common.generate_form())
        common.post_form_user(GATEWAY_URL, form["id"], common.generate_form())
        common.post_form_user(GATEWAY_URL, form2["id"], common.generate_form())
        common.post_form_user(GATEWAY_URL, form2["id"], common.generate_form())
        common.post_form_user(GATEWAY_URL, form3["id"], common.generate_form())
        common.post_form_user(GATEWAY_URL, form3["id"], common.generate_form())


        N = 100

        @timeit(f"test_get_responses", N)
        def run():
            for _ in range(N):
                common.get_responses(GATEWAY_URL, auth)
        run()

    def test_patch_response(self):

        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())
        common.post_form_user(GATEWAY_URL, form["id"], common.generate_form())

        resp = common.get_responses_params(GATEWAY_URL, auth, {"formID": form["id"]})[0]

        N = 100

        @timeit(f"test_patch_response", N)
        def run():
            new_state = False
            for _ in range(N):
                common.patch_response(GATEWAY_URL, auth, resp["id"], new_state)
                new_state = not new_state
        run()

    def test_create_tag(self):
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())
        common.post_form_user(GATEWAY_URL, form["id"], common.generate_response())
        resp = common.get_responses_params(GATEWAY_URL, auth, {"formID": form["id"]})[0]
        N = 100

        tag_vals = [f"test_{i}" for i in range(N)]

        @timeit("test_create_tag", N)
        def run():
            for i in range(N):
                common.make_tag(GATEWAY_URL, auth, resp["id"], tag_vals[i])

        run()


    def test_mailto(self):
        body = {
            "to": ["test@example.com"],
            "subject": "test subject",
            "body": "test body"
        }

        N = 100

        @timeit(f"test_mailto", N)
        def run():
            for _ in range(N):
                common.post_mailto(GATEWAY_URL, body)

        run()





        

if __name__ == '__main__':
    unittest.main()