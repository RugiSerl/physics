# n-body simulation (optimised)
Famous simulation using Newton gravity formula $f_{a \to b} = G \frac {m_a m_b} {r^2} \overrightarrow{u_{ba}}$.

This version uses a quadtree to make approximations in the computations of forces, with the [Barnes-Hut algorithm](https://jheer.github.io/barnes-hut/)