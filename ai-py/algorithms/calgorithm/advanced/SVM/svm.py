# -*- encoding: utf-8 -*-
"""
@File   : svm.py
@Time   : 2024-05-23 12:31:55
@Author : pan
"""
import numpy as np
import pandas as pd
from sklearn.datasets import load_iris
from sklearn.model_selection import train_test_split
import matplotlib.pyplot as plt


# 加载数据
def create_data():
    iris = load_iris()
    # 选取前两类作为数据
    Data = np.array(iris["data"])[:100]
    Label = np.array(iris["target"])[:100]
    Label = Label * 2 - 1
    print("dara shape:", Data.shape)
    print("label shape:", Label.shape)
    return Data, Label


class SVM:
    def __init__(self, max_iter=100, kernel="linear"):
        self.max_iter = max_iter
        self._kernel = kernel
        # 参数初始化

    def init_args(self, features, labels):
        self.m, self.n = features.shape
        self.X = features
        self.Y = labels
        self.b = 0.0
        self.alpha = np.ones(self.m)
        self.computer_product_matrix()
        # 为了加快训练速度创建一个内积矩阵
        # 松弛变量
        self.C = 1.0
        # 将Ei保存在一个列表里
        self.create_E()
        # KKT条件判断

    def judge_KKT(self, i):
        y_g = self.function_g(i) * self.Y[i]
        if self.alpha[i] == 0:
            return y_g >= 1
        elif 0 < self.alpha[i] < self.C:
            return y_g == 1
        else:
            return y_g <= 1
        # 计算内积矩阵#如果数据量较大，可以使用系数矩阵

    def computer_product_matrix(self):
        self.product_matrix = np.zeros((self.m, self.m)).astype(np.float)
        for i in range(self.m):
            for j in range(self.m):
                if self.product_matrix[i][j] == 0.0:
                    self.product_matrix[i][j] = self.product_matrix[j][i] = self.kernel(
                        self.X[i], self.X[j]
                    )
                    # 核函数

    def kernel(self, x1, x2):
        if self._kernel == "linear":
            return np.dot(x1, x2)
        elif self._kernel == "poly":
            return (np.dot(x1, x2) + 1) ** 2
        return 0

    # 将Ei保存在一个列表里
    def create_E(self):
        self.E = (np.dot((self.alpha * self.Y), self.product_matrix) + self.b) - self.Y

    # 预测函数g(x)
    def function_g(self, i):
        return self.b + np.dot((self.alpha * self.Y), self.product_matrix[i])

    # 选择变量
    def select_alpha(self):
        # 外层循环首先遍历所有满足0<a<C的样本点，检验是否满足KKT
        index_list = [i for i in range(self.m) if 0 < self.alpha[i] < self.C]
        # 否则遍历整个训练集
        non_satisfy_list = [i for i in range(self.m) if i not in index_list]
        index_list.extend(non_satisfy_list)
        for i in index_list:
            if self.judge_KKT(i):
                continue
            E1 = self.E[i]
            # 如果E2是+，选择最小的；如果E2是负的，选择最大的
            if E1 >= 0:
                j = np.argmin(self.E)
            else:
                j = np.argmax(self.E)
            return i, j

        # 剪切
        def clip_alpha(self, _alpha, L, H):
            if _alpha > H:
                return H
            elif _alpha < L:
                return L
            else:
                return _alpha

        # 训练函数，使用SMO算法
        def Train(self, features, labels):
            self.init_args(features, labels)
            # SMO算法训练
            for t in range(self.max_iter):
                i1, i2 = self.select_alpha()
                # 边界
                if self.Y[i1] == self.Y[i2]:
                    L = max(0, self.alpha[i1] + self.alpha[i2] - self.C)
                    H = min(self.C, self.alpha[i1] + self.alpha[i2])
                else:
                    L = max(0, self.alpha[i2] - self.alpha[i1])
                    H = min(self.C, self.C + self.alpha[i2] - self.alpha[i1])

                E1 = self.E[i1]
                E2 = self.E[i2]
                # eta=K11+K22-2K12
                eta = (
                    self.kernel(self.X[i1], self.X[i1])
                    + self.kernel(self.X[i2], self.X[i2])
                    - 2 * self.kernel(self.X[i1], self.X[i2])
                )
                if eta <= 0:
                    # print('eta <= 0')
                    continue

                alpha2_new_unc = self.alpha[i2] + self.Y[i2] * (E1 - E2) / eta
                # 此处有修改，根据书上应该是E1 - E2，书上130-131页
                alpha2_new = self.clip_alpha(alpha2_new_unc, L, H)
                alpha1_new = self.alpha[i1] + self.Y[i1] * self.Y[i2] * (
                    self.alpha[i2] - alpha2_new
                )
                b1_new = (
                    -E1
                    - self.Y[i1]
                    * self.kernel(self.X[i1], self.X[i1])
                    * (alpha1_new - self.alpha[i1])
                    - self.Y[i2]
                    * self.kernel(self.X[i2], self.X[i1])
                    * (alpha2_new - self.alpha[i2])
                    + self.b
                )
                b2_new = (
                    -E2
                    - self.Y[i1]
                    * self.kernel(self.X[i1], self.X[i2])
                    * (alpha1_new - self.alpha[i1])
                    - self.Y[i2]
                    * self.kernel(self.X[i2], self.X[i2])
                    * (alpha2_new - self.alpha[i2])
                    + self.b
                )
                if 0 < alpha1_new < self.C:
                    b_new = b1_new
                elif 0 < alpha2_new < self.C:
                    b_new = b2_new
                else:  # 选择中点
                    b_new = (b1_new + b2_new) / 2

                # 更新参数
                self.alpha[i1] = alpha1_new
                self.alpha[i2] = alpha2_new
                self.b = b_new
                self.create_E()
                # 这里与书上不同，，我选择更新全部E

        def predict(self, data):
            r = self.b
            for i in range(self.m):
                r += self.alpha[i] * self.Y[i] * self.kernel(data, self.X[i])

            return 1 if r > 0 else -1

        def score(self, X_test, y_test):
            right_count = 0
            for i in range(len(X_test)):
                result = self.predict(X_test[i])
                if result == y_test[i]:
                    right_count += 1
            return right_count / len(X_test)


if __name__ == "__main__":
    svm = SVM(max_iter=200)
    X, y = create_data()
    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.333, random_state=23323
    )
    svm.Train(X_train, y_train)
    print(svm.score(X_test, y_test))
