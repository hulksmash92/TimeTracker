import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { User } from 'src/app/models/user';

const headers = new HttpHeaders({
  'Content-Type': 'application/json'
});

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  readonly GH_API_URL = '/api/github';
  user: User;

  constructor(private http: HttpClient) { }

  /**
   * Gets the url for signing in with GitHub
   */
  gitHubUrl(): Observable<string> {
    return this.http.get(`${this.GH_API_URL}/url`).pipe(
      map((res: any) => res?.data)
    );
  }

  /**
   * Passes the session code to the api to get the users GH access token 
   * and create a log in session for the user
   * 
   * @param sessionCode session code returned by the github oauth server
   * 
   * @returns details of the user
   */
  loginGitHub(sessionCode: string): Observable<any> {
    const url = `${this.GH_API_URL}/login`;
    const body = JSON.stringify({sessionCode});
    return this.http.post(url, body, { headers });
  }

}
