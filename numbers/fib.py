def fib(n):
    if n <= 2:
        return [0, 1][: n + 1]
    fibs = [None] * n
    fibs[:2] = [0, 1]
    for i in range(2, n):
        fibs[i] = fibs[i - 2] + fibs[i - 1]
    return fibs


print(fib(0))
print(fib(1))
print(fib(5))
print(fib(100))
