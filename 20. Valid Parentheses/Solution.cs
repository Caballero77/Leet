using System.Collections.Generic;

public class Solution {
    static Dictionary<char, char> brackets = new Dictionary<char, char>
        {
            { ')', '(' },
            { '}', '{' },
            { ']', '[' },
        };
    public bool IsValid(string s) {
        var stack = new Stack<char>();
        foreach (var c in s) {
            if (!brackets.TryGetValue(c, out var x)) {
                stack.Push(c);
            } else if (stack.Count()!=0 && stack.Peek() == x) {
                stack.Pop();
            } else {
                return false;
            }
        }

        return stack.Count == 0;
    }
}