class Config {
  private port: number;
  private serviceName: string;
  private nodeEnv: string;
  private jwtSecret: string;
  private jwtExpiresIn: string;

  constructor() {
    this.port = parseInt(process.env.PORT || "3000", 10);
    this.serviceName = process.env.SERVICE_NAME || "node-stats-api";
    this.nodeEnv = process.env.NODE_ENV || "development";
    this.jwtSecret =
      process.env.JWT_SECRET || "dev-secret-key-change-in-production";
    this.jwtExpiresIn = process.env.JWT_EXPIRES_IN || "24h";
  }

  getPort(): number {
    return this.port;
  }

  getServiceName(): string {
    return this.serviceName;
  }

  isProduction(): boolean {
    return this.nodeEnv === "production";
  }

  getJwtSecret(): string {
    return this.jwtSecret;
  }

  getJwtExpiresIn(): string {
    return this.jwtExpiresIn;
  }
}

export default new Config();
