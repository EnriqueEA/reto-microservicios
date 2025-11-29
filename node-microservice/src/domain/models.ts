export type Matrix = number[][];

export interface Statistics {
  max: number;
  min: number;
  sum: number;
  avg: number;
  diagonal: { [key: string]: boolean };
}

export interface StatisticsRequest {
  matrices: { [key: string]: Matrix };
}

export interface HealthResponse {
  status: string;
  service: string;
}

export interface UserPayload {
  userId: number;
  username: string;
  role: string;
}

export interface User extends UserPayload {
  id: number;
  password?: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  message: string;
  token: string;
  user: {
    id: number;
    username: string;
    role: string;
  };
}

export interface AuthenticatedRequest extends Request {
  user?: UserPayload;
}
