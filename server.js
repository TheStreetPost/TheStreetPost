const express = require('express');
const path = require('path');

const app = express();
const PORT = process.env.PORT || 3000;

// Set up view engine and static files
app.set('view engine', 'html');
app.set('views', path.join(__dirname, 'HTML'));
app.engine('html', (filepath, options, callback) => {
  const fs = require('fs');
  fs.readFile(filepath, (err, content) => {
    if (err) return callback(err);
    const rendered = content.toString();
    return callback(null, rendered);
  });
});

// Serve static files
app.use(express.static(path.join(__dirname, 'static')));

// Routes
app.get('/', (req, res) => {
  res.render('welcome.html', { title: 'Welcome' });
});

app.get('/signin', (req, res) => {
  res.render('signIn.html', { title: 'Sign In' });
});

app.get('/about', (req, res) => {
  res.render('about.html', { title: 'About' });
});

// Start server
app.listen(PORT, () => {
  console.log(`TheStreetPost server is running on http://localhost:${PORT}`);
});
