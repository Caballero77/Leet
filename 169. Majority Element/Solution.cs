namespace _169._Majority_Element;

public class Solution {
    public int MajorityElement(int[] nums)
    {
        var result = 0;
        var amount = 0;
        foreach (var number in nums)
        {
            if (amount == 0 || result == number)
            {
                result = number;
                amount++;
            } else if (result != number)
            {
                if (amount == 0)
                {
                    result = number;
                }
                else
                {
                    amount--;
                }
            }
        }
        return result;
    }
}