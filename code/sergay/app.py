import requests
import urllib3

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

url = 'http://vps-sni-site.site:5000/api/random'

def req():
    for i in range(50):
        r = requests.get(url, verify=False)
        data = r.json()
        text = data[str(data["id"])]
        print(f'серГЕЙ {text}')
req()