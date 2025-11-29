import galois
import numpy as np
import numpy.matlib as npml

def f(n):
    GF2 = galois.GF(2)

    O = np.zeros((n, n), dtype=int)
    O = GF2(O)

    B = np.zeros((n, n), dtype=int)
    for i in [-1, 0, 1]:
        B += npml.eye(n, k=i, dtype=int)
    B = GF2(B)

    # I = npml.identity(n, dtype=int)
    I = B
    I = GF2(I)

    a = []
    for i in range(n):
        row = []
        for j in range(n):
            if i == j:
                row.append(B)
            elif i == j + 1 or i == j - 1:
                row.append(I)
            else:
                row.append(O)
        a.append(row)
    
    A = np.block(a)
    A = GF2(A)

    E = A.row_reduce()
    print(n, n ** 2 - np.linalg.matrix_rank(E))

    cols = E.null_space()

    for col in cols:
        for i in range(n):
            for j in range(n):
                print(col[i*n + j], end=" ")
            print()
        print()

def main():
    f(4)
    f(5)
    f(6)

if __name__ == "__main__":
    main()
