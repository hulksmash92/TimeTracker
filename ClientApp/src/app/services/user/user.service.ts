import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  readonly API_URL = '/api/user';

  constructor(private http: HttpClient) { }

  get(): Observable<any> {
    return this.http.get(`/api/user`, { withCredentials: true });
  }

  // TODO: Add user modification calls

}
