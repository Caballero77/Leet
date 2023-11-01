namespace _151._Reverse_Words_in_a_String;

public class Solution {
    public string ReverseWords(string s)
    {
        var reversedWords = s
            .Split(' ')
            .Where(word => word.Length > 0)
            .Reverse();

        return String.Join(' ', reversedWords);
    }
}