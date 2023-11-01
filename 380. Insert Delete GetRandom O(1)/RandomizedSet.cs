using System.Collections;

namespace _380._Insert_Delete_GetRandom_O_1_;

using System.Collections.Generic;

public class RandomizedSet
{
    private Random rand = new Random();

    private Dictionary<int, int> dict = new Dictionary<int, int>();

    private List<int> list = new List<int>();

    public bool Insert(int val)
    {
        if (dict.TryGetValue(val, out _))
        {
            return false;
        }
        list.Add(val);
        dict.Add(val, list.Count - 1);
        return true;
    }
    
    public bool Remove(int val) {
        if (dict.TryGetValue(val, out var index))
        {
            (list[index], list[^1]) = (list[^1], list[index]);
            dict[list[index]] = index;
            list.RemoveAt(list.Count - 1);
            dict.Remove(val);
            return true;
        }
        return false;
    }
    
    public int GetRandom()
    {
        return list[rand.Next(0, list.Count)];
    }
}