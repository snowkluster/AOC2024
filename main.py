#!/usr/bin/env python3

from bs4 import BeautifulSoup
import subprocess
import requests
import datetime
import sys

time = datetime.datetime.now()

year = time.year
month = time.month
day = time.day
if (month != 12 or day > 25):
    print("Not december or day greater than 25, no advent of code")
    sys.exit()

result = subprocess.run(['bash', 'example.sh'], text=True)
page = requests.get(f"https://adventofcode.com/{year}/day/{day}")
soup = BeautifulSoup(page.content, 'html.parser')
article_content = soup.find('article', class_='day-desc')
with open(f"Day {day}/{day}.md", 'w', encoding='utf-8') as file:
    file.write(str(article_content))
