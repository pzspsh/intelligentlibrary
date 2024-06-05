# -*- encoding: utf-8 -*-
'''
@File   : main2.py
@Time   : 2024-06-05 12:04:19
@Author : pch
'''
import numpy as np
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn import svm
from sklearn.metrics import accuracy_score
from sklearn.datasets import make_classification
from sklearn.preprocessing import StandardScaler

# 定义SVM模型
def create_svm(C, gamma):
    return svm.SVC(kernel="rbf", C=C, gamma=gamma)


# 目标函数（用于PSO评估），返回模型在测试集上的准确率
def fitness(params, X_train, y_train, X_test, y_test):
    C, gamma = params
    svm_model = create_svm(C, gamma)
    svm_model.fit(X_train, y_train)
    y_pred = svm_model.predict(X_test)
    accuracy = accuracy_score(y_test, y_pred)
    return -accuracy  # 因为PSO是最大化问题，而准确率越高越好，所以取负值


# 粒子类
class Particle:
    def __init__(self, dim, omega_start, omega_end, max_iter, bounds):
        self.position = np.random.uniform(bounds[:, 0], bounds[:, 1], dim)
        self.velocity = np.random.randn(dim)
        self.best_position = np.copy(self.position)
        self.best_score = float("-inf")
        self.omega_start = omega_start
        self.omega_end = omega_end
        self.max_iter = max_iter

    def update_velocity(self, t, personal_best, global_best):
        r1, r2 = np.random.rand(2, len(self.position))
        omega = self.omega_start - (self.omega_start - self.omega_end)*t / (self.max_iter - 1)
        cognitive_velocity = omega * self.velocity + r1 * (self.best_position - self.position)
        social_velocity = r2 * (global_best - self.position)
        self.velocity = cognitive_velocity + social_velocity

    def update_position(self, bounds):
        self.position += self.velocity
        self.position = np.clip(self.position, bounds[:, 0], bounds[:, 1])

    def evaluate(self, objective_function):
        score = objective_function(self.position)
        if score > self.best_score:
            self.best_score = score
            self.best_position = np.copy(self.position)
        return score


# PSO算法
def PSO(objective_function, dim, omega_start, omega_end, bounds, num_particles, num_iters):
    swarm = [Particle(dim, omega_start, omega_end, num_iters, bounds) for _ in range(num_particles)]
    global_best_score = float("-inf")
    global_best_position = None
    for t in range(num_iters):
        scores = [particle.evaluate(objective_function) for particle in swarm]
        best_particle_index = np.argmax(scores)
        best_particle = swarm[best_particle_index]
        if best_particle.best_score > global_best_score:
            global_best_score = best_particle.best_score
            global_best_position = best_particle.best_position
        for particle in swarm:
            particle.update_velocity(t=t,personal_best=particle.best_position, global_best=global_best_position)
            particle.update_position(bounds)
    return (global_best_position, -global_best_score)  # 返回最优参数和对应的准确率（取反）


if __name__ == "__main__":
    # 初始化PSO和SVM相关参数
    dim = 2  # SVM的C和gamma两个参数
    bounds = np.array([[0.1, 100], [0.001, 10]])  # C和gamma的搜索范围
    num_particles = 30
    num_iters = 100
    omega_start = 0.9
    omega_end = 0.4
    df = pd.read_excel("data.xlsx")
    Y = df.y
    X = df.drop("y", axis=1)
    scaler = StandardScaler()
    X1 = scaler.fit_transform(X)
    X_train, X_test, y_train, y_test = train_test_split(X1, Y, test_size=0.3, random_state=42)
    objective_function = lambda params: fitness(params, X_train, y_train, X_test, y_test)
    best_postion, best_score = PSO(objective_function, dim,omega_start, omega_end, bounds, num_particles, num_iters)
    best_c, best_gamma = best_postion
    best_accuracy = -best_score
    print(f"Optimized C: {best_c:.2f}, Gamma: {best_gamma:.4f}, Accuracy: {best_accuracy:.2f}")
