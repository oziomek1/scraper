# scraper

BASIC OTOMOTO.PL SCRAPPER.

OTOMOTO does not provide free/open API, it's only available for dealers and theirs developer

~This app is based on beautifulsoup to parse HTML and selenium-webdriver for user-like interaction~
#### Currently the app has been changed from python to golang 

This scraper should be able to collect data for different:
- makes
- models
- versions
- fuel type
- price range
- mileage
- year
- engine capacity & power
- gearbox type
- additional equipment

The aim is to collect data and apply machine learning to find various patterns.

Further plans:
* Add tests
* ~Move from python to node.js~
* Create client side in react/angular with user-friendly interface/graphs etc

### In case of any build problems, Travis is constantly looking after
<img src="https://travis-ci.org/oziomek1/scraper.svg?branch=master"/>

There's a possibility that it'll be moved from otomoto to olx/allegro scrapping due to not ideal interface there..


### chromedriver 2.40 got some issue and cannot provide proper behaviour of Chrome webbrowser
