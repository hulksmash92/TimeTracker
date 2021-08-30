import { Observable, of } from 'rxjs';

export class MockUserService {
    readonly API_URL = '/api/user';

    get(): Observable<any> {
        return of({});
    }
}
