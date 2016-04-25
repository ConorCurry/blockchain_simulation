# Approach

Before we could conduct experimentation, we first had to develop a model for the blockchain and then implement it. This proces was composed of 4 phases. 

0. We collected public information from the block chain. 
0. We modeled the data we collected. 
0. Next, we generated a simulation based on the model we generated.
0. Finally, we validated our model against the empirical data.

## Data Collection

To collect the data, we queried an API for the transaction history of the block chain. Because the block chain is open source, anyone can create a node of the distributed database and join the system. This means that all data in the blockchain is public. This is in contrast to traditional equities markets, where pricing information and transactions are not necessarily released instantaniously; markets like _dark pools_ may report transactions well after they've occurred, and companies like Bloomburgh license a stream of ticker information to customers who want to see prices in real-time. Instead, we were able to use a free public API to collect the history of the blockchain.

We used [Blockchain.info](blockchain.info) to collect the transaction information. Blockchain.info exposes an [API](https://blockchain.info/api) for aggregating information on recorded blocks. We collected metadata for 10000 blocks spanning the course of a week. 

We collected only contiguous blocks, instead of a random sampling of blocks, because a random sampling could introduce much more bias. Supposing you sample randomly, you may end up with a sample consisting of a high density of transactions at a certain point in the day; time of day could influence transaction rate, depending on how many Bitcoin users are at that time of day. Additionally, because our study requires that we find the duration of time for which there is a backlog of transations (i.e. the amount of time in between non-full blocks) we needed continuous data; if the blocks were non-contiguous, then it's possible that a intermediate block could be non-full, resulting in an error. (Here, we use **contiguous** to refer to blocks, as they are sequential, discrete entities, and __continuous__ to refer to data, which has an associated timestamp and thus is non-discrete.)

In order to avoid rate limiting, we calculated the acceptible number of requests per minute. When executing our collection script, we included a function that would sleep the main thread if our number of requests in the last minute exceeded our rate. Then, we parsed the JSON response and transformed the data to CSV format before storing the results on disk. The data we collected represented ____ KB of storage, and 10000 outbound network requests. For each block, we were only interested in five attributes: number of transactions, fee cost, the time at which the block was added to the blockchain, and the block size. All remaining attributes are computed properties derived from the above.

## Input Modeling

Once we had collected the data for our model, we began analysis to determine the distribution of the arrival of blocks in the blockchain. Our goal was to be able to determine how frequently new blocks are added to the chain. Each block is added to the chain following a random process: a block is added once a Proof of Work procedure has been completed. 

A Proof of Work procedure is a problem generally difficult and time consuming to solve, but easily varifiable. This done in part for rate-limiting purposes, as a block can be generated only so fast as the Proof of Work is completed. The more difficult the Proof of Work, the longer amount of time between generating blocks, so gauging the difficulty helps to keep the generation rate constant. In the case of Bitcoin, the Proof of Work is computing the preimage of a hash, based on the SHA-256 algorithm. This involves taking two strings and hashing them together to produce a result with a set number of leading zeros. If the amount of leading zeros on the result meet a given threshold, the block is solved. Otherwise, a different hash string must be input.

---- Explain:
---- The Proof of Work is computing the preimage of a hash. SHA 256. Double hashing.

Once the interarrival times for each block were modeled, the Kolmogrov-Smirnov test on the resulting data. It tested for both the normal distribution as well as the exponential distribution. In addition, the rate of the transaction times was modeled, as well as the the number of transactions per block.

## Simulation

To run the simulation, arrival times were generated for each block. A lambda value was generated for the arrival times of transactions per block. Two separate simulations were run, one as the control group using the original block size of 1 MB, the other being the experimental group using a doubled block size. Running the control group is necessary to ensure that the simulation was modeled correctly after the real data, and that the experimental data is being compared in an equal environment. The experimental group was tested specifically to investigate if the new block size could saturate the number of incoming transactions per rate time, looking closely at the backlog run of transaction data.

## Model Validation

For the model validation, the control group data was compared to the real data previously collected. The results of the control group were then compared to those of the experimental group. Calculations and results focused most on comparing the backlog run data, as this would tell whether an increase in block size had any impact on the amount of transactions waiting at any given time. 
