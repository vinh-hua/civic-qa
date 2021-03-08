import os
import sys

def extract_key_phrases(input):
    print(
        "In this sample, we want to find the articles that mention Microsoft to read."
    )
    # [START extract_key_phrases]
    from azure.core.credentials import AzureKeyCredential
    from azure.ai.textanalytics import TextAnalyticsClient

    endpoint = os.environ["AZURE_TEXT_ANALYTICS_ENDPOINT"]
    key = os.environ["COGNITIVE_SERVICE_KEY"]
    

    text_analytics_client = TextAnalyticsClient(endpoint=endpoint, credential=AzureKeyCredential(key))
    articles = [
        """
        Washington, D.C. Autumn in DC is a uniquely beautiful season. The leaves fall from the trees
        in a city chockful of forrests, leaving yellow leaves on the ground and a clearer view of the
        blue sky above...
        """,
        """
        Redmond, WA. In the past few days, Microsoft has decided to further postpone the start date of
        its United States workers, due to the pandemic that rages with no end in sight...
        """,
        """
        Redmond, WA. Employees at Microsoft can be excited about the new coffee shop that will open on campus
        once workers no longer have to work remotely...
        """
    ]

    result = text_analytics_client.extract_key_phrases(articles)
    print(result, file=sys.stderr)
    resultMap = {}
    for idx, doc in enumerate(result):
        if not doc.is_error:
            resultMap[idx] = doc.key_phrases
            # print("Key phrases in article #{}: {}".format(
            #     idx + 1,
            #     ", ".join(doc.key_phrases)
            # ))
    return resultMap