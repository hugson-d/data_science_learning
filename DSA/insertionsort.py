array = [3,56,777,44444,2133,23,131221313,131313,1,222,2]


def insertion_sort(array):
    for i in range(1, len(array)):
        key = array[i]
        # Start comparing with the previous item
        for j in range(i - 1, -1, -1):
            print(array[j],key)
            if array[j] > key:
                array[j + 1] = array[j]  # Shift element to the right
                array[j] = key           # Insert key at the correct position
            else:
                break  # Exit the loop if the current element is less than or equal to key
    return array



sorted_array = insertion_sort(array)
print(sorted_array) 

