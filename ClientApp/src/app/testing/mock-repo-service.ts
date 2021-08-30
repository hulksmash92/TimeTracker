import { Observable, of } from 'rxjs';
import { RepoSearchResult } from 'src/app/models/repos';

export class MockRepoService {
    readonly GH_API_URL: string = '/api/github';
  
    /**
     * Searches for GitHub repositories
     * @param query repository search query, for example search by name
     * @returns a list of repos for the search query
     */
    searchGitHub(query: string): Observable<RepoSearchResult[]> {
      return of([]);
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
        return of([]);
    }
  
}
