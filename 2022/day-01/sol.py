with open('./input1') as input1 :
    cals = input1.read().split('\n\n')
    cals_per_elf = list(map(lambda x: sum(map(int, x.split('\n'))), cals))
    ans = max(cals_per_elf)
    ans2 = sum(list(sorted(cals_per_elf))[-3:])
    print(ans, ans2)
