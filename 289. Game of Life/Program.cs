using _289._Game_of_Life;

var board = new []
{
    new [] { 0,1,0},
    new [] { 0,0,1},
    new [] { 1,1,1},
    new [] { 0,0,0},
};
for (var i = 0; i < board.Length; i++)
{
    for (int j = 0; j < board[i].Length; j++)
    {
        Console.Write($"{board[i][j]} ");
    }
    Console.WriteLine();
}

Console.WriteLine("____________________________");

new Solution().GameOfLife(board);

for (var i = 0; i < board.Length; i++)
{
    for (int j = 0; j < board[i].Length; j++)
    {
        Console.Write($"{board[i][j]} ");
    }
    Console.WriteLine();
}