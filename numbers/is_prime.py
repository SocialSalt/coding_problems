import time
import numpy as np


def prime_sieve(n):
    is_num_prime = np.ones(n + 1)
    is_num_prime[[0, 1]] = 0
    # is_num_prime = [True] * (n + 1)
    # is_num_prime[0] = False
    # is_num_prime[1] = False

    for i, p in enumerate(is_num_prime):
        if p:
            yield i
            # computing the number of elements isn't as fast as taking len of the list
            # the number if elements is max(ceil((n + 1 - i * i) / i), 0)
            # we only need to start from the square here
            # is_num_prime[i * i :: i] = [False] * len(is_num_prime[i * i :: i])
            is_num_prime[i * i :: i] = False


def is_prime_sieve(n):
    if n % 1 != 0 or n == 1:
        return False
    if n == 2:
        return True
    for i in prime_sieve(np.sqrt(n).astype(int) + 1):
        if n % i == 0:
            return False
    return True


def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    for i in range(2, int(np.sqrt(n)) + 1):
        if n % i == 0:
            return False
    return True


def is_prime_np(n):
    if n < 2:
        return False
    if n == 2:
        return True
    nums = np.arange(2, np.sqrt(n) + 1)
    nums = n % nums
    return not any(nums == 0)


def main():
    np.random.seed(1)
    nums = np.random.randint(2, 1_000_000, 10_000)
    s = time.time()
    for num in nums:
        is_prime(num)
    print(f"took: {time.time() - s}")


main()
