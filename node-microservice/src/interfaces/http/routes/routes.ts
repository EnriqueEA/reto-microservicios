import { Application, Response, NextFunction } from "express";
import HealthController from "../controllers/healthController";
import StatisticsController from "../controllers/statisticsController";
import AuthController from "../controllers/authController";
import { AuthRequest } from "../middleware/authMiddleware";

function setupRoutes(
  app: Application,
  healthController: HealthController,
  statisticsController: StatisticsController,
  authController: AuthController,
  authMiddleware: (req: AuthRequest, res: Response, next: NextFunction) => void,
): void {
  app.get("/health", healthController.check);
  app.post("/auth/login", authController.login);

  app.post("/stats", authMiddleware, statisticsController.calculateStats);
  app.get("/auth/verify", authMiddleware, authController.verify);
}

export default setupRoutes;
