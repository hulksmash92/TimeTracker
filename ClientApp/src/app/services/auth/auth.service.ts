import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

import { Observable, of } from 'rxjs';
import { catchError, map } from 'rxjs/operators';

import { User } from 'src/app/models/user';

const headers = new HttpHeaders({
  'Content-Type': 'application/json'
});

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  readonly API_URL = '/api/auth';
  readonly GH_API_URL = '/api/github';
  user: User;

  constructor(private http: HttpClient, private router: Router) { }

  /**
   * Checks if the user is authenticated
   */
  async isAuthenticated(): Promise<boolean> {
    try {
      const success = await this.http.get(`${this.API_URL}/isAuthenticated`).pipe(
        map((res: any) => res?.success),
        catchError((err: any) => of(false))
      ).toPromise();

      if (!success) {
        this.user = null;
        this.router.navigate(['']);
      }
      
      return success;
    } catch (e) {
      return false;
    }
  }

  /**
   * Gets the url for signing in with GitHub
   */
  gitHubUrl(): Observable<string> {
    // TODO: Move redirect to GH login page, and return void (will require injecting the WindowService into here)
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
