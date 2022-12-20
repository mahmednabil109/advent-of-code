with open('./input1') as input1:
    guide = input1.read() \
        .replace('X', 'A') \
        .replace('Y', 'B') \
        .replace('Z', 'C') \
        .split('\n')
    # guide = list(guide)
    beats = {
        'A': 'C',
        'B': 'A',
        'C': 'B',
    }
    scores = {
        'A': 1,
        'B': 2,
        'C': 3,
    }
    total_score, total_score2 = 0, 0
    for game in guide:
        if not game: continue
        p1, p2 = game.split(' ')
        total_score += scores[p2]
        if p1 == p2:
            total_score += 3
        elif beats[p1] != p2:
            total_score += 6
        # else you lost the round
    for game in guide:
        if not game: continue
        p1, p2 = game.split(' ')
        if p2 == 'A':
            total_score2 += scores[beats[p1]]
        if p2 == 'B':
            total_score2 += 3 + scores[p1]
        if p2 == 'C':
            total_score2 += 6 + scores[[x for x in scores.keys() if beats[x] == p1][0]]
    print(total_score, total_score2)
