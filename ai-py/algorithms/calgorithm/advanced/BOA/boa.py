# -*- encoding: utf-8 -*-
"""
@File   : boa.py
@Time   : 2024-05-23 10:49:02
@Author : pan
"""
import numpy as np


class ButterflyOptimizationAlgorithm:
    def __init__(self, num_butterflies, dimensions, obj_func, lb, ub, max_iter):
        self.num_butterflies = num_butterflies
        self.dimensions = dimensions
        self.obj_func = obj_func
        self.lb = np.array(lb)
        self.ub = np.array(ub)
        self.max_iter = max_iter
        self.butterflies = (
            np.random.uniform(0, 1, (self.num_butterflies, self.dimensions))
            * (self.ub - self.lb)
            + self.lb
        )
        self.best_position = self.butterflies[0]
        self.best_score = float("inf")

    def update_butterfly_position(self, butterfly, best_position, dimension):
        r1, r2 = np.random.rand(2)
        c1, c2 = 2, 1  # Constants controlling the attraction and repulsion forces
        diff = best_position[dimension] - butterfly[dimension]
        new_pos = (
            butterfly[dimension]
            + (c1 * np.exp(-r1 * abs(diff)) - c2 * np.exp(-r2 * abs(diff))) * diff
        )
        return max(min(new_pos, self.ub[dimension]), self.lb[dimension])

    def optimize(self):
        for iter_count in range(self.max_iter):
            for butterfly_index, butterfly in enumerate(self.butterflies):
                for dimension in range(self.dimensions):
                    new_pos = self.update_butterfly_position(
                        butterfly, self.best_position, dimension
                    )
                    butterfly[dimension] = new_pos

                    # Evaluate the new position
                    score = self.obj_func(butterfly)
                    if score < self.best_score:
                        self.best_score = score
                        self.best_position = butterfly

            # Optional: Print progress
            print(
                f"Iteration {iter_count+1}/{self.max_iter}, Best Score: {self.best_score}"
            )

        return self.best_position, self.best_score


# Example usage:
def objective_function(x):
    # Your optimization problem's objective function goes here
    return np.sum(x**2)  # For example, minimizing the sum of squares


if __name__ == "__main__":
    lb = [-10] * 2  # Lower bounds for each dimension
    ub = [10] * 2  # Upper bounds for each dimension
    num_butterflies = 10
    dimensions = len(lb)
    max_iter = 100

    boa = ButterflyOptimizationAlgorithm(
        num_butterflies, dimensions, objective_function, lb, ub, max_iter
    )
    best_position, best_score = boa.optimize()
    print(f"Best Position: {best_position}, Best Score: {best_score}")
