namespace _289._Game_of_Life;

public class Solution
{
    private IEnumerable<(int, int)> Around(int i, int j)
    {
        yield return (i + 1, j + 1);
        yield return (i, j + 1);
        yield return (i - 1, j + 1);
        yield return (i + 1, j);
        yield return (i - 1, j);
        yield return (i + 1, j - 1);
        yield return (i, j - 1);
        yield return (i - 1, j - 1);
    }
    
    public void GameOfLife(int[][] board) {
        for (var i = 0; i < board.Length; i++)
        {
            for (var j = 0; j < board[i].Length; j++)
            {
                var ones = Around(i, j)
                    .Where(pair =>
                        pair.Item1 >= 0 && pair.Item1 < board.Length && pair.Item2 >= 0 &&
                        pair.Item2 < board[pair.Item1].Length)
                    .Count(pair => board[pair.Item1][pair.Item2]%2 == 1);

                if (ones <= 1)
                {
                    board[i][j] += 2;
                }
                if (ones == 3)
                {
                    board[i][j] += 4;
                }
                if (ones >= 4)
                {
                    board[i][j] += 2;
                }
            }
        }

        for (var i = 0; i < board.Length; i++)
        {
            for (var j = 0; j < board[i].Length; j++)
            {
                if (board[i][j] == 2 || board[i][j] == 3)
                {
                    board[i][j] = 0;
                }

                if (board[i][j] == 4 || board[i][j] == 5)
                {
                    board[i][j] = 1;
                }
            }
        }
    }
}