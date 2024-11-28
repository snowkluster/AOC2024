#!/usr/bin/env python3

from bs4 import BeautifulSoup
from dotenv import load_dotenv
import requests
import datetime
import sys
import os

load_dotenv()
COOKIE = os.getenv('COOKIE')
time = datetime.datetime.now()
s = requests.Session()

cookies = {
    "session" : COOKIE
}
s.cookies.update(cookies)

year = time.year
month = time.month
day = time.day
if (month != 12 or day > 25):
    print("Not december or day greater than 25, no advent of code")
    sys.exit()

input = s.get(f"https://adventofcode.com/{year}/day/{day}/input")
page = s.get(f"https://adventofcode.com/{year}/day/{day}")
soup = BeautifulSoup(page.content, 'html.parser')
article_content = soup.find('article', class_='day-desc')
with open(f"Day {day}/{day}.md", 'w', encoding='utf-8') as file:
    file.write(str(article_content))
with open(f"Day {day}/input.txt",'w', encoding='utf=8') as file:
    file.write(input.text)