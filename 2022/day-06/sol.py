with open('input1') as input1:
    stream = input1.read().strip()
    num_of_distinct = 14 # for part one use 4
    for i in range(0, len(stream)):
        _set = set(stream[i:i+num_of_distinct])
        if len(_set) == num_of_distinct:
            print(stream[i:i+num_of_distinct])
            print(i + num_of_distinct)
            break
