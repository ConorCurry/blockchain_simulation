from __future__ import division
from scipy import stats
import json

def main():
    file1 = 'data/output1.json'
    file2 = 'data/output2.json'
    file05 = 'data/output05.json'

    control_fees = []
    exp2_fees = []
    exp05_fees = []

    with open(file1) as f:
        contents = f.read()
        data = json.loads(contents)
        blocks = data['blocks']
        txs = [ block['transactions'] for block in blocks ]
        control_fees = [ float(tx[0]['fee']) for tx in txs ]
        
    with open(file2) as f:
        contents = f.read()
        data = json.loads(contents)
        blocks = data['blocks']
        txs = [ block['transactions'] for block in blocks ]
        exp2_fees = [ float(tx[0]['fee']) for tx in txs ]

    with open(file05) as f:
        contents = f.read()
        data = json.loads(contents)
        blocks = data['blocks']
        txs = [ block['transactions'] for block in blocks ]
        exp05_fees = [ float(tx[0]['fee']) for tx in txs ]

    t, p = stats.ttest_rel(control_fees, exp2_fees)
    print '--------------------------------------------'
    print 'Comparing the control group to the 2x Group |'
    print 'T Stat: {}'.format(t)
    print 'P value: {}'.format(p)
    print '--------------------------------------------'

    t, p = stats.ttest_rel(control_fees, exp05_fees)
    print '--------------------------------------------'
    print 'Comparing the control group to the 0.5x Group |'
    print 'T Stat: {}'.format(t)
    print 'P value: {}'.format(p)
    print '--------------------------------------------'

    control_mean = sum(control_fees) / len(control_fees)
    exp2_mean = sum(exp2_fees) / len(exp2_fees)

    print 'Control mean: {}'.format(control_mean)
    print 'Exp2 Mean: {}'.format(exp2_mean)


if __name__ == '__main__':
    main()
