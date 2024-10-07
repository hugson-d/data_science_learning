from typing import List

def binary_search(arr: List[int], target: int) -> int:
    low: int = 0
    high: int = len(arr) - 1

    while low <= high:
        mid: int = (low + high) // 2  # Calculate mid based on the current low and high

        if arr[mid] == target:
            return mid  # Return the index, not the target value

        elif arr[mid] < target:
            low = mid + 1  # Narrow the search to the right half

        else:
            high = mid - 1  # Narrow the search to the left half

    return -1  # Return -1 if the target is not found

# Example usage
myArray: List[int] = [i for i in range(0, 20, 2)]  # List of even numbers
result = binary_search(myArray, 4)

if result != -1:
    print(f"Element found at index {result}.")
else:
    print("Element not found.")
