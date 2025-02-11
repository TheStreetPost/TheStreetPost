from flask import Flask, render_template

app = Flask(__name__, template_folder='HTML')


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
