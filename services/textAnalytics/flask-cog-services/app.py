from flask import Flask, render_template, url_for, jsonify, request, Response
from azure.core.credentials import AzureKeyCredential
from azure.ai.textanalytics import TextAnalyticsClient
from controller.controller import Controller
from analytics.textAnalytics import AzureAnalytics
from analytics.mockAnalytics import MockAnalytics
from analytics.rakeAnalytics import RakeAnalytics
import os
import sys

# VERSION_BASE is the API route base
VERSION_BASE = "/v0"
NUM_TAGS = 3


# need to add logging aggregator request

analytics_impl = os.getenv("ANALYTICS_IMPL", "rake")
analyticsController = None

if analytics_impl == "azure":
    
    endpoint = os.environ["AZURE_TEXT_ANALYTICS_ENDPOINT"]
    key = os.environ["COGNITIVE_SERVICE_KEY"]
    text_analytics_client = None
    try:
        text_analytics_client = AzureAnalytics(TextAnalyticsClient(endpoint=endpoint, credential=AzureKeyCredential(key)))
        analyticsController = Controller(text_analytics_client)

        print("authenticated azure", file=sys.stderr)
    except Exception as e:
        print(e, file=sys.stderr)
        print("Couldn't authenticate to Azure Text Analytics", file=sys.stderr)

if analytics_impl == "mock":
    analyticsController = Controller(MockAnalytics())
    print("Mock analytics")

if analytics_impl == "rake":
    analyticsController = Controller(RakeAnalytics())
    print("Rake analytics")


app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False

@app.route(f'{VERSION_BASE}/key-phrase', methods=['POST'])
def extract_key_phrases_api():

    if not request.json:
        data = {'error': 'Bad request type'}
        return jsonify(data), 415


    data = request.get_json()
    text_input = [f"{data['subject']} . {data['body']}"]


    try:
        resp = analyticsController.extract_keyphrases(text_input, NUM_TAGS)
        return jsonify(resp)
    except Exception as e:
        print(e.with_traceback(), file=sys.stderr)
        data = {'error': 'InternalServerError from Azure Analytics API - keyphrase'}
        return jsonify(data), 500

if __name__ == "__main__":
    addr = os.getenv("ADDR", 9090)
    app.run(host='0.0.0.0', port=addr)