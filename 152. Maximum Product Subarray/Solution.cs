namespace _152._Maximum_Product_Subarray;

public class Solution {

    public int MaxProduct(int[] nums)
    {
        var max = nums[0];
        var min = nums[0];
        var result = nums[0];
        for (var i = 1; i < nums.Length; i++)
        {
            var buf = max;
            max = Math.Max(nums[i], Math.Max(buf * nums[i], min * nums[i]));
            min = Math.Min(nums[i], Math.Min(buf * nums[i], min * nums[i]));
            result = Math.Max(result, max);
        }

        return result;
    }
}