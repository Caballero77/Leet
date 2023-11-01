// See https://aka.ms/new-console-template for more information

using _380._Insert_Delete_GetRandom_O_1_;

var r = new RandomizedSet();

r.Insert(1);
r.Insert(10);
r.Insert(20);
r.Insert(30);

for (var i = 0; i < 100; i++)
    Console.WriteLine(r.GetRandom());
