using System.Collections.Generic;

public class Solution {
    public bool CanMeasureWater(int jug1Capacity, int jug2Capacity, int targetCapacity) {
        if (jug1Capacity + jug2Capacity < targetCapacity)
        {
            return false;
        }

        var queue = new Queue<int>();
        var hash = new HashSet<int>();

        var directions = new []{ jug1Capacity, jug2Capacity, -jug1Capacity, -jug2Capacity};

        queue.Enqueue(0);
        while (queue.Count > 0)
        {
            var current = queue.Dequeue();
            if (current == targetCapacity)
            {
                return true;
            }

            foreach (var direction in directions)
            {
                var next = current + direction;
                if (next > 0 && next <= jug1Capacity + jug2Capacity && !hash.Contains(next))
                {
                    hash.Add(next);
                    queue.Enqueue(next);
                }
            }
        }

        return false;
    }
}