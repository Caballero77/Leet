namespace _209._Minimum_Size_Subarray_Sum;

public class Solution {
    public int MinSubArrayLen(int target, int[] nums)
    {
        var result = Int32.MaxValue;
        var i = 0;
        var j = 1;
        var sum = nums[i];
        while (true)
        {
            if (sum < target && j < nums.Length)
            {
                sum += nums[j];
                j++;
                continue;
            }

            if (sum >= target && i < nums.Length)
            {
                if (result > j - i)
                {
                    result = j - i;
                }

                sum -= nums[i];
                i++;
                continue;
            }
            
            break;
        }

        return result == Int32.MaxValue ? 0 : result;
    }
}