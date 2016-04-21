#!/bin/python

import csv

def get_csv_data():
    with open("../data/output.csv") as f:
        reader = csv.DictReader(f)
        return [x for x in reader]

def compute_mean(block_list, f):
    return sum([f(x) for x in block_list]) / len(block_list)

def parse_as_ints(block_list):
    result = []
    for block in block_list:
        result.append({
            'num_tx':       int(block['num_tx']),
            'satoshi_fee':  int(block['satoshi_fee']),
            'time':         int(block['time']),
            'block_size':   int(block['block_size']),
            'hash':         block['hash']
        })
    return result

def transform_offset_to_zero(block_list):
    result = []

    zero = block_list[0]['time']
    for i, block in enumerate(block_list):

        result.append({
            'num_tx':       block['num_tx'],
            'satoshi_fee':  block['satoshi_fee'],
            'time':         block['time'] - zero,
            'block_size':   block['block_size'],
            'hash':         block['hash']
        })
    return result

def calculate_interarrival_times(block_list):
    result = []
    for i, block in enumerate(block_list):
        if i == 0: 
            continue
        previous = block_list[i - 1]
        result.append({
            'num_tx':       block['num_tx'],
            'satoshi_fee':  block['satoshi_fee'],
            'time':         block['time'] - previous['time'],
            'block_size':   block['block_size'],
            'hash':         block['hash']
        })
    return result

def find_all_negative_arrivals(block_list):
    for block in block_list:
        if block['time'] < 0:
            print 'Hash: {0}, time: {1}'.format(block['hash'], block['time'])

class DataSet(object):

    def __init__(self, sorted=None, interarrival=None, t0=None, mean_time=-1, mean_interarrival_time=-1):
        self.sorted = sorted
        self.interarrival=interarrival
        self.t0 = t0
        self.mean_time = mean_time
        self.mean_interarrival_time = interarrival_time


def generate_dataset():
    block_list_as_str       = get_csv_data()
    block_list              = parse_as_ints(block_list_as_str)

    sorted_block_list       = sorted(block_list, key=lambda x: x['time'])

    block_list_t0           = transform_offset_to_zero(sorted_block_list)
    block_list_interarrival = calculate_interarrival_times(sorted_block_list)

    mean_time = compute_mean(sorted_block_list, lambda x: x['time'])
    mean_interarrival_time = compute_mean(block_list_interarrival, lambda x: x['time'])

    return DataSet(
            sorted=sorted_block_list,
            interarrival=block_list_interarrival,
            t0=block_list_t0,
            mean_time=mean_time,
            mean_interarrival_time=mean_interarrival_time
    )

def main(): 
    pass   


if __name__ == '__main__':
    main()
