# 我们要找出只出现一次的那个数，其他数都出现了3次，如果我们把所有数的每一位（0-31位）加起来
# 那些出现了3次的数，在某一位上1的总个数，一定是3的倍数
# 唯一那个特殊的数，在某些位上会多出1，导致不是3的倍数
# 所以只要对每一位的1总数 %3， 就能知道这个唯一数在这一位 是不是1
def singleNumber(nums: list[int]) -> int:
    res = 0
    for i in range(32):
        # 统计所有数字在当前第i位上有多少个1
        bit_count = 0
        for num in nums:
            # num>>i, 将num右移i位，与 1做 &, 只保留最低位(即原来第i位)
            # 完整意思是 取出num的第i位，看它是不是1
            bit_count += (num >> i) & 1
        # 如果count%3 !=0 ,说明 第i位上，不是3的倍数 -> 肯定是那个只出现一次的数 “贡献”
        if bit_count % 3:
            if i == 31:  # 处理负数, 32位整数中，第31位是符号位，正数是0，负数是1
                res -= 1 << 31
            else:
                # 把第i位设成1更新结果
                res |= 1 << i
    return res


def singleNumberII(nums: list[int]) -> int:
    ones = twos = 0
    for num in nums:
        ones = (ones ^ num) & ~twos
        twos = (twos ^ num) & ~ones
    return ones
