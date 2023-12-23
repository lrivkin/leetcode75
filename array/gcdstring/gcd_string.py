import string


class Solution:

    def divides_string(self, str1: str, base: str) -> bool:
        # print(f"divides_string({str1}, {base})")
        if len(str1)%len(base) != 0:
            return False
        
        repeats = int(len(str1)/len(base))
        # print(f"k= {repeats}, base*k = {base * repeats}, does divide? {base * repeats == str1}")

        return base * repeats == str1

    def gcdOfStrings(self, str1: str, str2: str) -> str:
        base = str1
        if len(str2) < len(str1):
            base = str2


        while base != "":
            # print(f"base: {base}")
            if self.divides_string(str1, base) and self.divides_string(str2, base):
                return base

            base = base[0:len(base)-1]

        return ""


if __name__ == '__main__':
    s = Solution()
    print(s.gcdOfStrings("ABCABC", "ABC"))
    print(s.gcdOfStrings("ABABAB", "AB"))
    print(s.gcdOfStrings("LEET", "CODE"))
