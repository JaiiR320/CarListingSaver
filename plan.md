Create an API to add car listings from multiple websites
    The listings will hold
    - Name of listing
    - price
    - car meta data (mileage, title, transmission type etc.)
    - original URL of the listing
    - date
    - ...

To add a listing to the API, make a POST with URL and other meta data in request

API innerworkings
POST
    - verify website source (facebook, craigslist etc)
    - scrape website for information and fill in a struct
    - store that struct on the db
    - store the user's ID on the struct to access in the db
        - SELECT * FROM users WHERE id = ?
    - maybe JWT
    - logins
GET
    - verify user
    - get listings user created
    - display as HTML on a dashboard

TODO
1. webscrape facebook for proof of concept
2. build api to interface with webscraping
3. build dashboard to show listings and interact with API

x. make a firefox/chrome extension to quickly add 