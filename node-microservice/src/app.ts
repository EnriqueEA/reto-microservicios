import express, { Application } from "express";
import bodyParser from "body-parser";
import config from "./infrastructure/config/config";
import StatisticsService from "./application/services/statisticsService";
import AuthService from "./application/services/authService";
import HealthController from "./interfaces/http/controllers/healthController";
import StatisticsController from "./interfaces/http/controllers/statisticsController";
import AuthController from "./interfaces/http/controllers/authController";
import createAuthMiddleware from "./interfaces/http/middleware/authMiddleware";
import setupRoutes from "./interfaces/http/routes/routes";

function createApp(): Application {
  const app = express();
  app.use(bodyParser.json());

  // Inicializar servicios
  const statisticsService = new StatisticsService();
  const authService = new AuthService(
    config.getJwtSecret(),
    config.getJwtExpiresIn(),
  );

  // Inicializar controladores
  const healthController = new HealthController();
  const statisticsController = new StatisticsController(statisticsService);
  const authController = new AuthController(authService);

  // Crear middleware de autenticación
  const authMiddleware = createAuthMiddleware(authService);

  // Configurar rutas
  setupRoutes(
    app,
    healthController,
    statisticsController,
    authController,
    authMiddleware,
  );

  return app;
}

function startServer(): void {
  const app = createApp();
  const port = config.getPort();

  app.listen(port, () => {
    console.log(`Node stats API listening on port ${port}`);
  });
}

if (require.main === module) {
  startServer();
}

export { createApp, startServer };
