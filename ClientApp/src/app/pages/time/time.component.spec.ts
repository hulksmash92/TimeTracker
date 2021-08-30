import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatNativeDateModule } from '@angular/material/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatListModule } from '@angular/material/list';
import { MatMenuModule } from '@angular/material/menu';
import { MatSelectModule } from '@angular/material/select';
import { MatSortModule } from '@angular/material/sort';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatTableModule } from '@angular/material/table';

import { TimeComponent } from './time.component';
import { TimeFormComponent } from './time-form/time-form.component';
import { TimeFormTagsComponent } from './components/time-form-tags/time-form-tags.component';
import { TimeTableComponent } from './time-table/time-table.component';
import { RepoItemSearchComponent } from './components/repo-item-search/repo-item-search.component';
import { RepoSearchComponent } from './components/repo-search/repo-search.component';
import { TimeService } from 'src/app/services/time/time.service';
import { RepoService } from 'src/app/services/repo/repo.service';
import { MockRepoService, MockTimeService } from 'src/app/testing';

describe('TimeComponent', () => {
  let component: TimeComponent;
  let fixture: ComponentFixture<TimeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        TimeComponent,
        TimeFormComponent,
        TimeFormTagsComponent,
        TimeTableComponent,
        RepoItemSearchComponent,
        RepoSearchComponent
      ],
      imports: [
        NoopAnimationsModule,
        FormsModule,
        ReactiveFormsModule,
        MatAutocompleteModule,
        MatButtonModule,
        MatCardModule,
        MatChipsModule,
        MatDatepickerModule,
        MatFormFieldModule,
        MatIconModule,
        MatInputModule,
        MatListModule,
        MatNativeDateModule,
        MatMenuModule,
        MatSelectModule,
        MatSortModule,
        MatPaginatorModule,
        MatTableModule
      ],
      providers: [
        { provide: RepoService, useClass: MockRepoService },
        { provide: TimeService, useClass: MockTimeService },
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TimeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
