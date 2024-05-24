# -*- encoding: utf-8 -*-
"""
@File   : pso_demo.py
@Time   : 2024-05-23 10:29:31
@Author : pan
"""
import numpy as np


class Particle:
    def __init__(self, dimension, bounds):
        # 粒子的位置和速度初始化
        self.position = np.random.uniform(low=bounds[0], high=bounds[1], size=dimension)
        self.velocity = np.random.uniform(low=-1, high=1, size=dimension)
        # 粒子的最佳位置和最佳适应度值初始化
        self.best_position = np.copy(self.position)
        self.best_value = float("inf")


def objective_function(x):
    # 目标函数示例：求解最小化问题
    return np.sum(x**2)


def update_velocity(particles, gbest_position, w=0.5, c1=2.0, c2=2.0):
    # 更新每个粒子的速度
    for particle in particles:
        r1, r2 = np.random.rand(), np.random.rand()
        cognitive_velocity = c1 * r1 * (particle.best_position - particle.position)
        social_velocity = c2 * r2 * (gbest_position - particle.position)
        particle.velocity = w * particle.velocity + cognitive_velocity + social_velocity


def update_position(particles, bounds):
    # 更新每个粒子的位置，并确保位置在边界内
    for particle in particles:
        particle.position += particle.velocity
        particle.position = np.clip(particle.position, bounds[0], bounds[1])


def pso(objective_function, bounds, num_particles, max_iter, dimension):
    # 初始化粒子群
    particles = [Particle(dimension, bounds) for _ in range(num_particles)]
    gbest_value = float("inf")
    gbest_position = None

    for i in range(max_iter):
        for particle in particles:
            # 评估粒子的当前适应度
            current_value = objective_function(particle.position)
            # 更新个体最优
            if current_value < particle.best_value:
                particle.best_value = current_value
                particle.best_position = particle.position
            # 更新全局最优
            if current_value < gbest_value:
                gbest_value = current_value
                gbest_position = particle.position

        # 打印当前迭代次数和全局最优适应度值
        print(f"迭代次数 {i+1}/{max_iter}, 全局最优适应度值: {gbest_value}")

        # 根据全局最优更新粒子的速度和位置
        update_velocity(particles, gbest_position)
        update_position(particles, bounds)

    return gbest_position, gbest_value


if __name__ == "__main__":
    # 设置参数和运行PSO算法
    dimension = 2  # 解的维度
    bounds = np.array([[-10, 10], [-10, 10]])  # 解空间的边界
    num_particles = 30  # 粒子数量
    max_iter = 100  # 最大迭代次数
    gbest_position, gbest_value = pso(
        objective_function, bounds, num_particles, max_iter, dimension
    )
    print(f"找到的最优解: {gbest_position}, 最优适应度值: {gbest_value}")
