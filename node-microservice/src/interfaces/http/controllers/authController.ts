import { Request, Response } from "express";
import AuthService from "../../../application/services/authService";

interface AuthRequest extends Request {
  user?: any;
}

class AuthController {
  private authService: AuthService;

  constructor(authService: AuthService) {
    this.authService = authService;
  }

  login = (req: Request, res: Response): Response => {
    try {
      const { username, password } = req.body;

      if (!username || !password) {
        return res.status(400).json({
          error: "Bad Request",
          message: "Username and password are required",
        });
      }

      const user = this.authService.authenticate(username, password);

      if (!user) {
        return res.status(401).json({
          error: "Authentication failed",
          message: "Invalid username or password",
        });
      }

      const token = this.authService.generateToken({
        userId: user.id,
        username: user.username,
        role: user.role,
      });

      return res.status(200).json({
        message: "Login successful",
        token,
        user: {
          id: user.id,
          username: user.username,
          role: user.role,
        },
      });
    } catch (error: any) {
      return res.status(500).json({
        error: "Internal Server Error",
        message: error.message,
      });
    }
  };

  verify = (req: AuthRequest, res: Response): Response => {
    return res.status(200).json({
      message: "Token is valid",
      user: req.user,
    });
  };
}

export default AuthController;
