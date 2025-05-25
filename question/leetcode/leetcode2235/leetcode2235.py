class Solution:
    # •	a ^ b：表示不带进位的加法结果（按位异或）
    # •	a & b：表示产生进位的位（按位与），但没有移动一位。
    def sum(self, num1: int, num2: int) -> int:
        return (num1 ^ num2) + ((num1 & num2) << 1)

    def sum2(self, num1: int, num2: int) -> int:
        return num1 + num2
