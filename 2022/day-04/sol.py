import re

with open('input1') as input1:
    ranges = input1.read().strip().split('\n')
    fully_containted, overlaped = 0, 0
    for range_pair in ranges:
        x1, y1, x2, y2 = map(int, re.split(r'-|,', range_pair))
        if (x1 >= x2 and y1 <= y2) or (x1 <= x2 and y1 >= y2):
           fully_containted += 1
        if (x1 >= x2 and x1 <= y2) or (y1 >= x2 and y1 <= y2) or (x2 >= x1 and x2 <= y1):
            overlaped += 1
    print(fully_containted, overlaped)

