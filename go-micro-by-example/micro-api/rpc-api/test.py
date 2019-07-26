# curl -H 'Content-Type: application/json' -d '{"name": "小小先"}' "http://localhost:8080/example/call"

import json
import requests

host = 'localhost:8080'
headers = {'Content-Type': 'application/json'}

def test_example_call():
    url = 'http://' + host + '/example/call'
    data = '{"name": "tony stack"}'
    r = requests.post(url, headers=headers, data=data)
    print('post {url}, status code: {status_code}'.format(url=url, status_code=r.status_code))
    print(r.text)

if __name__ == '__main__':
    test_example_call()
