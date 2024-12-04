# 3329 is too high 


# We will search for horizontal, vertical, and diagonal occurences of the words to search
# Check if the word can fit in the row,
def search_horizontal(input, x, y, word):
    input = input[y]
    count = 0
    word_right = x + len(word)

    if word_right <= len(input):
        found_word = "".join(input[x:word_right])
        if found_word == word:
            count += 1

    return count


def search_vertical(input, x, y, word):
    count = 0
    word_down = y + len(word)

    if word_down <= len(input):
        found_word = "".join([input[y + i][x] for i in range(len(word))])
        if found_word == word:
            count += 1

    return count

def search_diagonal(input, x, y, word):
    count = 0

    word_down_right = (x + len(word), y + len(word))
    if word_down_right[0] <= len(input[0]) and word_down_right[1] <= len(input):
        word_found = "".join([input[y + i][x + i] for i in range(len(word))])
        if word_found == word:
            count += 1

    word_down_left = (x - len(word) + 1, y + len(word))
    if word_down_left[0] >= 0 and word_down_left[1] <= len(input):
        word_found = "".join([input[y + i][x - i] for i in range(len(word))])
        if word_found == word:
            count += 1
    return count 

def search_cell(input, x, y, word):
    count = 0
    count += search_horizontal(input, x, y, word)
    count += search_vertical(input, x, y, word)
    count += search_diagonal(input, x, y, word)
    return count

def part_one(input):
    count = 0

    word_to_search = "XMAS"
    words_to_search = [word_to_search, word_to_search[::-1]]

    max_width = len(input[0])
    max_height = len(input)


    for y in range(max_height):
        for x in range(max_width):
            for word in words_to_search:
                count += search_cell(input, x, y, word)
                
    print(count)

# 1950 is too high
def part_two(input):
    # We solve this with a filter, we look for an A, and check if the diagonal corners are S or M
    # If there are two of each, we add 1 to the count
    # We do this for all A's in the grid
    count = 0
    
    max_width = len(input[0])
    max_height = len(input)

    diagonal_patterns = [
        ('M', 'A', 'S'),
        ('S', 'A', 'M')
    ]

    for y in range(1, max_height-1):
        for x in range(1, max_width-1):
            mas_count = 0

            if input[y][x] == "A":
                # Look for the diagonal corners
                for pattern in diagonal_patterns:
                    # Check if the pattern fits from left to right
                    if input[y-1][x-1] == pattern[0] and input[y+1][x+1] == pattern[2]:
                        mas_count += 1
                    # Check if the pattern fits from right to left
                    if input[y-1][x+1] == pattern[0] and input[y+1][x-1] == pattern[2]:
                        mas_count += 1
                if mas_count == 2:
                    count += 1

    print(count)

def main():
    # Read input.txt into input
    with open("./4/input.txt", "r") as file:
        input = file
        input = input.read()

    input = input.split("\n")
    # Convert input to 2D array
    input = [list(row) for row in input]


    test_input = """MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"""

    test_input = test_input.split("\n")
    test_input = [list(row) for row in test_input]


    part_one(input)
    part_two(input)

if __name__ == "__main__":
    main()