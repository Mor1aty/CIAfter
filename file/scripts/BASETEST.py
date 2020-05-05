import os
import sys
import time

from appium import webdriver

apk_path = os.path.abspath(os.path.join(os.path.dirname(__file__), "."))  # 获取当前项目的根路径
desired_caps = {}
desired_caps['platformName'] = sys.argv[1]
desired_caps['platformVersion'] = sys.argv[2]  # 手机系统版本
desired_caps['deviceName'] = sys.argv[3]  # 刚才的devicename
desired_caps['app'] = apk_path + '\\' + sys.argv[4]

driver = webdriver.Remote('http://localhost:4723/wd/hub', desired_caps)  # 运行该脚本desired_caps

result = sys.argv[5]
package_name = 'com.example.electric_chamberlain_app'
time.sleep(2)
driver.get_screenshot_as_file(result + '\\1.png')
time.sleep(5)
driver.get_screenshot_as_file(result + '\\2.png')
time.sleep(5)
driver.get_screenshot_as_file(result + '\\3.png')
time.sleep(5)
driver.remove_app(package_name)
driver.quit()
print(0)
