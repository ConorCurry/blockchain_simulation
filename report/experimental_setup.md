# Experiment Setup and Design

The purpose of the experiment was to test the effects of a modified block size on transaction fees. As of late, Bitcoin has been suffering from an increase of transaction fees stemming from a dearth in the number of available blocks. The correlation between runs of full blocks and transaction fees was measured, and there was determined to be a moderate correlation (r = 0.5). Thus, reducing the number of runs of full blocks should decrease the cost to perform a transaction for the average Bitcoin owner.

Our hypothesis were as follows:

    H0:     Adjusting the block will not significantly affect fees.
    Ha:     Changing the block size will significantly affect the fee cost.
   
In our simulation, the fee will increase if a waiting queue builds up with more elements than a block can hold. When that occurs, transactions are taken off of the queue according to their fee price, instead of the expected first-in first-out ordering. Given parameters determined via input validated data from the blockchain, the likelihood of runs of full blocks, likely cost estimates per byte on the blockchain, and overall the possible effects of increasing the block size were sought to be determined. 

# Input Modeling

To do the input modeling, data validation and to construct graphs we utilized the Python libraries scipy, numpy, and matplotlib.

![Figure 1. Block Interarrival Times](figures/interarrival-exponential-dist.png)

The Kolmogorov-Smirnov test was used to model the interarrival time of the blocks. This was found to be exponential.

![Figure 2. Transaction Arrival Times](figures/transaction-rate-histogram.png)

The number of transactions per block was found to follow the Poisson arrival, also via the Kolmogorov-Smirnov test.

![Figure 3. Kolmogrov-Smirnov of Erlang Distributions: Y-axis values represent D-Statistic, x-axis values represent K, and should start at x = 1. Thus, k = 3 instead of 2, as shown](figures/erlang-shape-parameters-ks-tests.png)

In addition, we found that the rate of transactions per block followed the Erlang distribution. This was also calculated with the Kolmogrov-Smirnov test.

# Simulation

To run the simulation, certain simplifying decisions were made. The actual transaction arrival times were not recorded, instead a model was created and validated. This was done due to the nature of the data not being flat. 

Based on the input modeling done, there were several parameters to be calculated for the simulation. The block arrival lambda value, alpha and beta values to find the transaction arrival lambda, the alpha and beta values for finding the fee cost, and maximum number of transactions per block had to all be determined. The block arrival lambda was found via the mean interarrival time using Python's libraries, 0.001745. The beta value could be determined from the Kolmogorov-Smirnov test for the Erlang distribution, arriving at a K value of degree 3. The beta value was found to be 1. For the fee cost, the apha value was found to be 7, and beta value 1200. The maximum transactions for block was 2048.

Table 1. Parameter values based on true data collected.

| Parameter                  | Value    |
|:--------------------------:|:--------:|
| Block Arrival Lambda       | 0.001745 |
| Transaction Alpha          | 3.0      |
| Transaction Beta           | 1.0      |
| Fee Alpha                  | 7.0      |
| Fee Beta                   | 1200.0   |
| Maximum Transactions       | 2048     |

The standard block size, 2048 kb served as our control group. Our experimental groups were powers of 2 on the standard size, as those have been proposed changes to the block size. All other factors were kept constant between simulation runs.

Based on these parameters, the simulation was constructed around the use of a priority queue to collect blocks. For each block an arrival time was collected. Within this block, each transaction collected had an associated arrival time and fee. Based on the block arrival time modeling, the amount of time passed was measured, by the end of the simulated time collecting approximately 10,000 blocks.

Table 2. Comparison of the control group to the experimental 2x group.

| T-Statistic | P-Value     |
|:-----------:|:-----------:|
| 8.0835      | 9.1085      |

Table 3. Comparison of the control group to the experimental 0.5x group.

| T-Statistic | P-Value     |
|:-----------:|:-----------:|
| -1.3841     | 0.1680      |

The block arrival generator used the same random number seed, so all blocks arrived at the same time across all experiments. As result, a Paired T Test was run between the control and experimental data produced. The calculated p-value was compared at the 0.01 significance level. Based on these calculations, the 2x Group is statistically significant, while the 0.5x group is not statistically significant. 
