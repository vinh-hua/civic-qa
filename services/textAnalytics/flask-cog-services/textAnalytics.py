import os
import sys

def authenticate_client():
    from azure.core.credentials import AzureKeyCredential
    from azure.ai.textanalytics import TextAnalyticsClient

    endpoint = os.environ["AZURE_TEXT_ANALYTICS_ENDPOINT"]
    key = os.environ["COGNITIVE_SERVICE_KEY"]
    text_analytics_client = TextAnalyticsClient(endpoint=endpoint, credential=AzureKeyCredential(key))
    return text_analytics_client

def extract_key_phrases(input):
    print(
        "In this sample, we want to find the articles that mention Microsoft to read."
    )
    client = authenticate_client()
    result = client.extract_key_phrases(input)
    print(result, file=sys.stderr)
    resultMap = {}
    for idx, doc in enumerate(result):
        if not doc.is_error:
            resultMap[idx] = doc.key_phrases
    return resultMap

def extract_pii_entities(input):
    client = authenticate_client()
    response = client.recognize_pii_entities(input, language="en")
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
