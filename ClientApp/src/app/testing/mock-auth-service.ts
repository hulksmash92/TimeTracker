import { Observable, of } from 'rxjs';

export class MockAuthService {
    gitHubUrl(): Observable<string> {
        return of(null);
    }

    loginGitHub(sessionCode: string): Observable<any> {
        return of(null);
    }
}
