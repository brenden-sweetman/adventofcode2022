count = 0
pairs = []
with open("input", "r") as input:
    for line in input:
        pair = line.split(",")
        pair = [(int(s.split("-")[0]), int(s.split("-")[1])) for s in pair]
        pairs.append(pair)
print ("len: {}".format(len(pairs)))     
for pair in pairs:
    elf1 = pair[0]
    elf2 = pair[1]
    if (elf1[0] in range(elf2[0],elf2[1]+1) and elf1[1] in range(elf2[0],elf2[1]+1)) or (
        elf2[0] in range(elf1[0],elf1[1]+1) and elf2[1] in range(elf1[0],elf1[1]+1)
    ):
        count += 1
        print (pair)
print(count)

