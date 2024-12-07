import itertools


def part_one(equations):
    
    operators = ['+','*']
    total = 0

    # Test combinations of operators and numbers to make the equation true
    for equation in equations:
        left_handside = equation[0]
        right_handside = equation[1:]

        # Get cartesian product of operators and numbers (remember an operator is always between two numbers)
        possible_combinations = itertools.product(operators, repeat=len(right_handside)-1)

        for combination in possible_combinations:
            sum = right_handside[0]
            # Apply the combination to the right handside
            for i, operator in enumerate(combination):
                index = i + 1

                if operator == '+':
                    sum += right_handside[index]
                elif operator == '*':
                    sum *= right_handside[index]
            # Check if the sum is equal to the left handside
            if sum == left_handside:
                total += sum
                break



    print(total)
        

            
def part_two(equations):

    operators = ['+','*', '||']
    total = 0

    # Test combinations of operators and numbers to make the equation true
    for equation in equations:

        left_handside = equation[0]
        right_handside = equation[1:]

        # Get cartesian product of operators and numbers (remember an operator is always between two numbers)
        possible_combinations = itertools.product(operators, repeat=len(right_handside)-1)

        for combination in possible_combinations:
            sum = right_handside[0]
            # Apply the combination to the right handside
            for i, operator in enumerate(combination):
                index = i + 1

                if operator == '+':
                    sum += right_handside[index]
                elif operator == '*':
                    sum *= right_handside[index]
                elif operator == '||':
                    sum = int(str(sum) + str(right_handside[index]))
            # Check if the sum is equal to the left handside
            if sum == left_handside:
                total += sum
                break

    print(total)


def main():
    input_file = open('./7/input.txt', 'r')
    
    # Read the input file
    lines = input_file.readlines()
    input_file.close()

    # Build the equations to solve
    equations = []

    # Split on :
    for line in lines:
        equation = []
        parts = line.split(':')
        left_handside = int(parts[0])
        equation.append(left_handside)
        # Strip and cast to int the right handside
        right_handside = parts[1].strip().split(' ')
        right_handside = [int(x) for x in right_handside]
        equation.extend(right_handside)
        equations.append(equation)
        



    part_one(equations)
    part_two(equations)

if __name__ == '__main__':
    main()