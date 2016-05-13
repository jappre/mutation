
import urllib2

responses=urllib2.urlopen("http://www.baidu.com")
print responses.read()