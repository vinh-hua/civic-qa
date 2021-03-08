from flask import Flask, render_template, url_for, jsonify, request
import textAnalytics
app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False

# articles = [
#     """
#     Washington, D.C. Autumn in DC is a uniquely beautiful season. The leaves fall from the trees
#     in a city chockful of forrests, leaving yellow leaves on the ground and a clearer view of the
#     blue sky above...
#     """,
#     """
#     Redmond, WA. In the past few days, Microsoft has decided to further postpone the start date of
#     its United States workers, due to the pandemic that rages with no end in sight...
#     """,
#     """
#     Redmond, WA. Employees at Microsoft can be excited about the new coffee shop that will open on campus
#     once workers no longer have to work remotely...
#     """
# ]
# documents = [
#     "The employee's SSN is 859-98-0987.",
#     "The employee's phone number is 555-555-5555."
# ]

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/key-phrase', methods=['POST'])
def extract_key_phrases_api():
    data = request.get_json()
    text_input = data['body']
    response = textAnalytics.extract_key_phrases(text_input)
    return jsonify(response)

@app.route('/pii-entities', methods=['POST'])
def extract_pii_entities_api():
    data = request.get_json()
    text_input = data['body']
    response = textAnalytics.extract_pii_entities(text_input)
    return jsonify(response)