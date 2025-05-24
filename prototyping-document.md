## Probability logic

Setting probabilities can be done in two ways,

* Marginal and conditional
* Joint

Specifying both **conditional** and **joint** may imply two existing dependencies, this is not advised.
The formula suggests that joint doesn't necessarily mean independence when it doesn't equal the product of marginals.
However, you can't simply specify all marginals with all known joints. In large nodes, this is difficult to track.
Therefore, a sensible way to design this is whether to describe all the events upfront (joint) or granularly using conditional/marginal.