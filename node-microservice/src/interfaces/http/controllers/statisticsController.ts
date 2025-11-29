import { Request, Response } from "express";
import StatisticsService from "../../../application/services/statisticsService";

class StatisticsController {
  private statisticsService: StatisticsService;

  constructor(statisticsService: StatisticsService) {
    this.statisticsService = statisticsService;
  }

  calculateStats = (req: Request, res: Response): Response => {
    try {
      const { matrices } = req.body;

      if (!matrices) {
        return res.status(400).json({
          error: "Missing matrices in request body",
        });
      }

      const statistics = this.statisticsService.calculateStatistics(matrices);
      return res.json(statistics);
    } catch (error: any) {
      return res.status(400).json({
        error: error.message,
      });
    }
  };
}

export default StatisticsController;
