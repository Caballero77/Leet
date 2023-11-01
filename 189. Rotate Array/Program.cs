// See https://aka.ms/new-console-template for more information
using _189._Rotate_Array;

var arr = new[] { 1,2,3,4,5,6 };

new Solution()
    .Rotate(arr, 6);

foreach (var i in arr)
{
    Console.Write($"{i} ");
}
Console.WriteLine();
