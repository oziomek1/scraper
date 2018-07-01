import urllib3
from bs4 import BeautifulSoup
from selenium import webdriver

"""
BASIC OTOMOTO.PL SCRAPPER.
OTOMOTO does not provide free/open API, it's only available for dealers and theirs developer

This app is based on beautifulsoup to parse HTML and selenium-webdriver for user-like interaction

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
Add tests
Move from python to node.js
Create client side in react/angular with user-friendly interface/graphs etc


There's a possibility that it'll be moved from otomoto to olx/allegro scrapping due to not ideal interface there..
"""

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


def start_chrome():
    option = webdriver.ChromeOptions()
    option.add_argument('--disable-dev-shm-usage')
    option.add_argument('--no-sandbox')
    option.add_argument('start-maximized')
    option.add_argument('--user-data-dir=/Users/oziomek/Library/Application Support/Google/Chrome/')
    option.add_argument('--disable-browser-side-navigation')
    web_driver = webdriver.Chrome(chrome_options=option, executable_path='/usr/local/bin/chromedriver')
    return web_driver


BASE_URL = 'https://www.otomoto.pl/'
TYPE = {
    'passenger': 'osobowe/',
    'delivery': 'dostawcze/',
    'motorcycle': 'motocykle-i-quady/',
    'truck': 'ciezarowe/',
    'construction': 'maszyny-budowlane/',
    'trailer': 'przyczepy/',
    'agro': 'maszyny-rolnicze'
}

QUERY_STRING_MARK = '?';
QUERY_START = 'search%5B';
QUERY_JOIN = '&';

SORT_TYPE = {
    'time': 'search%5Border%5D=created_at%3A',
    'price': 'search%5Border%5D=filter_float_price%3A',
    'mileage': 'search%5Border%5D=filter_float_mileage%3A',
    'power' : 'search%5Border%5D=filter_float_engine_power%3A'
}

SORT_TYPE_MODE = {
    'asc': 'asc',
    'dsc': 'desc'
}

ENGINE_CAPACITY = {
    'from': 'search%5Bfilter_float_engine_capacity%3Afrom%5D=',
    'to': 'search%5Bfilter_float_engine_capacity%3Ato%5D='
}

YEAR_SINCE = 'od-'
YEAR_TO = 'search%5Bfilter_float_year%3Ato%5D='


http = urllib3.PoolManager()

driver = start_chrome()
driver.get(BASE_URL)

driver.quit()
# make_box = soup.find('div', attrs={'data-name': 'filter_enum_make'}).select('option')
# makes = [x['value'] for x in make_box[1::]]


make_click = driver.find_element_by_link_text('Volkswagen')
driver.implicitly_wait(5)
make_click.click()

# print(makes)

print(driver.page_source)

