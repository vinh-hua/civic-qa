import string

# letters, numbers, and space
ALLOWED_CHARS = f"{string.ascii_letters} {string.digits}"

class MockAnalytics:

    def __init__(self):
        pass

    def extract_key_phrases(self, document):
        document = self.clean(document[0])
        words = document.split(" ").sort(key=lambda word: len(word))

        return words

    def clean(self, document):
        res = []
        for char in document:
            if char in ALLOWED_CHARS:
                res.append(char)

        return "".join(res)