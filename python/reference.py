print 'Simple Assignment'
shoplist = ['apple','mango','caroot','banana']
mylist = shoplist

del shoplist[0]

print 'Shoplist is', shoplist
print 'mylist is', mylist

mylist = shoplist[:]
del mylist[0]

print 'Shoplist is', shoplist
print 'mylist is', mylist


