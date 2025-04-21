from flask import Flask, render_template
from flask_lambda import FlaskLambda  # Import FlaskLambda for AWS Lambda compatibility

app = FlaskLambda(__name__, template_folder='HTML')  # Use FlaskLambda instead of Flask


@app.route('/')
def main():
    return render_template("welcome.html", title="Welcome")


@app.route('/signin')
def signin():
    return render_template("signin.html", title="Sign In")


@app.route('/about')
def about():
    return render_template("about.html", title="About")


if __name__ == "__main__":
    app.run()