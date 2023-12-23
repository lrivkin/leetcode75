import numpy as np

class Solution:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        if str1 + str2 != str2 + str1:
            return ""
        gcd = np.gcd(len(str1), len(str2))

        return str1[0:gcd]

if __name__ == '__main__':
    s = Solution()
    print(s.gcdOfStrings("ABCABC", "ABC"))
    print(s.gcdOfStrings("ABABAB", "AB"))
    print(s.gcdOfStrings("LEET", "CODE"))
