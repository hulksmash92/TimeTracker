import { Observable, of } from 'rxjs';
import { User } from '../models/user';

export class MockAuthService {
    readonly API_URL = '/api/auth';
    readonly GH_API_URL = '/api/github';
    user: User = {
        id: 1,
        name: 'Alex',
        email: 'alex@example.com',
        githubUserId: 'hulksmash92',
        created: new Date(),
        updated: new Date(),
        avatar: null,
        organisations: [],
        apiClients: []
    };

    isAuthenticated(): Promise<boolean> {
        return Promise.resolve(true);
    }

    async signOut(): Promise<void> {}

    async gitHubLogin(): Promise<void> {}

    loginGitHub(sessionCode: string): Observable<any> {
        return of(null);
    }
}
