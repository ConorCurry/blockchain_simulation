from data_cleaning import generate_dataset

from scipy import stats
import numpy

def main():
    dataset = generate_dataset()

    def times_only(dataset):
        return map(lambda x: x['time'], dataset)

    def normalize(dataset):
        mean = sum(dataset) / len(dataset)
        standard_dev = numpy.std(dataset)
        normal = lambda x: (x - mean)/standard_dev
        return map(normal, dataset)

    def fit_norm(vect):
        norm_vect = normalize(vect)
        print 'Performing the KS test on normal data:'
        dstat, p = stats.kstest(norm_vect, 'norm')
        print 'D-Statistic:\t{}'.format(dstat)
        print 'P Value:\t\t{}'.format(p)
        print '----------------------------'

    def mean(vect):
        return sum(vect) / len(vect)

    def fit_erlang(vect):
        pass

    def fit_exponential(vect):
        exponential_dist = stats.expon(scale=mean(vect)) # lambda = 1 / mean
        print 'Performing the KS test on exponential data:'
        dstat, p = stats.kstest(vect, exponential_dist.pdf)
        print 'D-Statistic:\t{}'.format(dstat)
        print 'P Value:\t\t{}'.format(p)
        print '----------------------------'

    interarrival = times_only(dataset.t0)
    fit_norm(interarrival[:100])
    fit_exponential(interarrival[:100])
    

if __name__ == '__main__':
    main()
