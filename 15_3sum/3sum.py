#!/usr/bin/env python3

"""
Just for fun/comparison to the Go version.
Python -> ~30s
Golang -> ~ 5s
"""

from typing import List

def three_sum(arr: List[int]) -> str:
    m = { e : i for i, e in enumerate(arr) }
    for i in range(len(arr)):
        for j in range(i, len(arr)):
            if (k:=m.get(-(arr[i]+arr[j]), None)) is not None:
                return f"{i+1} {j+1} {k+1}"
    return "-1"

if __name__ == "__main__":
    with open("data.txt", "r") as file:
        data = file.read().strip().split("\n")
    
    for arr in data[1:]:
        arr = list(map(int, arr.split(" ")))
        print(three_sum(arr))
