import { ApiClient } from './api-client';
import { Organisation } from './organisation';

export interface User {
    id: number;
    name: string;
    email: string;
    created: Date;
    updated: Date;
    githubUserId: string;
    avatar: string;
    organisation: Organisation[];
    apiClient: ApiClient[];
}
