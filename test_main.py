import unittest
from selenium import webdriver
from selenium.webdriver.common.by import By


class MyTestCase(unittest.TestCase):

    def setUp(self):
        option = webdriver.ChromeOptions()
        option.add_argument('--disable-dev-shm-usage')
        option.add_argument('--no-sandbox')
        option.add_argument('start-maximized')
        option.add_argument('--user-data-dir=/Users/oziomek/Library/Application Support/Google/Chrome/')
        option.add_argument('--disable-browser-side-navigation')
        web_driver = webdriver.Chrome(chrome_options=option, executable_path='/usr/local/bin/chromedriver')
        self.browser = web_driver

    def test_google_homepage(self):
        self.browser.implicitly_wait(10)
        self.browser.get('http://www.google.com/')
        header = self.browser.find_element(By.ID, 'site-header')
        self.assertEqual(header.is_displayed())

    def tearDown(self):
        self.browser.close()


if __name__ == '__main__':
    unittest.main()
