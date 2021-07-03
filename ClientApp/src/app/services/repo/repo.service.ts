import { RepoSearchResult } from './../../models/repos';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { of, Observable } from 'rxjs';
import { map } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class RepoService {
  readonly GH_API_URL: string = '/api/github';

  constructor(private http: HttpClient) { }

  /**
   * Searches for GitHub repositories
   * @param query repository search query, for example search by name
   * @returns a list of repos for the search query
   */
  searchGitHub(query: string): Observable<RepoSearchResult> {
    if (!query) {
      return of(null);
    }
    let params = new HttpParams();
    params = params.append('query', query);

    return this.http.get(`${this.GH_API_URL}/search`, {params}).pipe(
      map((res: any) => res.data)
    );
  }


  getGitHubRepoItems(owner: string, repo: string, itemType: string): Observable<any[]> {
    return this.http.get(`${this.GH_API_URL}/repo/${owner}/${repo}/${itemType}`).pipe(
      map((res: any) => res.data)
    );
  }

}
