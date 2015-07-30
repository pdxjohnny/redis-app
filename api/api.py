#! /usr/bin/env python
import sys
import json
import urllib
try:
    # For Python 3.0 and later
    from urllib.request import urlopen
except ImportError:
    # Fall back to Python 2's urllib2
    from urllib2 import urlopen

host = "http://localhost:5000/"
USAGE = "Usage: python {} <function> paramaters"
# Help string

def fix_type(number):
    "Corrects the type of a sting"
    try:
        return json.loads(number)
    except ValueError:
        pass
    try:
        return float(number)
    except ValueError:
        pass
    try:
        return int(number)
    except ValueError:
        pass
    return number

def page(url):
    "Text of webpage"
    url = host + urllib.quote(url, safe="")
    response = urlopen(url)
    return response.read()

def post_json(url, jsonable_object):
    "Posts a json to a webpage"
    url = host + urllib.quote(url, safe="")
    req = Request(myurl)
    req.add_header('Content-Type', 'application/json; charset=utf-8')
    jsondata = json.dumps(jsonable_object)
    jsondataasbytes = jsondata.encode('UTF-8')
    req.add_header('Content-Length', len(jsondataasbytes))
    response = urlopen(req, jsondataasbytes)
    return response.read()

def get(key):
    "Retrives the value of the key"
    res = page(key)
    res = fix_type(res)
    return res

def set(key, value):
    "Sets the key to the value"
    res = page(u"set/" + str(key) + u"/" + str(value))
    res = fix_type(res)
    return res

def main():
    "CLI"
    if len(sys.argv) < 2:
        print(USAGE.format(sys.argv[0]))
        print("")
        help(sys.modules[__name__])
    else:
        try:
            function = sys.argv[1]
            function = getattr(sys.modules[__name__], function)
            res = function(*sys.argv[2:])
            print(res)
        except Exception as error:
            print(error)

if __name__ == "__main__":
    main()
