namespace _392._Is_Subsequence;

public class Solution {
    public bool IsSubsequence(string s, string t)
    {
        if (s.Length == 0 || (s.Length == 0 && t.Length == 0)) return true;
        var i = 0;
        var j = 0;
        while (j < t.Length)
        {
            if (s[i] == t[j])
            {
                i++;
                if (i == s.Length)
                {
                    return true;
                }
            }

            j++;
        }

        return false;
    }
}