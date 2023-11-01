namespace _208._Implement_Trie__Prefix_Tree_;

public class Trie {
    private class TrieNode
    {
        public IDictionary<char, TrieNode> Next;

        public bool Last;
    }

    private TrieNode Root;

    public Trie()
    {
        this.Root = new TrieNode { Next = new Dictionary<char, TrieNode>() };
    }

    private void InsertInternal(TrieNode root, string word)
    {
        if (word.Length == 0)
        {
            return;
        }
        if (root.Next.TryGetValue(word[0], out var next))
        {
            InsertInternal(next, word.Remove(0,1));
            if (word.Length == 1)
            {
                next.Last = true;
            }
        }
        else
        {
            root.Next.Add(word[0], new TrieNode { Next = new Dictionary<char, TrieNode>(), Last = word.Length == 1 });
            InsertInternal(root.Next[word[0]], word.Remove(0,1));
        }
    }
    
    public void Insert(string word) => InsertInternal(this.Root, word);

    private bool SearchInternal(TrieNode root, string word)
    {
        if (root.Next.TryGetValue(word[0], out var next))
        {
            if (word.Length == 1)
            {
                return next.Last;
            }
            return SearchInternal(next, word.Remove(0,1));
        }

        return false;
    }

    public bool Search(string word) => SearchInternal(this.Root, word);
    
    private bool StartsWithInternal(TrieNode root, string word)
    {
        if (word.Length == 0)
        {
            return true;
        }
        if (root.Next.TryGetValue(word[0], out var next))
        {
            return StartsWithInternal(next, word.Remove(0,1));
        }

        return false;
    }


    public bool StartsWith(string prefix) => StartsWithInternal(this.Root, prefix);
}