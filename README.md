# FSI - Fast Streamlet Identifiers

FSI is a simple library that generates identifiers that respect Streamlet database format identifiers. It supports multithreading as it uses atomics.

## Format for Streamlet identifiers

22_random_base/10_incremental_numbers
Streamlet IDs are composed of the following: `22-length random base | "/" |  10-length incremental numbers`. An example Streamlet ID would be: `Hkh2bjNWQ4yxuMI++kTixA/0000000000`.

## Benchmarks

On average, it can achieve up to 6.5 million ops/sec on a single thread and 11 million ops/sec on 3 threads.
