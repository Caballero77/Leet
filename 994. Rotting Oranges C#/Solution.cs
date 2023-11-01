using System.Collections.Generic;

public class Solution {
    public int OrangesRotting(int[][] grid) {
        var minutes = 0;
        var fresh = 0;

        var rotten = new Queue<(int, int)>();
        for (var i = 0; i < grid.Length; i++) {
            for (var j = 0; j < grid[i].Length; j++) {
                if (grid[i][j] == 2) {
                    rotten.Enqueue((i, j));
                } else if (grid[i][j] == 1) {
                    fresh++;
                }
            }
        }

        if (fresh == 0) {
            return 0;
        }
        
        rotten.Enqueue((-1, -1));

        while (rotten.Count > 0) {
            var (i, j) = rotten.Dequeue();

            if (i == j && i == -1 && rotten.Count > 0) {
                minutes++;
                rotten.Enqueue((-1, -1));
            } else {
                ProcessIfFresh(i + 1, j);
                ProcessIfFresh(i - 1, j);
                ProcessIfFresh(i, j + 1);
                ProcessIfFresh(i, j - 1);
            }
        }

        return fresh == 0 ? minutes : -1;

        void ProcessIfFresh(int i, int j){
            if(i >= 0 && i < grid.Length && j >= 0 && j < grid[i].Length) {
                if (grid[i][j] == 1) {
                    rotten?.Enqueue((i,j));
                    grid[i][j] = 2;
                    fresh--;
                }
            }
        }
    }
}