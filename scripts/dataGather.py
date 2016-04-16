#! /usr/bin/python3

import requests
import json
from time import localtime, strftime

url_template = 'https://blockchain.info/block-height/{}?format=json'


start = 400001

#note that max block size is 1MB
#account for this in deciding how many blocks to retrieve
#newer blocks will be larger on average
desired_blocks = [x + start for x in range(100)]


#determine total fees for each block, and number of transactions
#this will help find expected value, for calculation of lambda

for b in desired_blocks:
    url = url_template.format(b)

    resp = requests.get(url)
    block_dat = resp.json()['blocks'][0]
    important_data = {'n_tx':block_dat['n_tx'],'satoshi_fee':block_dat['fee'],'t':block_dat['time'],'block_size':block_dat['size']}

    print("{} -- {}".format(strftime('%Y-%m-%d %H:%M:%S', localtime(important_data['t'])), repr(important_data.items())))
    #print(resp.json()['blocks'][0].keys())



