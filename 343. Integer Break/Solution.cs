namespace _343._Integer_Break;

public class Solution
{
    private int[] dp;

    public Solution()
    {
        dp = new int[59];
        dp[0] = 0;
        dp[1] = 0;
        dp[2] = 1;
        for (var i = 3; i < 59; i++) dp[i] = -1;
    }
    
    public int IntegerBreak(int n)
    {
        if (dp[n] != -1)
        {
            return dp[n];
        }

        var max = 0;
        var l = 1;
        var r = n - 1;
        while (l <= r)
        {
            var left = IntegerBreak(l);
            var right = IntegerBreak(r);

            max = Math.Max(max, Math.Max(left, l) * Math.Max(right, r));
            l++;
            r--;
        }

        dp[n] = max;
        return max;
    }
}