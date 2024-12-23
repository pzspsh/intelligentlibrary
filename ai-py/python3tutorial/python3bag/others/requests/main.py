# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2023-12-07 09:53:00
@Author : pan
"""
import os
import re
import html
import requests
from lxml import etree


def nunusfcrawler(url: str):
    """ """
    response = requests.get(url=url, timeout=10)
    content = response.json()
    contentList = content["list"]
    for text in contentList:
        title = text["title"]
        urlpic = text["pic"]
        res = requests.get(url=urlpic, timeout=10)
        # contentpage = res.content.decode("utf8")
        # html = etree.HTML(contentpage)
        # article = html.xpath("//div/p/strong/text()")
        # article = html.xpath("//div/p/text()")
        # article = html.xpath("//div/p/text()||strong/text()")
        # print(article)
        contentdata = res.content.decode("utf-8")
        datasList = re.findall(
            r"<p>(.*?)</p><p><strong>(.*?)</strong>(.*?)</p>", contentdata
        )
        dirpath = "path/python3tutorial/python3bag/others/requests/"
        file = open(os.path.join(dirpath, title + ".txt"), "w", encoding="utf-8")
        for datas in datasList:
            for data in datas:
                data = html.unescape(data)
                file.write(data + "\n")


def spliceurl(urlpath: str, page: str):
    return urlpath.format(page)


def pingfan(id: str, page: str):
    urlhead = "https://www.nunusf.net/e/extend/bookpage/pages.php?id="
    urlpage = "&pageNum={0}&dz=asc&pageCount=1"
    response = requests.get(
        url=urlhead + id + spliceurl(urlpath=urlpage, page=page), timeout=10
    )
    content = response.json()
    contentList = content["list"]
    totalpage = content["totalPage"]
    for text in contentList:
        title = text["title"]
        urlpic = text["pic"]
        res = requests.get(url=urlpic, timeout=10)
        contentdata = res.content.decode("utf-8")
        datasList = re.findall(r"<p>(.*?)</p>", contentdata)
        dirpath = "path/others/requests/"
        file = open(os.path.join(dirpath, title + ".txt"), "w", encoding="utf-8")
        for datas in datasList:
            data = html.unescape(datas)
            print(data)
            file.write(data + "\n\n")
    if len(totalpage) > 0:
        pass


if __name__ == "__main__":
    # ch = "&#24773"
    # print(html.unescape(ch))
    # crawlerUrl = "https://www.nunusf.net/e/extend/bookpage/pages.php?id=16623&pageNum=0&dz=asc&pageCount=1"
    # nunusfcrawler(url=crawlerUrl)
    # crawlerUrl = "https://www.nunusf.net/e/extend/bookpage/pages.php?id=3541&pageNum=0&dz=asc&pageCount=1"
    pingfan(id="3541", page="0")
