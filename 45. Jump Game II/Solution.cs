namespace _45._Jump_Game_II;

public class Solution
{
    public int Jump(int[] nums)
    {
        var dp = new int[nums.Length];
        for (var i = nums.Length - 2; i >= 0; i--)
        {
            if (nums[i] == 0)
            {
                dp[i] = 10001;
                continue;
            }
            var min = 10001;
            for (var j = i+1; j <= Math.Min(i + nums[i], nums.Length-1); j++)
            {
                if (min > dp[j])
                {
                    min = dp[j];
                }
            }

            dp[i] = min+1;
        }

        return dp[0];
    }
}