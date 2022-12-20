with open('input1') as input1:
    rucksacks = input1.read().split('\n')
    rucksacks = list(filter(lambda x: x, rucksacks))
    p_sum = 0
    for rucksack in rucksacks:
        if not rucksack: continue
        l = len(rucksack)
        c1, c2 = set(rucksack[:l//2]), set(rucksack[l//2:])
        inter = list(c1.intersection(c2))
        assert len(inter) == 1
        inter0 = ord(inter[0])
        if ord('a') <=  inter0 <= ord('z'):
            p_sum += 1  + inter0 - ord('a')
        else:
            p_sum += 27 + inter0 - ord('A')
    p_sum2 = 0
    chunkify = lambda b, n: (b[i:i+n] for i in range(0,len(b),n))
    for rucksack in chunkify(rucksacks, 3):
        sets = list(map(set, rucksack))
        inter = list(sets[0].intersection(*sets))
        assert len(inter) == 1
        inter0 = ord(inter[0])
        if ord('a') <= inter0 <= ord('z'):
            p_sum2 += 1 + inter0 - ord('a')
        else:
            p_sum2 += 27 + inter0 - ord('A')
    print(p_sum, p_sum2)
