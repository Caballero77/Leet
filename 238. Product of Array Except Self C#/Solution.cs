public class Solution {
    public int[] ProductExceptSelf(int[] nums) {
        ProductExceptSelfRec(nums, 0, 1);
        return nums;
    }

    private int ProductExceptSelfRec(int[] nums, int i, int n) {
        if (i == nums.Length - 1) {
            var buf = nums[i];
            nums[i] = n;
            return buf;
        }

        var next = ProductExceptSelfRec(nums, i + 1, n*nums[i]);
        var x = nums[i] * next;
        nums[i] = next * n;
        return x;
    }
}