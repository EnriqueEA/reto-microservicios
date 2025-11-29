import { Request, Response } from "express";
import config from "../../../infrastructure/config/config";

class HealthController {
  check(_req: Request, res: Response): Response {
    return res.json({
      status: "ok",
      service: config.getServiceName(),
    });
  }
}

export default HealthController;
