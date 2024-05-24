# -*- encoding: utf-8 -*-
"""
@File   : pso_demo3.py
@Time   : 2024-05-23 10:34:05
@Author : pan
"""
import numpy as np
from sklearn.datasets import make_regression
from sklearn.neural_network import MLPRegressor


class Particle:
    def __init__(self, dimension):
        self.position = np.random.uniform(-10, 10, dimension)
        self.velocity = np.random.uniform(-1, 1, dimension)
        self.best_position = np.copy(self.position)
        self.best_score = float("inf")


class PSO:
    def __init__(self, num_particles, dimension, alpha=0.5, beta=0.8, gamma=0.9):
        self.num_particles = num_particles
        self.particles = [Particle(dimension) for _ in range(num_particles)]
        self.g_best_position = np.random.uniform(-10, 10, dimension)
        self.g_best_score = float("inf")
        self.alpha = alpha
        self.beta = beta
        self.gamma = gamma

    def optimize(self, function, max_iter):
        for _ in range(max_iter):
            for particle in self.particles:
                fitness = function(particle.position)
                if fitness < particle.best_score:
                    particle.best_score = fitness
                    particle.best_position = particle.position.copy()

                if fitness < self.g_best_score:
                    self.g_best_score = fitness
                    self.g_best_position = particle.position.copy()

            for particle in self.particles:
                inertia = self.alpha * particle.velocity
                personal_attraction = (
                    self.beta
                    * np.random.random()
                    * (particle.best_position - particle.position)
                )
                global_attraction = (
                    self.gamma
                    * np.random.random()
                    * (self.g_best_position - particle.position)
                )
                particle.velocity = inertia + personal_attraction + global_attraction
                particle.position += particle.velocity

        return self.g_best_position, self.g_best_score


# 创建一个模拟数据集
X, y = make_regression(n_samples=100, n_features=2, noise=0.1)


# 定义一个用于评估粒子的适应度的函数
def fitness(position):
    # 将位置向量重新塑形为权重和偏置
    hidden_layer_weights = position[:20].reshape(10, 2)
    hidden_layer_bias = position[20:30]
    output_weights = position[30:40]
    output_bias = position[40]

    model = MLPRegressor(hidden_layer_sizes=(10,), max_iter=1, warm_start=True)
    model.coefs_ = [hidden_layer_weights, output_weights.reshape(-1, 1)]
    model.intercepts_ = [hidden_layer_bias, np.array([output_bias])]

    model.partial_fit(X, y)
    predictions = model.predict(X)
    mse = ((predictions - y) ** 2).mean()
    return mse


if __name__ == "__main__":
    # 初始化PSO并优化
    dimension = 41  # 20权重+10偏置+10权重+1偏置
    pso = PSO(num_particles=30, dimension=dimension)
    best_position, best_score = pso.optimize(fitness, max_iter=1000)
    print(f"Best MSE: {best_score}")
