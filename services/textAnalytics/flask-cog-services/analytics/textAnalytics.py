import os
import sys

class AzureAnalytics:

    def __init__(self, analytics_client) -> None:
        self.analytics_client = analytics_client

    def extract_key_phrases(self, input):
        result = self.analytics_client.extract_key_phrases(input)
        result_phrases = []
        for doc in result:
            if not doc.is_error:
                result_phrases = doc.key_phrases
        return result_phrases

    def extract_pii_entities(self, input):
        response = self.analytics_client.recognize_pii_entities(input, language="en")
        result = [doc for doc in response if not doc.is_error]
        personal_info = {}
        for idx, doc in enumerate(result):
            list_of_entities = []
            for entity in doc.entities:
                entity_map = {}
                entity_map["text"] = entity.text
                entity_map["category"] = entity.category
                entity_map["confidence_score"] = entity.confidence_score
                list_of_entities.append(entity_map)
            personal_info[idx] = list_of_entities
        return personal_info
