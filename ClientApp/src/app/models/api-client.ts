export interface ApiClient {
    clientId?: string;
    secretKey?: string;
    appName: string;
    description: string;
    validTill: Date;
    created?: Date;
    updated?: Date;
    userId: number;
}