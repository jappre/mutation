shoplist = ['apple','mango','carrot','banana']
print 'I have', len(shoplist), 'items to purchase.'
print 'These items are:',
for item in shoplist:
    print item,

print '\nI also have to buy rice.'
shoplist.append('rice')
print 'My shopping list is now', shoplist
