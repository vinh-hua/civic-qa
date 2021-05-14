from rake_nltk import Rake
import string

class RakeAnalytics:

    def __init__(self):
        pass

    def extract_key_phrases(self, document):
        rake_nltk_var = Rake(max_length=2)
        rake_nltk_var.extract_keywords_from_text(document)
        keyword_extracted = rake_nltk_var.get_ranked_phrases()[:3]
        return keyword_extracted