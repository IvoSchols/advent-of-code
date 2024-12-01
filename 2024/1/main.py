import pandas as pd
import numpy as np

def part_one(df):
    column_one = df[0].values
    column_two = df[1].values

    column_one = np.sort(column_one)
    column_two = np.sort(column_two)

    diff = np.abs(column_one - column_two)

    sum_diff = np.sum(diff)

    print(sum_diff)

def part_two(df: pd.DataFrame):
    column_one = df[0].values

    column_two_occurence = df[1].value_counts()
    column_two_occurence = column_two_occurence.to_dict()

    summed_similarity_score = 0
    # Fill in missing values with 0 & sum
    for c1 in column_one:
        column_two_occurence[c1] = 0 if c1 not in column_two_occurence else column_two_occurence[c1]
        summed_similarity_score += c1 * column_two_occurence[c1]


    print(summed_similarity_score)

def main():
    df = pd.read_csv("./1/input.csv", header=None)
    part_one(df)
    part_two(df)

if __name__ == "__main__":
    main()