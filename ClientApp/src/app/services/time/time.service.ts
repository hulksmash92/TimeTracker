import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

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
   * Gets all time entries for the selected conditions
   * 
   * @param from date range start
   * @param to date range end
   * @param repo Git repository that the time entries relate to, null for all
   * 
   * @returns array of the time entries
   */
  get(from: Date, to: Date, repo: string): Observable<TimeEntry[]> {
    let params = new HttpParams();
    if (!!from) {
      params = params.append('from', this.dateToString(from));
    }
    if (!!to) {
      params = params.append('to', this.dateToString(to));
    }

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


  dateToString(d: Date): string {
    return `${d.getFullYear()}-${d.getMonth()}-${d.getDay()}`;
  }

}
