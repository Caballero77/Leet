public class NumMatrix {

    private int[,] sum;

    public NumMatrix(int[][] matrix) {
        sum = new int[matrix.Length, matrix[0].Length];

        for (var i = 0; i < matrix.Length; i++) {
            for (var j = 0; j < matrix[i].Length; j++) {
                var x = matrix[i][j];
                if (i >  0) {
                    x += sum[i-1, j];
                }
                if (j > 0) {
                    x += sum[i, j-1];
                }
                if (i > 0 & j > 0) {
                    x -= sum[i-1,j-1];
                }
                sum[i, j] = x;
            }
        }
    }
    
    public int SumRegion(int row1, int col1, int row2, int col2) {
        var x = sum[row2, col2];
        if (row1 > 0) {
            x -= sum[row1-1, col2];
        }
        if (col1 > 0) {
            x -= sum[row2, col1-1];
        }
        if (row1 > 0 && col1 > 0) {
            x += sum[row1-1, col1-1];
        }
        return x;
    }
}