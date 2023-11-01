namespace _80._Remove_Duplicates_from_Sorted_Array_II;

public class Solution {
    public int RemoveDuplicates(int[] nums) {
        if (nums.Length == 1)
        {
            return 1;
        }

        var length = 1;
        var amount = 1;
        for (var i = 1; i < nums.Length; i++)
        {
            if (nums[i] == nums[i - 1])
            {
                if (amount == 1)
                {
                    nums[length] = nums[i];
                    amount++;
                    length++;
                } else if (amount == 2) continue;
            }
            else
            {
                nums[length] = nums[i];
                amount = 1;
                length++;
            }
        }

        return length;
    }
}