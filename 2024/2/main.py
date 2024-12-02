import pandas as pd
import numpy as np

def is_monotonic_increasing(x):
    dx = np.diff(x)
    return np.all(dx >= 0)

def is_monotonic_decreasing(x):
    dx = np.diff(x)
    return np.all(dx <= 0)

def is_level_legit(x):
    # Iterate over the row, check current and next value
    is_diff = True
    for i in range(len(x) - 1):
        if x[i] == x[i+1] or abs(x[i] - x[i+1]) > 3:
            is_diff = False
            break
    return is_diff

def part_one(df: pd.DataFrame):

    count = 0
    # Any two adjacent levels differ by at least one and at most three.
    for (index, series) in df.iterrows():
        row = series.to_numpy()
        row = row[~np.isnan(row)]

        if not (is_monotonic_increasing(row) or is_monotonic_decreasing(row)):
            continue

        # Iterate over the row, check current and next value
        if is_level_legit(row):
            count += 1

    print(count)


def part_two(df: pd.DataFrame):
    count = 0
    # Any two adjacent levels differ by at least one and at most three.
    for (index, series) in df.iterrows():
        row = series.to_numpy()
        row = row[~np.isnan(row)]

        # Check if any of the combinations are legit
        if (is_monotonic_increasing(row) or is_monotonic_decreasing(row)) and is_level_legit(row):
            count += 1
            continue 

        # Get leave one out combinations of the row
        combinations = np.array([np.delete(row, i) for i in range(len(row))])

        # Check if any of the combinations are legit
        for combination in combinations:
            if not (is_monotonic_increasing(combination) or is_monotonic_decreasing(combination)):
                continue
            if is_level_legit(combination):
                count += 1
                break
       
    print(count)


def main():
    # df = pd.read_csv("./2/example.csv", header=None, delimiter=' ')
    df = pd.read_csv("./2/input.csv", header=None, delimiter=' ')
    part_one(df)
    part_two(df)

if __name__ == "__main__":
    main()