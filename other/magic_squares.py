import math


def encryption(s):
    L = len(s)
    rL = math.sqrt(L)
    r = math.floor(rL)
    c = math.ceil(rL)

    if c * r < L:
        r += 1

    rows = []
    for i in range(r):
        rows.append(s[i * c : (i + 1) * c])

    enc = ["".join([rows[i][j] for i in range(r - 1)]) for j in range(c)]
    for i, item in enumerate(rows[-1]):
        enc[i] += item

    return " ".join(enc)


if __name__ == "__main__":
    s = "haveaniceday"
    print(encryption(s))

    s = "feedthedog"
    print(encryption(s))

    s = "chillout"
    print(encryption(s))
