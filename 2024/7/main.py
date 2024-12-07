import itertools


def solve_equations(equations, operators):
    total = 0

    for equation in equations:
        left_handside = equation[0]
        right_handside = equation[1:]

        possible_combinations = itertools.product(operators, repeat=len(right_handside) - 1)

        for combination in possible_combinations:
            result = right_handside[0]

            for i, operator in enumerate(combination):
                index = i + 1

                if operator == '+':
                    result += right_handside[index]
                elif operator == '*':
                    result *= right_handside[index]
                elif operator == '||':
                    result = int(str(result) + str(right_handside[index]))

            if result == left_handside:
                total += result
                break

    return total


def main():
    with open('./7/input.txt', 'r') as input_file:
        lines = input_file.readlines()

    equations = [
        [int(parts[0])] + list(map(int, parts[1].strip().split(' ')))
        for line in lines
        for parts in [line.split(':')]
    ]

    part_one = solve_equations(equations, operators=['+', '*'])
    part_two = solve_equations(equations, operators=['+', '*', '||'])

    print(part_one)
    print(part_two)


if __name__ == '__main__':
    main()