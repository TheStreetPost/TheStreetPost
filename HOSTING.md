# TheStreetPost - Node.js Version

## Installation & Setup

1. **Install Node.js** (if not already installed)
   - Download from https://nodejs.org/

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Run locally**
   ```bash
   npm start
   ```
   - Visit http://localhost:3000 in your browser

## Hosting on Hostinger

### Using cPanel (Recommended for Hostinger)

1. **Upload your files via FTP or File Manager**
   - Connect to your Hostinger account
   - Upload the entire project folder to your public_html directory

2. **Install Node.js on Hostinger**
   - Go to cPanel → Node.js Manager
   - Create a new Node.js application
   - Set the app root to your project folder
   - Click "Create"

3. **Install dependencies on the server**
   - In the Node.js Manager, click "NPM" for your app
   - Run: `npm install`

4. **Set environment**
   - Click on your application
   - Make sure it's running
   - Note the URL provided (usually port 3000 or custom port)

### Alternative: Using Hostinger's Application Manager

1. Upload files to your hosting account
2. Go to Hosting → Manage → Application Manager
3. Create a new Node.js application
4. Select your project folder
5. Click "Create" and wait for it to be deployed

## Project Structure

```
TheStreetPost/
├── server.js           # Main server file (replaces app.py)
├── package.json        # Node.js dependencies
├── HTML/               # HTML templates
│   ├── welcome.html
│   ├── signIn.html
│   ├── about.html
│   └── template.html
└── static/             # Static files (CSS, images)
    ├── style.css
    └── images/
```

## Notes

- The app runs on port 3000 by default (can be changed with PORT environment variable)
- All routes from the Flask app have been migrated
- Static files are served from the `static/` folder
- Hostinger will provide a public URL once deployed
