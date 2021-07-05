import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TimeFormTagsComponent } from './time-form-tags.component';

describe('TimeFormTagsComponent', () => {
  let component: TimeFormTagsComponent;
  let fixture: ComponentFixture<TimeFormTagsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TimeFormTagsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TimeFormTagsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
