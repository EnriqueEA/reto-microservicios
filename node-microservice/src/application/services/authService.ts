import jwt from "jsonwebtoken";
import { UserPayload, User } from "../../domain/models";

class AuthService {
  private jwtSecret: string;
  private jwtExpiresIn: string;

  constructor(jwtSecret: string, jwtExpiresIn: string) {
    this.jwtSecret = jwtSecret;
    this.jwtExpiresIn = jwtExpiresIn;
  }

  generateToken(payload: UserPayload): string {
    if (!payload || typeof payload !== "object") {
      throw new Error("Payload must be a valid object");
    }

    return jwt.sign(payload, this.jwtSecret, {
      expiresIn: this.jwtExpiresIn as unknown as any,
    });
  }

  verifyToken(token: string): UserPayload {
    if (!token) {
      throw new Error("Token is required");
    }

    try {
      return jwt.verify(token, this.jwtSecret) as UserPayload;
    } catch (error: any) {
      if (error.name === "TokenExpiredError") {
        throw new Error("Token has expired");
      }
      if (error.name === "JsonWebTokenError") {
        throw new Error("Invalid token");
      }
      throw error;
    }
  }

  authenticate(
    username: string,
    password: string,
  ): Omit<User, "password"> | null {
    const users: User[] = [
      {
        id: 1,
        username: "admin",
        password: "admin123",
        role: "admin",
        userId: 1,
      },
      { id: 2, username: "user", password: "user123", role: "user", userId: 2 },
    ];

    const user = users.find(
      (u) => u.username === username && u.password === password,
    );

    if (!user) {
      return null;
    }

    const { password: _, ...userWithoutPassword } = user;
    return userWithoutPassword;
  }
}

export default AuthService;
