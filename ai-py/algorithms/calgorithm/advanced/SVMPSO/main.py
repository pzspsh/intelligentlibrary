import numpy as np
from sklearn import datasets
from sklearn.model_selection import train_test_split
from sklearn import svm
from sklearn.metrics import accuracy_score

# 加载数据集（例如鸢尾花数据集）
iris = datasets.load_iris()
X = iris.data
y = iris.target

# 划分训练集和测试集
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.3, random_state=42)


# 定义SVM模型
def create_svm(C, gamma):
    return svm.SVC(kernel="rbf", C=C, gamma=gamma)


# 目标函数（用于PSO评估），返回模型在测试集上的准确率
def objective_function(params):
    C, gamma = params
    svm_model = create_svm(C, gamma)
    svm_model.fit(X_train, y_train)
    y_pred = svm_model.predict(X_test)
    accuracy = accuracy_score(y_test, y_pred)
    return -accuracy  # 因为PSO是最大化问题，而准确率越高越好，所以取负值


# 粒子类
class Particle:
    def __init__(self, dim, bounds):
        self.position = np.random.uniform(bounds[:, 0], bounds[:, 1], dim)
        self.velocity = np.zeros(dim)
        self.best_position = np.copy(self.position)
        self.best_score = float("-inf")

    def update_velocity(self, omega, personal_best, global_best):
        r1, r2 = np.random.rand(2, len(self.position))
        cognitive_velocity = omega * self.velocity + r1 * (
            self.best_position - self.position
        )
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
def pso(objective_function, dim, bounds, num_particles, num_iterations, omega=0.5):
    swarm = [Particle(dim, bounds) for _ in range(num_particles)]
    global_best_score = float("-inf")
    global_best_position = None

    for _ in range(num_iterations):
        scores = [particle.evaluate(objective_function) for particle in swarm]
        best_particle_index = np.argmax(scores)
        best_particle = swarm[best_particle_index]
        if best_particle.best_score > global_best_score:
            global_best_score = best_particle.best_score
            global_best_position = best_particle.best_position

        for particle in swarm:
            particle.update_velocity(
                omega, particle.best_position, global_best_position
            )
            particle.update_position(bounds)

    return (global_best_position,-global_best_score)  # 返回最优参数和对应的准确率（取反）


if __name__ == "__main__":
    # 设置PSO和SVM的参数范围
    dim = 2  # SVM的C和gamma两个参数
    bounds = np.array([[0.1, 100], [0.001, 10]])  # C和gamma的搜索范围
    num_particles = 30
    num_iterations = 100
    omega = 0.5

    # 运行PSO算法优化SVM参数
    best_position, best_score = pso(objective_function, dim, bounds, num_particles, num_iterations, omega)
    best_C, best_gamma = best_position
    best_accuracy = -best_score  # 取反得到真实的准确率

    print(f"Optimized C: {best_C:.2f}, Gamma: {best_gamma:.4f}, Accuracy: {best_accuracy:.2f}")    
