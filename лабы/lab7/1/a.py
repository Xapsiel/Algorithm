import numpy as np
from scipy.optimize import linprog

# Целевая функция
c = np.array([-3, -2])

# Ограничения-неравенства
A_ub = np.array([[3, 7], [1, 1]])
b_ub = np.array([21, 4])

# Границы переменных
bounds = [(0, 4), (0, 3)]

# Решение релаксированной задачи
res = linprog(-c, A_ub=A_ub, b_ub=b_ub, bounds=bounds)
x_relax = res.x

# Округление до целых значений
x1 = int(x_relax[0])
x2 = int(x_relax[1])

# Проверка целочисленности решения
if x1 == x_relax[0] and x2 == x_relax[1]:
    print(f"Оптимальное решение: x1 = {x1}, x2 = {x2}")
    print(f"Максимальное значение функции: {-res.fun}")
else:
    # Ветвление по переменной с дробной частью
    if x_relax[0] - x1 > x_relax[1] - x2:
        x1_new = x1 + 1
        x1_new_bounds = (x1_new, x1_new)
        x2_new_bounds = bounds[1]
    else:
        x2_new = x2 + 1
        x1_new_bounds = bounds[0]
        x2_new_bounds = (x2_new, x2_new)

    # Рекурсивный вызов для поддерева с большим значением целевой функции
    new_bounds = [x1_new_bounds, x2_new_bounds]
    res_new = linprog(-c, A_ub=A_ub, b_ub=b_ub, bounds=new_bounds)
    if -res_new.fun > -res.fun:
        x1, x2 = int(res_new.x[0]), int(res_new.x[1])
        print(f"Оптимальное решение: x1 = {x1}, x2 = {x2}")
        print(f"Максимальное значение функции: {-res_new.fun}")
    else:
        print(f"Оптимальное решение: x1 = {x1}, x2 = {x2}")
        print(f"Максимальное значение функции: {-res.fun}")
