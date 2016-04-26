#! /usr/bin/python3

import requests
import json
from time import localtime, strftime
import csv
import time

# note that max block size is 1MB
# account for this in deciding how many blocks to retrieve
# newer blocks will be larger on average
# determine total fees for each block, and number of transactions
# this will help find expected value, for calculation of lambda

fieldnames = ['num_tx', 'satoshi_fee', 'time', 'block_size', 'hash']

def main():
    url_template = 'https://blockchain.info/block-height/{0}?format=json'
    start = 406861 - 10000
    desired_blocks = [x + start for x in range(10000)]

    with open('/output.csv', 'w') as f:

        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()

        now = time.clock()

        for block in desired_blocks:

            def request_block(block):
                url = url_template.format(block)
                resp = requests.get(url)
                block_dict = marshal_resp(resp)
                writer.writerow(block_dict)

            def block_if_rate_limited(now, index, f):
                if index % 100 == 0:
                    print 'Processed a hundred blocks'
                    f.flush()
                    if time.clock() - now < 62:
                        time.sleep(62 - (time.clock() - now))
                    now = time.clock()

                elif index % 50 == 0:
                    print 'Process block number {0}'.format(index)

            request_block(block)
            block_if_rate_limited(now, block, f)


def marshal_resp(resp):
    block_dat = resp.json()['blocks'][0]
    important_data = {
        fieldnames[0]:  block_dat['n_tx'],
        fieldnames[1]:  block_dat['fee'],
        fieldnames[2]:  block_dat['time'],
        fieldnames[3]:  block_dat['size'], 
        fieldnames[4]:  block_dat['hash']
    }
    return important_data

if __name__ == '__main__':
    main()
