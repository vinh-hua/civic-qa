import common
import unittest
import time

GATEWAY_URL = "http://localhost/v0"


def timeit(name):
    def decor(f):

        def inner(*args, **kwargs):
            t_start = time.perf_counter()
            f(*args, **kwargs)
            print(f"{name} executed in {time.perf_counter()} seconds")

        return inner
    return decor

class TestLoad(unittest.TestCase):

    
    def test_signup(self):
        N = 100
        users = [common.generate_user() for _ in range(N)]

        @timeit(f"test_signup {N=}")
        def run():
            for u in users:
                common.make_user(GATEWAY_URL, u)
        run()

    def test_login(self):
        N = 100
        user = common.generate_user()
        common.make_user(GATEWAY_URL, user)

        creds = {"email": user["email"], "password": user["password"]}

        @timeit(f"test_login {N=}")
        def run():
            for _ in range(N):
                common.login(GATEWAY_URL, creds)

        run()

    def test_get_form_user(self):
        auth = common.make_user(GATEWAY_URL, common.generate_user())
        form = common.make_form(GATEWAY_URL, auth, common.generate_form())

        N = 100

        @timeit(f"test_get_form_user {N=}")
        def run():
            for _ in range(N):
                common.get_form_user(GATEWAY_URL, form["id"])

        run()

        

if __name__ == '__main__':
    unittest.main()