import { Observable, of } from 'rxjs';

export class MockAuthService {
    isAuthenticated(): Promise<boolean> {
        return Promise.resolve(true);
    }

    async gitHubLogin(): Promise<void> {}

    loginGitHub(sessionCode: string): Observable<any> {
        return of(null);
    }
}
