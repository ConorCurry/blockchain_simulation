#! /bin/python

from __future__ import division
import matplotlib.pyplot as plt
import matplotlib.mlab as mlab
import numpy as np
import scipy.stats as sp
from data_cleaning import generate_dataset
import math

def test_erlang(data):
    CRIT = 1.36 / math.sqrt(len(data))
    print "Critical Value: {}".format(CRIT)
    best_k = 0
    best_D = 1
    res = list()
    for k in range(1,10):
        res.append((sp.kstest(data, 'erlang', (k,))[0], k))
        print res[k-1]
        if res[k-1][0] < best_D:
            best_D = res[k-1][0]
            best_k = res[k-1][1]
    print 'best k parameter: {}\nD value: {}'.format(best_k, best_D)
    dat = [x[0] for x in res]
    print dat
    plt.plot([CRIT]*len(dat), 'r--', label='KS CRITICAL VALUE')
    plt.plot(dat, 'bH:', label='Erlang Shape Parameter D values')
    plt.show()

ds = generate_dataset()

times = sorted([block['time'] for block in ds.interarrival])
#times = [x/max(times) for x in times]
trimmed_times = times[::len(times)//100]
test_erlang(trimmed_times)
plt.hist(times,100)
plt.show()

tx_rate = sorted([block['num_tx']/block['time'] for block in ds.interarrival if block['time'] != 0 and block['num_tx'] != 0])
print 'mean: {}'.format(sum(tx_rate)/len(tx_rate))
print 'max: {}'.format(max(tx_rate))
print 'min: {}'.format(min(tx_rate))

trimmed_tx_rate = tx_rate[len(tx_rate)//25::len(tx_rate)//100]

test_erlang(trimmed_tx_rate)

plt.hist(tx_rate[:-len(tx_rate)//25], 100)
plt.show()


