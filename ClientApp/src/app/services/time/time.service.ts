import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { PaginatedTable } from 'src/app/models/paginated-table';
import { Tag } from 'src/app/models/tag';
import { TimeEntry } from 'src/app/models/time-entry';

const headers = new HttpHeaders({
  'Content-Type': 'application/json'
});

@Injectable({
  providedIn: 'root'
})
export class TimeService {
  readonly API_URL: string = '/api/time';

  constructor(private http: HttpClient) { }

  /**
   * Gets all tags for adding onto time entries
   */
  getTags(): Observable<Tag[]> {
    return this.http.get(`${this.API_URL}/tags`).pipe(
      map((res: any) => res?.data || [])
    );
  }

  /**
   * Gets a page of time entries for the selected conditions
   * 
   * @param from date range start to filter on
   * @param to date range end to filter on
   * @param pageIndex index of the current page (starts at 0)
   * @param pageSize number of records per page
   * @param sort name of the column to sort by
   * @param sortDesc whether to sort descending or not
   * 
   * @returns array of the time entries
   */
  get(from: Date, to: Date, pageIndex: number, pageSize: number, sort: string, sortDesc: boolean): Observable<PaginatedTable<TimeEntry>> {
    let params = new HttpParams();
    if (!!from) {
      params = params.append('from', from.toISOString());
    }
    if (!!to) {
      params = params.append('to', to.toISOString());
    }
    params = params.append('pageIndex', `${(pageIndex ?? 0)}`);
    params = params.append('pageSize', `${(pageSize ?? 10)}`);
    params = params.append('sort', `${(sort || 'created')}`);
    params = params.append('sortDesc', `${(sortDesc ?? true)}`);

    return this.http.get(this.API_URL, { params }).pipe(
      map((res: any) => res?.data || [])
    );
  }

  /**
   * Creates a time entry for the logged in user from the passed in value
   * 
   * @param data values for the new time entry
   * 
   * @returns the newly created time entry
   */
  create(data: TimeEntry): Observable<TimeEntry> {
    return this.http.post(this.API_URL, JSON.stringify(data), { headers }).pipe(
      map((res: any) => res?.data)
    );
  }

  /**
   * Patches the selected properties on the selected time entry
   * 
   * @param id id of the time entry to update
   * @param data values to update
   * 
   * @returns the full new time entry value
   */
  update(id: number, data: any): Observable<TimeEntry> {
    return this.http.patch(`${this.API_URL}/${id}`, JSON.stringify(data), { headers }).pipe(
      map((res: any) => res?.data)
    );
  }

  /**
   * Deletes the selected time entry 
   * 
   * @param id time entry id to delete
   * 
   * @returns true or false if successful
   */
  delete(id: number): Observable<boolean> {
    return this.http.delete(`${this.API_URL}/${id}`).pipe(
      map((res: any) => res?.data)
    );
  }

}
