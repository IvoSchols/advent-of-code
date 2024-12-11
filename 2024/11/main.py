from collections import defaultdict

# If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
def rule_1(stone):
    return 1
# If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
def rule_2(stone):
    stone = str(stone)
    left = stone[:len(stone) // 2]
    right = stone[len(stone) // 2:]
    return (int(left), int(right))
# If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
def rule_3(stone):
    return stone * 2024

def apply_rules(stones):
    new_stones = []
    for stone in stones:
        if stone == 0:
            new_stones.append(rule_1(stone))
        elif len(str(stone)) % 2 == 0:
            new_stones.extend(rule_2(stone))
        else:
            new_stones.append(rule_3(stone))
    return new_stones

def part_one(stones, n_blinks=25):
    for _ in range(n_blinks):
        stones = apply_rules(stones)
    print(len(stones))



def part_two(stones, n_blinks=75):
    # Create a dictionary to store the count of occurrences of each stone
    stone_count = defaultdict(int)
    for stone in stones:
        stone_count[stone] += 1

    for _ in range(n_blinks):
        new_stones = defaultdict(int)
        # Update the stone count dictionary
        for stone, count in stone_count.items():
            if stone == 0:
                new_stones[rule_1(stone)] += count
            elif len(str(stone)) % 2 == 0:
                left, right = rule_2(stone)
                new_stones[left] += count
                new_stones[right] += count
            else:
                new_stones[rule_3(stone)] += count
        stone_count = new_stones

    total_stones = 0
    for count in stone_count.values():
        total_stones += count

    print(total_stones)

def main():
    with open("./11/input.txt") as file:
        # Read stones from file, e.g.: 0 1 5 10
        stones = [int(stone) for stone in file.readline().split()]
        
    part_one(stones)
    part_two(stones)

if __name__ == '__main__':
    main()