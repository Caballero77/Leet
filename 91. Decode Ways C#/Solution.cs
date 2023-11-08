public class Solution {

    private static readonly HashSet<string> Letters = new();

    private static readonly Dictionary<string, int> Memory = new();

    public Solution()
    {
        for (var i = 1; i <= 26; i++) {
            Letters.Add(i.ToString());
        }
    }

    public int NumDecodings(string s) {
        if (Memory.ContainsKey(s)) {
            return Memory[s];
        }
        if (s.Length == 0) {
            return 1;
        }
        var result = 0;

        if (s.Length >= 1)
        {
            var first = s.Substring(0, 1);
            if (Letters.Contains(first)) {
                result += NumDecodings(s.Substring(1));
            }
        }

        if (s.Length >= 2)
        {
            var second = s.Substring(0, 2);
            if (Letters.Contains(second)) {
                result += NumDecodings(s.Substring(2));
            }
        }
        
        Memory[s] = result;

        return result;
    }
}