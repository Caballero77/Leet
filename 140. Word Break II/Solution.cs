namespace _140._Word_Break_II;

public class Solution {
    private IList<string> PrepareResult(IList<string>[] results, int i)
    {
        if (i == 0)
        {
            return results[i];
        }

        var result = new List<string>();
        foreach (var word in results[i])
        {
            foreach (var prev in PrepareResult(results, i - word.Length))
            {
                result.Add(prev.Length > 0 ? prev + " " + word : word);
            }
            
        }

        return result;
    }
    
    public IList<string> WordBreak(string s, IList<string> words)
    {
        var dp = new List<string>[s.Length + 1];
        dp[0] = new List<string>{ "" };
        for (var i = 1; i < dp.Length; i++) dp[i] = new List<string>();

        for (var i = 1; i < dp.Length; i++)
        {
            for (var j = i - 1; j >= 0; j--)
            {
                var curr = s.Substring(j, i - j);
                if (dp[j].Count > 0 && words.Contains(curr))
                {
                    dp[i].Add(curr);
                } 
            }
        }
        
        return PrepareResult(dp, s.Length);
    }
}