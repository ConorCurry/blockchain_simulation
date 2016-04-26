#! /usr/bin/python

from __future__ import division
import matplotlib.pyplot as plt
import matplotlib.mlab as mlab
import numpy as np
import scipy.stats as sp
from itertools import groupby
from data_cleaning import generate_dataset
import math
from collections import defaultdict

mean = lambda x : sum(x)/len(x)

def calcAvgFee(blockList):
    tot_fee = sum([b['satoshi_fee'] for b in blockList])
    tot_size = sum([b['block_size'] for b in blockList])

    fee_per_byte = 0
    if tot_size != 0:
        fee_per_byte = tot_fee // tot_size

    return fee_per_byte
    

def runLenFees(fullSize=975000):
    ds = generate_dataset()
    
    grouped, not_full = list(), list()
    r = False
    total_tx = list()
    for b in ds.interarrival:
        total_tx.extend([b['satoshi_fee']//b['num_tx']]*b['num_tx'])
        if b['block_size'] <= fullSize:
            not_full.append(b)
            r = False
        elif b['block_size'] > fullSize:
            if not r:
                grouped.append([b])
                r = True
            else:
                grouped[-1].append(b)


    mean = sum(total_tx)/len(total_tx)
    print 'mean tx fee: {}'.format(mean)
    fig = plt.figure()
    fig.suptitle('Transaction Fee PDF/Histogram Fit')
    plt.hist(sorted(total_tx)[:-len(total_tx)//100],300, normed=True)
    plt.xlabel('Satoshi Tx Fee')
    plt.ylabel('Probability')
    plt.axvline(x=mean, color='red')
    rv = sp.erlang(7, loc=11500, scale=1200)
    x = np.linspace(0,50000)
    #plt.plot(x, rv.pdf(x))
    #plt.show()
    
    l1 = [i/len(total_tx) for i in range(len(total_tx)-(len(total_tx)//50)-1)]
    l2 = [e for e in sorted(total_tx)[:-len(total_tx)//50]]
    print len(l1), len(l2)
    fig.suptitle('Sorted Fees vs. Erlang CDF')
    plt.xlabel('Satoshi Tx Fee')
    plt.ylabel('Probablility')

    plt.plot(x, rv.cdf(x), color='red')
    plt.plot(l2, l1, color='blue')
    plt.show()

    dic = defaultdict(list)
    for blocklist in grouped:
        dic[len(blocklist)].extend(blocklist)

    avgFeeDic = dict()
    for key in dic:
        avgFeeDic[key] = calcAvgFee(dic[key])
    avgFeeDic[0] = calcAvgFee(not_full)
    return avgFeeDic

def feeVsRunScatter():
    x = runLenFees()
    print 'Run Length Fee Dict: {}'.format(x)
    print 'Correlation between full block run length and fee amounts: {}'.format(np.corrcoef(x.keys(), x.values())[0][1])
    plt.scatter(x.keys(), x.values())
    plt.show()

feeVsRunScatter()
