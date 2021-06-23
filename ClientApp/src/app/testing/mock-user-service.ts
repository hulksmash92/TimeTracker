import { Observable, of } from 'rxjs';

export class MockUserService {

    get(): Observable<any> {
        return of({});
    }
}
