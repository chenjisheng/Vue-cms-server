#!/usr/bin/env python 
# coding:utf-8 
# @Time : 2019/3/27 17:57
# @Author : chenjisheng
# @File : cms_client.py
# @Mail : mail_maomao@163.com
import requests
import time

BASE_URL = "http://127.0.0.1:8080/v1"
def Add_swipe():
    number = 1
    url = BASE_URL + "/swipe"
    data = [
        {"img_url": "https://dpic3.tiankong.com/g0/rw/QJ7109236255.jpg?x-oss-process=style/240h"},
        {"img_url": "https://dpic1.tiankong.com/8m/lj/QJ6212733281.jpg?x-oss-process=style/240h"},
        {"img_url": "https://dpic.tiankong.com/58/lu/QJ9109162040.jpg?x-oss-process=style/240h"},
        {"img_url": "https://dpic1.tiankong.com/uv/jb/QJ6104512293.jpg?x-oss-process=style/240h"}
    ]
    res = requests.post(url,json=data)
    print(res.json())

def Add_newsList():
    url = BASE_URL + "/news"
    numbers = 10
    _time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    for i in range(numbers):
        data =  {"title":str(i)*3,
         "click":0,"url":"https://dpic.tiankong.com/58/lu/QJ9109162040.jpg?x-oss-process=style/240h",
         "add_time": _time,
         "content": {"content":str(i)*10,
                     "news_type":"media"},
         }
        res = requests.post(url,json=data)
        print(res.json())

def Add_comments():
    number = 10
    for news in range(1,number):
        url = BASE_URL + "/news/comments/" + str(news)
        for i in range(10):
            _time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
            data = {
                "comment":"public==" + str(i),
                "add_time": _time,
            }
            res = requests.post(url,json=data)
            print(res.json())


if __name__ == "__main__":
    Add_swipe()
    Add_newsList()
    Add_comments()
    pass