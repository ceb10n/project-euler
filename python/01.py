# -*- coding: utf-8 -*-
"""Project Euler #1: Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

Project Euler: https://projecteuler.net/problem=1
Hackerrank: https://www.hackerrank.com/contests/projecteuler/challenges/euler001/problem
"""
import sys


def sum_multiples_of_3_or_5(number: int):
    return \
        3 * (((number // 3) * ((number // 3) + 1)) // 2) + \
        5 * (((number // 5) * ((number // 5) + 1)) // 2) - \
        15 * (((number // 15) * ((number // 15) + 1)) // 2)


if __name__ == "__main__":
    t = int(input().strip())

    numbers = []

    for _ in range(t):
        n = int(input().strip())
        numbers.append(sum_multiples_of_3_or_5(n - 1))

    for number in numbers:
        print(number, end="\n")
