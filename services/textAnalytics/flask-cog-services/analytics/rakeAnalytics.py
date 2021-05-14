from rake_nltk import Rake
import string

# letters, numbers, and space
ALLOWED_CHARS = f"{string.ascii_letters} {string.digits}"

class RakeAnalytics:

    def __init__(self):
        pass

    def extract_key_phrases(self, document):
        document = self.clean(document[0])
        rake_nltk_var = Rake(max_length=2)
        rake_nltk_var.extract_keywords_from_text(document)
        keyword_extracted = rake_nltk_var.get_ranked_phrases()[:3]
        return keyword_extracted

    def clean(self, document):
        res = []
        for char in document:
            if char in ALLOWED_CHARS:
                res.append(char)

        return "".join(res)