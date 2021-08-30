import { Observable, of } from 'rxjs';

export class MockAuthService {
    isAuthenticated(): Promise<boolean> {
        return Promise.resolve(true);
      }

    gitHubUrl(): Observable<string> {
        return of(null);
    }

    loginGitHub(sessionCode: string): Observable<any> {
        return of(null);
    }
}
