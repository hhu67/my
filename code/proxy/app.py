import requests
import urllib3

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

urlpr='http://vps-sni-site.site:507/api/random'
url='https://httpbin.org/ip'

r = requests.get(urlpr, verify=False)
data = r.json()
text = data["proxy"]

proxies = {
    'http': text,
    'https': text
}

j = requests.get(url, proxies=proxies, verify=False)
data2 = j.json()
text2 = data2["origin"]

f = requests.get(url, verify=False)
data3 = f.json()
text3 = data3['origin']

print(f'Прокси: {text2}\nБез прокси: {text3}')