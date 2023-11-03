public class Solution {
    public int[] AsteroidCollision(int[] asteroids) {
        var stack = new Stack<int>();
        for(var i = 0; i < asteroids.Length;) {
            var asteroid = asteroids[i];
            if (stack.Count == 0) {
                stack.Push(asteroid);
                i++;
                continue;
            }

            var peek = stack.Peek();
            if (asteroid < 0 && peek > 0)
            {
                if (Math.Abs(asteroid) == Math.Abs(peek)) {
                    stack.Pop();
                    i++;
                } else if (Math.Abs(asteroid) > Math.Abs(peek)) {
                    stack.Pop();
                } else {
                    i++;
                }
            } else {
                stack.Push(asteroid);
                i++;
            }
        }
        return stack.Reverse().ToArray();    
    }
}