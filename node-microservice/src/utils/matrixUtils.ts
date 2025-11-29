import { Matrix } from "../domain/models";

export function isDiagonal(matrix: Matrix): boolean {
  if (!Array.isArray(matrix) || matrix.length === 0) {
    return false;
  }

  const rows = matrix.length;
  const cols = matrix[0]?.length || 0;

  if (rows !== cols) {
    return false;
  }

  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      const value = Number(matrix[i][j]);
      if (i !== j && value !== 0) {
        return false;
      }
    }
  }

  return true;
}

export function extractValues(matrix: Matrix): number[] {
  const values: number[] = [];

  if (!Array.isArray(matrix)) {
    return values;
  }

  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      const value = Number(matrix[i][j]);
      if (!Number.isNaN(value)) {
        values.push(value);
      }
    }
  }

  return values;
}
