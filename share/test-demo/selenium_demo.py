from time import sleep

from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By


class TestCase(object):
    def __init__(self):
        self.driver = webdriver.Chrome(
            service=Service(executable_path="/Users/pzxy/WorkSpace/driver/chromedriver")
        )
        self.driver.get("http://172.16.236.113:3007/")
        self.driver.maximize_window()
        sleep(1)

    def test_id(self):
        self.driver.find_element(By.ID, "normal_login_username").send_keys("admin")
        self.driver.find_element(By.ID, "normal_login_password").send_keys("Admin123!")
        sleep(3)
        self.driver.find_element(By.XPATH, '//*[@id="normal_login"]/div[3]/div/div/div/button').click()
        sleep(2)
        self.driver.find_element(By.XPATH, '//*[@id="root"]/section/aside/div/ul/li[2]/div').click()
        sleep(3)
        self.driver.quit()


if __name__ == '__main__':
    case = TestCase()
    case.test_id()
