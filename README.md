This project is wholely based on the github.com/relab/hotstuff project, the main difference is that we add the pbeegees consensus package, which is a noval BFT algorithm I proposed that is an improvement to the BeeGees BFT algorithm[1]. However, the original framework that the relab project provides is not perfectly suitable for pBeeGees (and BeeGees), so there are lot of trivial modification in everywhere of the original project. The detailed specification of this project and the pBeeGees algorithm is still in progress.



[1]Neil Giridharan, Florian Suri-Payer, Matthew Ding, Heidi Howard, Ittai Abraham, and Natacha Crooks. 2023. BeeGees: Stayin' Alive in Chained BFT. In Proceedings of the 2023 ACM Symposium on Principles of Distributed Computing (PODC '23). Association for Computing Machinery, New York, NY, USA, 233–243. https://doi.org/10.1145/3583668.3594572
