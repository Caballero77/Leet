namespace _57._Insert_Interval;

public class Solution {
    public int[][] Insert(int[][] intervals, int[] newInterval)
    {
        var result = new List<int[]>();
        var inserted = false;
        for (var i = 0; i < intervals.Length; i++)
        {
            if (newInterval[0] > intervals[i][1])
            {
                result.Add(intervals[i]);
                continue;
            }
            if (newInterval[1] < intervals[i][0])
            {
                if (!inserted)
                {
                    result.Add(newInterval);
                    inserted = true;
                }
                result.Add(intervals[i]);
                continue;
            }

            newInterval[0] = Math.Min(newInterval[0], intervals[i][0]);
            newInterval[1] = Math.Max(newInterval[1], intervals[i][1]);
        }

        if (!inserted)
        {
            result.Add(newInterval);
        }

        return result.ToArray();

    }
}