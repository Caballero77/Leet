namespace _509._Fibonacci_Number;

public class Solution
{
    private int[] mem;

    public Solution()
    {
        mem = new int[31];
        mem[0] = 0;
        mem[1] = 1;
        for (var i = 2; i < 31; i++) mem[i] = -1;
    }
    
    public int Fib(int n)
    {
        if (mem[n] != -1) return mem[n];
        var res = Fib(n - 1) + Fib(n - 2);
        mem[n] = res;
        return res;
    }
}