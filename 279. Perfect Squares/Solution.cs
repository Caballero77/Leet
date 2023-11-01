namespace _279._Perfect_Squares;

public class Solution {
    public int NumSquares(int n)
    {
        var dp = new int[n+1];
        var sqrt = 1;
        for (var i = 1; i <= n; i++)
        {
            if (i == sqrt * sqrt)
            {
                dp[i] = 1;
                sqrt++;
                continue;
            }

            var min = Int32.MaxValue;
            for (var j = 1; j < i; j++)
            {
                min = Math.Min(min, dp[j] + dp[i-j]);
            }

            dp[i] = min;
        }

        return dp[n];
    }
}
