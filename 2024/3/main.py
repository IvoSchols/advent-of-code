
import re


def part_one():
    total = 0
    with open('./3/input.txt') as f:
        lines = f.readlines()
        for line in lines:
            # A line can hold many matches
            matches = re.findall(r'mul\((\d+),(\d+)\)', line)

            # Fold the matches into the sum
            total += sum([int(match[0]) * int(match[1]) for match in matches])

    print(total)

# 102210215 is too high (i.e., we are either missing a dont() or enabling a do())
# 88811886
def part_two():
    total = 0
    prune = False
    with open('./3/input.txt') as f:
        lines = f.readlines()
        for line in lines:
            # A line can hold many matches
            matches = re.findall(r'(mul\((\d+),(\d+)\))|(don\'t\(\))|(do\(\))', line)

            # Prune the matches which are in between don't() and do()
            for match in matches:
                if match[3] == 'don\'t()':
                    prune = True
                elif match[4] == 'do()':
                    prune = False
                # Check if match1 can be a number
                elif not prune:
                    number_one = int(match[1])
                    number_two = int(match[2])
                    total += number_one * number_two

    print(total)



def main():

    part_one()
    part_two()

if __name__ == "__main__":
    main()