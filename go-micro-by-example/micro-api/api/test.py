import requests

def test_example_call():
    url = 'http://localhost:8080/example/call'
    params = {"name": "tony stack"}
    headers = {'head-1': 'I am a header'}
    r = requests.get(url, params=params, headers=headers)
    print("test_example_call:", r.status_code, r.text)

def test_example_foo_bar():
    url = 'http://localhost:8080/example/foo/bar'
    headers = {'Content-Type': 'application/json'}
    data = '{"name": "tony stack"}'
    r = requests.post(url, headers=headers, data=data)
    print('test_example_foo_bar:', r.status_code, r.text)

if __name__ == '__main__':
    test_example_call()
    test_example_foo_bar()