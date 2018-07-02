import urllib3
from bs4 import BeautifulSoup
import constants
import re

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

make, model = 'volkswagen/', 'golf/'

http = urllib3.PoolManager()

response = http.request('GET', constants.BASE_URL + constants.TYPE['passenger'] + make + model)

soup = BeautifulSoup(response.data, 'html5lib')

offers = soup.find('div', {'class': 'offers list'})
offers = offers.findAll('article')

footer = soup.find('ul', {'class': 'om-pager rel'}).findAll('a')


# make_box = soup.find('div', attrs={'data-name': 'filter_enum_make'})
# make_box.select('option')
# makes = [x['value'] for x in make_box[1::]]

print('Number of offers in this page', len(offers))
print('Number of pages with this car:', re.findall(r'\d+', str(footer[-2]))[0])