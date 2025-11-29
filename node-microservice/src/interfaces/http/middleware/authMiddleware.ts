import { Request, Response, NextFunction } from "express";
import AuthService from "../../../application/services/authService";
import { UserPayload } from "../../../domain/models";

export interface AuthRequest extends Request {
  user?: UserPayload;
}

function createAuthMiddleware(authService: AuthService) {
  return (
    req: AuthRequest,
    res: Response,
    next: NextFunction,
  ): Response | void => {
    try {
      const authHeader = req.headers.authorization;

      if (!authHeader) {
        return res.status(401).json({
          error: "Authorization header is required",
          message: "No token provided",
        });
      }

      const parts = authHeader.split(" ");

      if (parts.length !== 2 || parts[0] !== "Bearer") {
        return res.status(401).json({
          error: "Invalid authorization header format",
          message: "Expected format: Bearer <token>",
        });
      }

      const token = parts[1];

      const decoded = authService.verifyToken(token);

      req.user = decoded;

      next();
    } catch (error: any) {
      return res.status(401).json({
        error: "Authentication failed",
        message: error.message,
      });
    }
  };
}

export default createAuthMiddleware;
