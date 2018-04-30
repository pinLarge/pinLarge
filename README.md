# pinLarge

Adapt pinterest rss feed to provide URL to large images and untruncated feed titles.

You can use this feed to autopost the original pinterest images to twitter for example.

Click the button below to deploy this application to Heroku. Heroku will host
it for you for free, as long as you don't use it more than 18 hours a day (this
is a soft limit). If you want to use it in 24/7, it'll cost you.

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## How to use?

The API is simple (replace _largepins_ with your Heroku app name):

* Use `https://largepins.herokuapp.com/?key=https://pinterest.com/hid3609481/feed.rss` to get the feed for user
  `hid3609481`.
  as the source.
* Use `https://largepins.herokuapp.com/?key=https://www.pinterest.com/hid3609481/home-decor-and-interior-design-ideas.rss/` to get the feed
  for the specific pin board. This will use
  `https://www.pinterest.com/hid3609481/home-decor-and-interior-design-ideas.rss` as the source.

