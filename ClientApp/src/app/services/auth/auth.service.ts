import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

const headers = new HttpHeaders({
  'Content-Type': 'application/json'
});

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  readonly GH_API_URL = '/api/github';

  constructor(private http: HttpClient) { }

  gitHubUrl(): Observable<string> {
    return this.http.get(`${this.GH_API_URL}/url`).pipe(
      map((res: any) => res?.data)
    );
  }

  loginGitHub(sessionCode: string): Observable<any> {
    const url = `${this.GH_API_URL}/login`;
    const body = JSON.stringify({sessionCode});
    return this.http.post(url, body, { headers, withCredentials: true });
  }

}
