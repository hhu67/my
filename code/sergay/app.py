import requests

url = "http://vps-sni-site.site:5000/api/random"

r = requests.get(url, verify=False)
data = r.json()
text = data[str(data["id"])]

print(f"серГЕЙ {text}\nбэээээ")