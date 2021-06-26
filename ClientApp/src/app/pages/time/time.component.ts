import { Component, OnDestroy, OnInit } from '@angular/core';
import { BehaviorSubject, merge, Subscription } from 'rxjs';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { TimeEntry } from 'src/app/models/time-entry';
import { TimeService } from 'src/app/services/time/time.service';

@Component({
  selector: 'app-time',
  templateUrl: './time.component.html',
  styleUrls: ['./time.component.scss']
})
export class TimeComponent implements OnInit, OnDestroy {
  private filterSub: Subscription = new Subscription();
  showForm: boolean;
  data: TimeEntry[] = [];
  editing: TimeEntry;
  dateFrom: BehaviorSubject<Date> = new BehaviorSubject<Date>(null);
  dateTo: BehaviorSubject<Date> = new BehaviorSubject<Date>(null);
  repo: BehaviorSubject<string> = new BehaviorSubject<string>(null);

  constructor(private readonly timeService: TimeService) {
    this.resetParams();
  }

  ngOnInit(): void {
    this.filterSub = merge(
      this.dateFrom.pipe(distinctUntilChanged(), debounceTime(200)),
      this.dateTo.pipe(distinctUntilChanged(), debounceTime(200)),
      this.repo.pipe(distinctUntilChanged()),
    )
    .subscribe(() => {
      this.get();
    });
  }

  ngOnDestroy(): void {
    this.filterSub.unsubscribe();
    this.dateFrom.complete();
    this.dateTo.complete();
    this.repo.complete();
  }

  /**
   * Resets the table filter params
   */
  resetParams(): void {
    this.repo.next(null);
    this.dateTo.next(new Date());

    const d = new Date();
    d.setDate(d.getDate() - 29);
    this.dateFrom.next(d);
  }

  /**
   * Gets time entries that a user has access to
   */
  get(): void {
    const dtF = this.dateFrom.value;
    const dtT = this.dateTo.value;
    const repo = this.repo.value;

    this.timeService.get(dtF, dtT, repo)
      .subscribe((res: TimeEntry[]) => {
        this.data = res;
      });
  }

  /**
   * Creates a new time entry with the selected values
   * 
   * @param data values for the new time entry
   */
  create(data: TimeEntry): void {
    this.timeService.create(data)
      .subscribe((res: TimeEntry) => {
        if (!!res) {
          this.data.push(res);
        }
      });
  }

  /**
   * Updates the selected time entry
   * 
   * @param id id of the item to update
   * @param newValues key-value pairs for the properties to update
   */
  update(id: number, newValues: any): void {
    if (id > 0) {
      this.timeService.update(id, newValues)
        .subscribe((entry: TimeEntry) => {
          if (!!entry) {
            const data = [...this.data];
            const index = data.findIndex(t => t.id === entry.id);
            if (index > -1) {
              data[index] = entry;
              this.data = [...data];
            }
          }
        });
    }
  }

  /**
   * Deletes the selected time entry
   * 
   * @param id id of the item to update
   */
  delete(id: number): void {
    if (id > 0) {
      this.timeService.delete(id)
        .subscribe((success: boolean) => {
          if (success) {
            const data = [...this.data];
            const index = data.findIndex(t => t.id === id);
            if (index > -1) {
              data.splice(index, 1);
              this.data = [...data];
            }
          }
        });
    }
  }

}
