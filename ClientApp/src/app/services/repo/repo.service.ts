import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { of, Observable } from 'rxjs';
import { map } from 'rxjs/operators';

import { RepoSearchResult } from 'src/app/models/repos';


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
  searchGitHub(query: string): Observable<RepoSearchResult[]> {
    if (!query) {
      return of(null);
    }
    let params = new HttpParams();
    params = params.append('query', query);

    return this.http.get(`${this.GH_API_URL}/search`, {params}).pipe(
      map((res: any) => res.data || [])
    );
  }

  /**
   * Gets the repo items of the selected type for the selected the repository
   * @param owner owner of the GitHub repo
   * @param repo name of the GitHub repo
   * @param itemType type of item to get i.e. branches or commits
   * @param from date range to search from
   * @param to date range to search until
   * @returns a list of the repo items
   */
  getGitHubRepoItems(owner: string, repo: string, itemType: string, from?: Date, to?: Date): Observable<any[]> {
    let params = new HttpParams();
    if (!!from) {
      params = params.append('from', from.toISOString());
    }
    if (!!to) {
      params = params.append('to', to.toISOString());
    }
    
    return this.http.get(`${this.GH_API_URL}/repo/${owner}/${repo}/${itemType}`, {params}).pipe(
      map((res: any) => res.data || [])
    );
  }

}
