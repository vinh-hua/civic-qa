

class Controller:

    def __init__(self, analytics_client) -> None:
        self.client = analytics_client

    def extract_keyphrases(self, documents: 'list[str]', num_phrases = 3) -> 'list[str]':
        return self.client.extract_key_phrases(documents)[:num_phrases]
