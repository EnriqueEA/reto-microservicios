import { isDiagonal, extractValues } from "../../utils/matrixUtils";
import { Statistics, Matrix } from "../../domain/models";

class StatisticsService {
  calculateStatistics(matrices: { [key: string]: Matrix }): Statistics {
    if (!matrices || typeof matrices !== "object") {
      throw new Error("Invalid matrices object");
    }

    const allValues: number[] = [];
    const diagonal: { [key: string]: boolean } = {};
    const matrixNames = Object.keys(matrices);

    matrixNames.forEach((name) => {
      const matrix = matrices[name];
      const values = extractValues(matrix);
      allValues.push(...values);
      diagonal[name] = isDiagonal(matrix);
    });

    if (allValues.length === 0) {
      throw new Error("No numeric values found in matrices");
    }

    const sum = allValues.reduce((acc, val) => acc + val, 0);
    const max = Math.max(...allValues);
    const min = Math.min(...allValues);
    const avg = sum / allValues.length;

    return {
      max,
      min,
      sum,
      avg,
      diagonal,
    };
  }
}

export default StatisticsService;
