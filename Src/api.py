import flask
from flask import render_template
import os

app = flask.Flask(__name__)

@app.route('/')
def index():
    return flask.render_template('welcome.html')

if __name__ == "__main__":
    app.run(debug=True)