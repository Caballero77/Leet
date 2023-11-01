namespace _211._Design_Add_and_Search_Words_Data_Structure;

public class WordDictionary
{
    private class Node
    {
        public Node[] Next = new Node[26];

        public bool End;
    }

    private Node root;

    public WordDictionary()
    {
        this.root = new Node { End = false };
    }

    private void AddWordInternal(Node root, string word)
    {
        var next = root.Next[(byte)word[0] - (byte)'a'];
        if (next != null)
        {
            if (word.Length == 1)
            {
                next.End = true;
                return;
            }
            AddWordInternal(next, word.Remove(0, 1));
        }
        else
        {
            root.Next[(byte)word[0] - (byte)'a'] = new Node { End = word.Length == 1 };
            if (word.Length == 1)
                return;
            AddWordInternal(root.Next[(byte)word[0] - (byte)'a'], word.Remove(0, 1));
        }
    }

    public void AddWord(string word) => AddWordInternal(this.root, word);

    private bool SearchInternal(Node root, string word)
    {
        if (word.Length == 0) return root.End;
        var curr = word[0];
        if (curr == '.')
        {
            foreach (var node in root.Next)
            {
                if (node != null && SearchInternal(node, word.Remove(0,1)))
                {
                    return true;
                }
            }

            return false;
        }
        
        var next = root.Next[(byte)word[0] - (byte)'a'];
        if (next != null)
        {
            return SearchInternal(next, word.Remove(0, 1));
        }

        return false;
    }

    public bool Search(string word) => SearchInternal(this.root, word);
}
