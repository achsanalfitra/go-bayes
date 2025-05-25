## Probability logic

Setting probabilities can be done in two ways,

* Marginal and conditional
* Joint


### Marginal and conditional (standard)

Let's evaluate the first method, marginal and conditional. This is the method of bayesian network.

The principle in this method is:

* **Each node is conditioned on their parents**: Node without a parent automatically has empty set as parents.
* **Each node understand their own states**: Marginal probability doesn't define each node states. They hold their own states.
* **Probability is set from each state**: Node probability is set based on their states. When a node has parents, the states are multiplied by the parent states.
* **Completion check is simple**: The question for the completion check is simple, does every node have their states defined the conditioned states which their parents have? Of course, check is now can be done arithmetically instead of calling everything.
* **Generally, there are two steps**: Setting: nodes are created, connected, and have their definition specified. Inference: when the context is clear, inference can happen. Probabilities are then calculated, states are calculated, reverse-connection can happen without breaking the structure. Adding a new node causes that node to be on the setting phase, not necessarily breaking the existing context.

#### Defining node structure

A node is identified by their own name. And, they hold their own states and probability spaces. They also acknowledge who their parents and their children are.

* **Name**: An identifier, similar to variable name.
* **States**: The domain where the states live.
* **Marginal**: The marginal probability space. This mimics the state exactly, but now has probability attached to it. When not defined or inferred, the complete states may not live here.
* **Conditional**: The conditional probability. It has all the states for each parent state.
* **Parents** (self explanatory)
* **Chiildren** (self explanatory)


