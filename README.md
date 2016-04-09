# Heroku echo telegram go bot
Small echo bot based on rockneurotiko/go-tgbot that can be hosted on Heroku

## How to install
Main steps are similar with https://devcenter.heroku.com/articles/getting-started-with-go .
But before pushing to heroku You have to set next env variables:
1. botToken - token received from Bot Father
2. hookURL - base url for hook. By default should be "https://application-heroku-domain"
