# -*- encoding: utf-8 -*-
'''
@File   : crawler.py
@Time   : 2024-06-06 15:01:27
@Author : pan
'''
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


def CrawlerShowMoreClick(Url: str): # 实现Show More...按钮点击
    """
    实现Show More...按钮点击
    :param Url: 要抓取的网页URL
    :return: None
    """
    # 初始化Chrome浏览器
    driver = webdriver.Chrome()

    # 打开目标网页
    driver.get(Url)
    try:
        # 假设我们有一个网页，其中有一个"Show more"按钮，它的HTML代码可能看起来像这样：
        # <button id="load-more-btn" class="show-more-button">Show more</button>
        # 在这个例子中，按钮有一个唯一的id属性load-more-btn和一个class属性show-more-button。
        # 我们可以使用这些属性来定位并点击这个按钮。

        # 确定"Show more"按钮的定位方式，例如通过id、class、xpath等
        # show_more_button = driver.find_element_by_id('showMoreButtonId')

        # 等待"Show more"按钮出现并变得可点击
        show_more_button = WebDriverWait(driver, 10).until(EC.element_to_be_clickable((By.ID, 'load-more-btn')))

        # 点击"Show more"按钮
        show_more_button.click()

        # 这里可以添加代码来处理加载出来的更多内容，例如等待内容加载完成后再进行抓取
        # ...

    finally:
        # 关闭浏览器
        driver.quit()


if __name__ == "__main__":
    Url = "http://example.com/page-with-show-more"
    CrawlerShowMoreClick(Url=Url)
