from flask import Flask, render_template, url_for, jsonify, request, Response
import textAnalytics
import os
import sys
from azure.core.credentials import AzureKeyCredential
from azure.ai.textanalytics import TextAnalyticsClient
# need to add logging aggregator request
endpoint = os.environ["AZURE_TEXT_ANALYTICS_ENDPOINT"]
key = os.environ["COGNITIVE_SERVICE_KEY"]
addr = os.environ["ADDR"]
text_analytics_client = {}
NUM_TAGS = 3
#VersionBase is the API route base
VersionBase =  "/v0"
try:
    text_analytics_client = TextAnalyticsClient(endpoint=endpoint, credential=AzureKeyCredential(key))
    print("authenticated azure", file=sys.stderr)
except Exception as e:
    print(e, file=sys.stderr)
    print("Couldn't authenticate to Azure Text Analytics", file=sys.stderr)

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False

@app.route(VersionBase + '/key-phrase', methods=['POST'])
def extract_key_phrases_api():
    if not request.json:
        data = {'error': 'Bad request type'}
        return jsonify(data), 400
    data = request.get_json()
    text_input = data['subject'] + ". " + data['body']
    text_input = [text_input]
    print(text_input, file=sys.stderr)
    try:
        resp = textAnalytics.extract_key_phrases(text_input, NUM_TAGS, text_analytics_client)
        return jsonify(resp)
    except Exception:
        data = {'error': 'InternalServerError from Azure Analytics API - keyphrase'}
        return jsonify(data), 500

@app.route(VersionBase + '/v0/pii-entities', methods=['POST'])
def extract_pii_entities_api():
    data = request.get_json()
    text_input = data['body']
    try: 
        resp = textAnalytics.extract_pii_entities(text_input, text_analytics_client)
        return jsonify(resp)
    except Exception:
        data = {'error': 'InternalServerError from Azure Analytics API'}
        return jsonify(data), 500

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=addr)