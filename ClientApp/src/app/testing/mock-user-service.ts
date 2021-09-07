import { Observable, of } from 'rxjs';

export class MockUserService {
    readonly API_URL = '/api/user';

    get(): Observable<any> {
        return of({});
    }

    update(newValues: any): Observable<any> {
        return of({success: true});
    }

    delete(): Observable<any> {
        return of({success: true});
    }
}
