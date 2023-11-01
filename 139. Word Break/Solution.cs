namespace _139._Word_Break;

public class Solution {
    public bool WordBreak(string s, IList<string> words)
    {
        var dp = new bool[s.Length + 1];
        dp[0] = true;
        for (var i = 1; i < dp.Length; i++) dp[i] = false;

        for (var i = 1; i < dp.Length; i++)
        {
            for (var j = i - 1; j >= 0; j--)
            {
                if (dp[j] && words.Contains(s.Substring(j, i - j)))
                {
                    dp[i] = true;
                    break;
                } 
            }
        }

        return dp[s.Length];
    }
}