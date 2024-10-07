import random
import time
from typing import List

def bubble_sort(unsortedList: List[int]):
    length = len(unsortedList) - 1
    while length > 0:
        for j in range(0, length):
            if unsortedList[j] > unsortedList[j+1]:
                unsortedList[j], unsortedList[j+1] = unsortedList[j+1], unsortedList[j]
        length = length -1
    return unsortedList




l = []
for i in range(9999):
    l.append(random.randint(0,99))

start = time.time()
l = bubble_sort(l)
print(l)
end = time.time()
print(end-start)
