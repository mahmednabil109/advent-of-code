with open('../in.txt') as file:
    result = 0
    for line in file.readlines():
        if not line: continue;
        digits = list(map(int, filter(lambda x:x.isdigit(), line)))
        result += digits[0] * 10 + digits[-1]
    print(result)