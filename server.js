const express = require('express');
const path = require('path');

const app = express();
const PORT = process.env.PORT || 3000;

// Set up view engine and static files
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));

// Serve static files
app.use(express.static(path.join(__dirname, 'static')));

// Helper function to render pages with template
const renderPage = (page) => {
  return (req, res) => {
    res.render('template', {
      title: getPageTitle(page),
      page: `pages/${page}`
    });
  };
};

// Get page titles
const getPageTitle = (page) => {
  const titles = {
    'welcome': 'Welcome',
    'news': 'News',
    'blog': 'Blog',
    'about': 'About',
    'contact': 'Contact',
    'signin': 'Sign In',
    'signup': 'Sign Up'
  };
  return titles[page] || 'TheStreetPost';
};

// Routes
app.get('/', renderPage('welcome'));
app.get('/news', renderPage('news'));
app.get('/blog', renderPage('blog'));
app.get('/about', renderPage('about'));
app.get('/contact', renderPage('contact'));
app.get('/signin', renderPage('signin'));
app.get('/signup', renderPage('signup'));

// Handle 404
app.use((req, res) => {
  res.status(404).render('pages/404', { title: 'Not Found' });
});

// Start server
app.listen(PORT, () => {
  console.log(`TheStreetPost server is running on http://localhost:${PORT}`);
});
