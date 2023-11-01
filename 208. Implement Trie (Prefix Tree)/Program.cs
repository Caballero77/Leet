// See https://aka.ms/new-console-template for more information
using _208._Implement_Trie__Prefix_Tree_;

var trie = new Trie();

trie.Insert("apple");
trie.Insert("app");
trie.Insert("app");
trie.Insert("apple");

Console.WriteLine(trie.Search("app"));
Console.WriteLine(trie.Search("appl"));
